package stdlib

import (
	"os/exec"

	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

func makeOSExecCommand(cmd *exec.Cmd) *value.ImmutableMap {
	return &value.ImmutableMap{
		Value: map[string]value.Object{
			// combined_output() => bytes/error
			"combined_output": &value.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &value.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &value.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &value.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &value.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &value.UserFunction{
				Name: "set_path",
				Value: func(args ...value.Object) (value.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					s1, ok := args[0].ToString()
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return value.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &value.UserFunction{
				Name: "set_dir",
				Value: func(args ...value.Object) (value.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					s1, ok := args[0].ToString()
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return value.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &value.UserFunction{
				Name: "set_env",
				Value: func(args ...value.Object) (value.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *value.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *value.ImmutableArray:
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
					return value.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &value.UserFunction{
				Name: "process",
				Value: func(args ...value.Object) (value.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
