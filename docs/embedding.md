# Embedding Kavun In Go

The recommended embedding API is `kavun.Script`.

It wraps parsing, compilation, globals setup, and VM execution into a higher-level workflow that is easier to integrate and maintain in Go applications.

Direct use of compiler and VM is still available when you need lower-level control, but this document focuses on Script-first usage.

## Quick Start (Script)

This is the primary pattern (also used by benchmarks): create a script, set imports, compile once, run many times.

```go
package main

import (
	"fmt"

	"github.com/jokruger/kavun"
	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/stdlib"
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

	compiled, err := script.Compile(nil, nil)
	if err != nil {
		panic(err)
	}

	if err := compiled.Run(); err != nil {
		panic(err)
	}

	fmt.Println("result:", compiled.GetValue("out"))
}
```

## Script Lifecycle

Use these two workflows depending on your needs:

1. One-shot execution

```go
compiled, err := script.Run() // compile + run
if err != nil {
	panic(err)
}
```

2. Reusable execution (preferred for performance)

```go
compiled, err := script.Compile(nil, nil) // compile once
if err != nil {
	panic(err)
}

for i := 0; i < 100; i++ {
	if err := compiled.Run(); err != nil {
		panic(err)
	}
}
```

`Compiled.Run()` is designed for recurrent execution. It reuses internal VM structures and resets internal runtime resources between runs (allocator-backed runtime values, globals runtime state, and VM state) so you can execute the same compiled script repeatedly.

At the lower-level VM API, reuse is also possible via `vm.Reset(...)`, including swapping bytecode to run different scripts.

## Inputs And Outputs

Set host variables before compile with `script.Add`.

```go
script.Add("x", core.IntValue(20))
script.Add("y", core.IntValue(22))
script.Add("out", core.Undefined)
```

Read values after run:

```go
out := compiled.GetValue("out")
sum, _ := out.AsInt()
fmt.Println(sum)
```

You can update compiled globals between runs:

```go
if err := compiled.Set("x", core.IntValue(50)); err != nil {
	panic(err)
}
if err := compiled.Set("y", core.IntValue(7)); err != nil {
	panic(err)
}
if err := compiled.Run(); err != nil {
	panic(err)
}
```

Other helpers:

- `compiled.Get(name)` returns a `*kavun.Variable`
- `compiled.GetAll()` returns all globals
- `compiled.IsDefined(name)` checks whether a global exists and is not `undefined`

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
- `script.SetMaxFrames(n)`
- `script.SetMaxStack(n)`
- `script.SetAssignmentMode(mode)`

Example:

```go
script.SetMaxConstObjects(10_000)
script.SetMaxFrames(4_096)
script.SetMaxStack(8_192)
script.SetAssignmentMode(kavun.AssignmentModeSmart)
```

## Concurrency

`Compiled` is safe for repeated use from one goroutine, and it can be cloned for parallel execution.

For concurrent runs, create one clone per goroutine with its own runtime arena:

```go
base, err := script.Compile(nil, nil)
if err != nil {
	panic(err)
}

arena := core.NewArena(nil)
clone, err := base.Clone(arena)
if err != nil {
	panic(err)
}

if err := clone.Run(); err != nil {
	panic(err)
}
```

`RunContext(ctx)` is also available for cancellable execution.

## Allocators

`Script.Compile(cta, rta)` accepts two allocators:

- `cta`: compile-time allocator
- `rta`: run-time allocator

If you pass `nil` for either allocator, Kavun creates a default allocator internally.

```go
compiled, err := script.Compile(nil, nil) // both allocators use defaults
if err != nil {
	panic(err)
}
```

Compile-time and run-time allocators are separated by design. This avoids mixing compiler allocations with runtime allocations and usually gives better reuse characteristics in recurrent execution.

If needed, you can still pass the same allocator instance for both:

```go
arena := core.NewArena(nil)
compiled, err := script.Compile(arena, arena)
if err != nil {
	panic(err)
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
