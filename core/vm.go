package core

type VM interface {
	Allocator() Allocator
	Abort()
	IsStackEmpty() bool
	Call(Object, ...Value) (Value, error)
	Run() error
}
