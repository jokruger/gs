package stdlib

import (
	"os"
	"syscall"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

func makeOSProcessState(state *os.ProcessState) *value.ImmutableMap {
	return &value.ImmutableMap{
		Value: map[string]core.Object{
			"exited": &value.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &value.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &value.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &value.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *value.ImmutableMap {
	return &value.ImmutableMap{
		Value: map[string]core.Object{
			"kill": &value.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &value.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &value.UserFunction{
				Name: "signal",
				Value: func(args ...core.Object) (core.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					i1, ok := args[0].AsInt()
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &value.UserFunction{
				Name: "wait",
				Value: func(args ...core.Object) (core.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
