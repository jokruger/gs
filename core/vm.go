package core

type VM interface {
	Abort()
	IsStackEmpty() bool
	Call(CompiledFunction, ...Object) (Object, error)
	Run() error
}
