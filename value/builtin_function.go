package value

import "github.com/jokruger/gs/core"

type BuiltinFunction struct {
	Object
	Name  string
	Value core.NativeFunc
}

func (o *BuiltinFunction) TypeName() string {
	return "builtin-function:" + o.Name
}

func (o *BuiltinFunction) String() string {
	return "<builtin-function>"
}

func (o *BuiltinFunction) Copy() core.Object {
	return &BuiltinFunction{Value: o.Value}
}

func (o *BuiltinFunction) Equals(core.Object) bool {
	return false
}

func (o *BuiltinFunction) Call(args ...core.Object) (core.Object, error) {
	return o.Value(args...)
}

func (o *BuiltinFunction) IsCallable() bool {
	return true
}
