package core

// CallableFunction is a function signature for the callable functions.
type CallableFunction = func(args ...Object) (ret Object, err error)
