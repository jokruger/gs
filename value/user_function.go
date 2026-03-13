package value

type UserFunction struct {
	ObjectImpl
	Name  string
	Value CallableFunc
}

func (o *UserFunction) TypeName() string {
	return "user-function:" + o.Name
}

func (o *UserFunction) String() string {
	return "<user-function>"
}

func (o *UserFunction) Copy() Object {
	return &UserFunction{Value: o.Value, Name: o.Name}
}

func (o *UserFunction) Equals(_ Object) bool {
	return false
}

func (o *UserFunction) Call(args ...Object) (Object, error) {
	return o.Value(args...)
}

func (o *UserFunction) CanCall() bool {
	return true
}
