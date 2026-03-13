package stdlib

import (
	"os/exec"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

func makeOSExecCommand(cmd *exec.Cmd) *gst.ImmutableMap {
	return &gst.ImmutableMap{
		Value: map[string]gst.Object{
			// combined_output() => bytes/error
			"combined_output": &gst.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &gst.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &gst.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &gst.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &gst.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &gst.UserFunction{
				Name: "set_path",
				Value: func(args ...gst.Object) (gst.Object, error) {
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
					return gst.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &gst.UserFunction{
				Name: "set_dir",
				Value: func(args ...gst.Object) (gst.Object, error) {
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
					return gst.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &gst.UserFunction{
				Name: "set_env",
				Value: func(args ...gst.Object) (gst.Object, error) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *gst.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *gst.ImmutableArray:
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
					return gst.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &gst.UserFunction{
				Name: "process",
				Value: func(args ...gst.Object) (gst.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
