package core

type VM interface {
	Allocator() Allocator
	Abort()
	IsStackEmpty() bool
	Call(Object, ...Object) (Object, error)
	Run() error
}
