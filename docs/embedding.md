# Embedding Kavun In Go

The recommended embedding API is `kavun.Script`. It wraps parsing, compilation, globals setup, and VM execution into a higher-level workflow that is easier to integrate and maintain in Go applications.

Direct use of compiler and VM is still available when you need lower-level control, but this document focuses on Script-first usage.

## Quick Start (Script)

This is the primary pattern: create a script, compile once, then run repeatedly with explicit runtime resources.

```go
package main

import (
	"fmt"

	"github.com/jokruger/kavun"
	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/stdlib"
	"github.com/jokruger/kavun/vm"
)

func main() {
	src := []byte(`
fib := func(x) {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
out = fib(10)
`)

	script := kavun.NewScript(src)
	script.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	script.Add("out", core.Undefined)

    // Compile-time allocator.
	cta := core.NewArena(nil)

    // Runtime allocator and VM.
	rta := core.NewArena(nil)
	machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)

    // Compile once.
	compiled, err := script.Compile(cta)
	if err != nil {
		panic(err)
	}

    // Run repeatedly with the same compiled code and runtime resources.
    for i := 0; i < 100; i++ {
	    if err := compiled.Run(rta, machine); err != nil {
		    panic(err)
	    }
    }

	fmt.Println("result:", compiled.GetValue("out"))
}
```

`Compiled.Run(...)` resets the runtime allocator and reinitializes VM state before each execution.

At lower-level, reuse is done with `rta.Reset()` and `machine.Reset(rta, bytecode, globals)`.

## Inputs And Outputs

Set host variables before compile with `script.Add`.

```go
script.Add("x", core.IntValue(20))
script.Add("y", core.IntValue(22))
script.Add("out", core.Undefined)
```

After compilation, `compiled.Set(...)` prepares input globals for the next execution.
It does not update the runtime state exposed by `Get`, `GetValue`, or `GetAll` directly.

```go
if err := compiled.Set("x", core.IntValue(50)); err != nil {
	panic(err)
}
if err := compiled.Set("y", core.IntValue(7)); err != nil {
	panic(err)
}
if err := compiled.Run(rta, machine); err != nil {
	panic(err)
}
```

`Get`, `GetValue`, and `GetAll` read runtime global variables produced by the last script execution.
In practice, `Set` configures the inputs, and `Get*` reads the outputs, so `Get*` should be used only after the script has run.

Read values after run:

```go
out := compiled.GetValue("out")
sum, _ := out.AsInt()
fmt.Println(sum)
```

Other helpers:

- `compiled.Get(name)` returns a `*kavun.Variable`
- `compiled.GetAll()` returns all globals

## Imports

Use a module map to control what `import("...")` can load.

```go
modules := vm.NewModuleMap()

// Selected stdlib modules
modules.AddMap(stdlib.GetModuleMap("math", "json"))

// Host builtin module
modules.AddBuiltinModule("host", map[string]core.Value{
	"answer": core.IntValue(42),
})

// In-memory source module
modules.AddSourceModule("helpers", []byte(`
export add := func(a, b) { return a + b }
`))

script.SetImports(modules)
```

Suggested default for general-purpose apps:

```go
script.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
```

## File Imports

Local file imports are disabled by default.

Enable them explicitly:

```go
script.EnableFileImport(true)
if err := script.SetImportDir("./scripts"); err != nil {
	panic(err)
}
```

`Script` does not expose file extension customization.
If you need custom source extensions for file imports, use the lower-level compiler API (`Compiler.SetImportFileExt`) directly.

## Runtime And Compiler Limits

`Script` exposes common limits and execution controls:

- `script.SetMaxConstObjects(n)`
- `script.SetAssignmentMode(mode)`

Example:

```go
script.SetMaxConstObjects(10_000)
script.SetAssignmentMode(kavun.AssignmentModeSmart)
```

## Concurrency

`Script`, `Compiled`, `VM`, and allocator helpers are not thread-safe.

If you need parallel execution, user code must provide synchronization and isolated runtime resources.

Safe pattern for parallel runs:

- each goroutine uses its own `Compiled` (for example via `Clone`)
- each goroutine uses its own runtime arena and VM
- shared resources are protected with explicit locking

Example:

```go
base, err := script.Compile(core.NewArena(nil))
if err != nil {
	panic(err)
}

clone, err := base.Clone(core.NewArena(nil))
if err != nil {
	panic(err)
}

rta := core.NewArena(nil)
machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)

if err := clone.Run(rta, machine); err != nil {
	panic(err)
}
```

Use `RunContext(ctx, rta, machine)` for cancellable execution.

## Allocators

You must separate compile-time and runtime allocators.

- `script.Compile(cta)` uses compile-time allocator
- `compiled.Run(rta, machine)` / `compiled.RunContext(ctx, rta, machine)` use runtime allocator

Do not reuse the same allocator instance for compile-time and runtime paths.
Runtime execution resets the runtime allocator, so using the same allocator for both can invalidate compile-time data when VM is reused.

```go
cta := core.NewArena(nil)
rta := core.NewArena(nil)
machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)

compiled, err := script.Compile(cta)
if err != nil {
	panic(err)
}

if err := compiled.Run(rta, machine); err != nil {
	panic(err)
}
```

If `cta` passed to `Compile` is `nil`, Kavun creates a default compile-time allocator internally.

## Lazy Resource Management And VM.Clear

By default, VM reuse is lazy: stack and frame references are not fully cleared on each run.
This improves performance but can keep some references alive longer (until overwritten).

If you prefer more aggressive release behavior, call `machine.Clear()` explicitly.

```go
if err := compiled.Run(rta, machine); err != nil {
	panic(err)
}

// Optional: release remaining references in stack/frames.
machine.Clear()
```

Use `Clear` when memory pressure is more important than peak throughput.

## One-Shot Helper Pattern

If you still want a one-shot flow in app code, build it explicitly:

```go
func RunOnce(src []byte) error {
	script := kavun.NewScript(src)
	cta := core.NewArena(nil)
	rta := core.NewArena(nil)
	machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)

	compiled, err := script.Compile(cta)
	if err != nil {
		return err
	}
	return compiled.Run(rta, machine)
}
```

### Custom Allocator Payload

Allocator behavior can be extended with a custom payload via `core.ArenaOptions.Payload`.
Payload must implement `Reset()` and is reset together with the arena.

This is useful when embedding user-defined types (see unit tests for custom type registration patterns) and you want type-specific allocation or caches to follow arena lifecycle.

```go
type MyPayload struct {
	buf []byte
}

func (p *MyPayload) Reset() {
	p.buf = p.buf[:0]
}

opts := core.DefaultArenaOptions()
opts.Payload = &MyPayload{}

arena := core.NewArena(opts)
_ = arena.Payload() // retrieve custom payload when needed
```
