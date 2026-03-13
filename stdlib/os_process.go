package stdlib

import (
	"os"
	"syscall"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

func makeOSProcessState(state *os.ProcessState) *gst.ImmutableMap {
	return &gst.ImmutableMap{
		Value: map[string]gst.Object{
			"exited": &gst.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &gst.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &gst.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &gst.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *gst.ImmutableMap {
	return &gst.ImmutableMap{
		Value: map[string]gst.Object{
			"kill": &gst.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &gst.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &gst.UserFunction{
				Name: "signal",
				Value: func(args ...gst.Object) (gst.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					i1, ok := args[0].ToInt64()
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
			"wait": &gst.UserFunction{
				Name: "wait",
				Value: func(args ...gst.Object) (gst.Object, error) {
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
