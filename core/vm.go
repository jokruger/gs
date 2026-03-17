package core

type VM interface {
	Abort()
	IsStackEmpty() bool
	Call(Object, ...Object) (Object, error)
	Run() error
}
