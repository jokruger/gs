package core

type BuiltinFunction struct {
	Value    NativeFunc
	Name     string
	Arity    int // number of positional arguments, or minimum number of arguments if variadic is true
	Variadic bool
}

func NewStaticBuiltinFunction(name string, val NativeFunc, arity int, variadic bool) Value {
	o := &BuiltinFunction{Value: val, Name: name, Arity: arity, Variadic: variadic}
	return BuiltinFunctionValue(o)
}

func (o *BuiltinFunction) Set(name string, value NativeFunc, arity int, variadic bool) {
	o.Value = value
	o.Name = name
	o.Arity = arity
	o.Variadic = variadic
}
