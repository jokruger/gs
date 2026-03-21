package core

type NativeFunc = func(Allocator, ...Object) (Object, error)
