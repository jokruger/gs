package value

type BuiltinFunction struct {
	ObjectImpl
	Name  string
	Value CallableFunc
}

func (o *BuiltinFunction) TypeName() string {
	return "builtin-function:" + o.Name
}

func (o *BuiltinFunction) String() string {
	return "<builtin-function>"
}

func (o *BuiltinFunction) Copy() Object {
	return &BuiltinFunction{Value: o.Value}
}

func (o *BuiltinFunction) Equals(Object) bool {
	return false
}

func (o *BuiltinFunction) Call(args ...Object) (Object, error) {
	return o.Value(args...)
}

func (o *BuiltinFunction) CanCall() bool {
	return true
}
