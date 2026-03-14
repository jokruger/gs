package value

import "github.com/jokruger/gs/core"

type UserFunction struct {
	ObjectImpl
	Name  string
	Value core.CallableFunction
}

func (o *UserFunction) TypeName() string {
	return "user-function:" + o.Name
}

func (o *UserFunction) String() string {
	return "<user-function>"
}

func (o *UserFunction) Copy() core.Object {
	return &UserFunction{Value: o.Value, Name: o.Name}
}

func (o *UserFunction) Equals(core.Object) bool {
	return false
}

func (o *UserFunction) Call(args ...core.Object) (core.Object, error) {
	return o.Value(args...)
}

func (o *UserFunction) CanCall() bool {
	return true
}
