package core

type NativeFunc = func(args ...Object) (ret Object, err error)
