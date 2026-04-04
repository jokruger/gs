package value

import (
	"fmt"

	"github.com/jokruger/gs/core"
)

/*
	CompiledFunction is different from other objects as it is created and managed by compiler and VM directly.
*/

type CompiledFunction struct {
	Object
	Instructions  []byte
	NumLocals     int // number of local variables (including function parameters)
	NumParameters int
	VarArgs       bool
	SourceMap     map[int]core.Pos
	Free          []*core.Value
}

func (o *CompiledFunction) TypeName() string {
	if o.VarArgs {
		return fmt.Sprintf("<compiled-function/%d+>", o.NumParameters)
	}
	return fmt.Sprintf("<compiled-function/%d>", o.NumParameters)
}

func (o *CompiledFunction) String() string {
	return o.TypeName()
}

func (o *CompiledFunction) Arity() int {
	return o.NumParameters
}

func (o *CompiledFunction) IsVariadic() bool {
	return o.VarArgs
}

func (o *CompiledFunction) IsImmutable() bool {
	return true
}

func (o *CompiledFunction) Size() int64 {
	return int64(len(o.Instructions) + len(o.SourceMap) + len(o.Free))
}

func (o *CompiledFunction) Copy(core.Allocator) core.Value {
	t := &CompiledFunction{
		Instructions:  append([]byte{}, o.Instructions...),
		NumLocals:     o.NumLocals,
		NumParameters: o.NumParameters,
		VarArgs:       o.VarArgs,
		Free:          append([]*core.Value{}, o.Free...), // DO NOT Copy() of elements; these are variable pointers
	}
	return core.NewObject(t, false)
}

func (o *CompiledFunction) Method(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	return core.NewUndefined(), core.NewInvalidMethodError(name, o.TypeName())
}

func (o *CompiledFunction) Access(core.VM, core.Value, core.Opcode) (core.Value, error) {
	return core.NewUndefined(), core.NewNotAccessibleError(o.TypeName())
}

func (o *CompiledFunction) SourcePos(ip int) core.Pos {
	for ip >= 0 {
		if p, ok := o.SourceMap[ip]; ok {
			return p
		}
		ip--
	}
	return core.NoPos
}

func (o *CompiledFunction) IsCallable() bool {
	return true
}

func (o *CompiledFunction) IsCompiledFunction() bool {
	return true
}

func (o *CompiledFunction) Call(vm core.VM, args ...core.Value) (core.Value, error) {
	return vm.Call(o, args...)
}
