package value

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/token"
)

type BuiltinFunction struct {
	Object
	value    core.NativeFunc
	name     string
	arity    int // number of positional arguments, or minimum number of arguments if variadic is true
	variadic bool
}

func NewStaticBuiltinFunction(name string, val core.NativeFunc, arity int, variadic bool) core.Value {
	o := &BuiltinFunction{}
	o.Set(name, val, arity, variadic)
	return core.ObjectValue(o)
}

func (o *BuiltinFunction) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	var name string
	if err := dec.Decode(&name); err != nil {
		return err
	}

	var arity int
	if err := dec.Decode(&arity); err != nil {
		return err
	}

	var variadic bool
	if err := dec.Decode(&variadic); err != nil {
		return err
	}

	o.Set(name, nil, arity, variadic)
	return nil
}

func (o *BuiltinFunction) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(o.name); err != nil {
		return nil, err
	}
	if err := enc.Encode(o.arity); err != nil {
		return nil, err
	}
	if err := enc.Encode(o.variadic); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *BuiltinFunction) Set(name string, value core.NativeFunc, arity int, variadic bool) {
	o.name = name
	o.value = value
	o.arity = arity
	o.variadic = variadic
}

func (o *BuiltinFunction) Name() string {
	return o.name
}

func (o *BuiltinFunction) TypeName() string {
	if o.variadic {
		return fmt.Sprintf("<builtin-function:%s/%d+>", o.name, o.arity)
	}
	return fmt.Sprintf("<builtin-function:%s/%d>", o.name, o.arity)
}

func (o *BuiltinFunction) String() string {
	return o.TypeName()
}

func (o *BuiltinFunction) Interface() any {
	return o.value
}

func (o *BuiltinFunction) Arity() int {
	return o.arity
}

func (o *BuiltinFunction) BinaryOp(vm core.VM, op token.Token, rhs core.Value) (core.Value, error) {
	return core.UndefinedValue(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
}

func (o *BuiltinFunction) Copy(alloc core.Allocator) core.Value {
	return alloc.NewBuiltinFunctionValue(o.name, o.value, o.arity, o.variadic)
}

func (o *BuiltinFunction) Method(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	return core.UndefinedValue(), core.NewInvalidMethodError(name, o.TypeName())
}

func (o *BuiltinFunction) Access(core.VM, core.Value, core.Opcode) (core.Value, error) {
	return core.UndefinedValue(), core.NewNotAccessibleError(o.TypeName())
}

func (o *BuiltinFunction) Assign(core.Value, core.Value) error {
	return core.NewNotAssignableError(o.TypeName())
}

func (o *BuiltinFunction) Call(vm core.VM, args ...core.Value) (core.Value, error) {
	if o.value == nil {
		return core.UndefinedValue(), core.NewLogicError(fmt.Sprintf("built-in function %s is referencing nil", o.name))
	}
	return o.value(vm, args...)
}

func (o *BuiltinFunction) IsCallable() bool {
	return true
}

func (o *BuiltinFunction) IsImmutable() bool {
	return true
}

func (o *BuiltinFunction) IsVariadic() bool {
	return o.variadic
}

func (o *BuiltinFunction) IsBuiltinFunction() bool {
	return true
}
