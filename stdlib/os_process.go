package stdlib

import (
	"os"
	"syscall"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

func makeOSProcessState(state *os.ProcessState) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			"exited": &gs.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &gs.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &gs.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &gs.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			"kill": &gs.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &gs.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &gs.UserFunction{
				Name: "signal",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					i1, ok := gs.ToInt64(args[0])
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
			"wait": &gs.UserFunction{
				Name: "wait",
				Value: func(args ...gs.Object) (gs.Object, error) {
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
