package stdlib

import (
	"os/exec"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

func makeOSExecCommand(cmd *exec.Cmd) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			// combined_output() => bytes/error
			"combined_output": &gs.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &gs.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &gs.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &gs.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &gs.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &gs.UserFunction{
				Name: "set_path",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					s1, ok := gs.ToString(args[0])
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return gs.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &gs.UserFunction{
				Name: "set_dir",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					s1, ok := gs.ToString(args[0])
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return gs.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &gs.UserFunction{
				Name: "set_env",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *gs.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *gs.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return gs.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &gs.UserFunction{
				Name: "process",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
