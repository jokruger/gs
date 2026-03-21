package stdlib

import (
	"os/exec"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

func makeOSExecCommand(alloc core.Allocator, cmd *exec.Cmd) *value.Record {
	cmdRun := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.run", "0", len(args))
		}
		return wrapError(alloc, cmd.Run()), nil
	}

	cmdStart := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.start", "0", len(args))
		}
		return wrapError(alloc, cmd.Start()), nil
	}

	cmdWait := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.wait", "0", len(args))
		}
		return wrapError(alloc, cmd.Wait()), nil
	}

	cmdCombinedOutput := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.combined_output", "0", len(args))
		}
		res, err := cmd.CombinedOutput()
		if err != nil {
			return wrapError(alloc, err), nil
		}
		if len(res) > core.MaxBytesLen {
			return nil, core.NewBytesLimitError("os.exec.combined_output")
		}
		return alloc.NewBytes(res), nil
	}

	cmdOutput := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.output", "0", len(args))
		}
		res, err := cmd.Output()
		if err != nil {
			return wrapError(alloc, err), nil
		}
		if len(res) > core.MaxBytesLen {
			return nil, core.NewBytesLimitError("os.exec.output")
		}
		return alloc.NewBytes(res), nil
	}

	cmdSetPath := func(alloc core.Allocator, args ...core.Object) (core.Object, error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("os.exec.set_path", "1", len(args))
		}
		s1, ok := args[0].AsString()
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("os.exec.set_path", "first", "string(compatible)", args[0])
		}
		cmd.Path = s1
		return alloc.NewUndefined(), nil
	}

	cmdSetDir := func(alloc core.Allocator, args ...core.Object) (core.Object, error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("os.exec.set_dir", "1", len(args))
		}
		s1, ok := args[0].AsString()
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("os.exec.set_dir", "first", "string(compatible)", args[0])
		}
		cmd.Dir = s1
		return alloc.NewUndefined(), nil
	}

	cmdSetEnv := func(alloc core.Allocator, args ...core.Object) (core.Object, error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("os.exec.set_env", "1", len(args))
		}

		var env []string
		var err error
		switch arg0 := args[0].(type) {
		case *value.Array:
			env, err = stringArray(arg0.Value(), "first")
			if err != nil {
				return nil, err
			}
		default:
			return nil, core.NewInvalidArgumentTypeError("os.exec.set_env", "first", "array(string)", args[0])
		}
		cmd.Env = env
		return alloc.NewUndefined(), nil
	}

	cmdProcess := func(alloc core.Allocator, args ...core.Object) (core.Object, error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("os.exec.process", "0", len(args))
		}
		return makeOSProcess(alloc, cmd.Process), nil
	}

	return alloc.NewRecord(map[string]core.Object{
		"combined_output": alloc.NewBuiltinFunction("combined_output", cmdCombinedOutput, 0, false), // combined_output() => bytes/error
		"output":          alloc.NewBuiltinFunction("output", cmdOutput, 0, false),                  // output() => bytes/error
		"run":             alloc.NewBuiltinFunction("run", cmdRun, 0, false),                        // run() => error
		"start":           alloc.NewBuiltinFunction("start", cmdStart, 0, false),                    // start() => error
		"wait":            alloc.NewBuiltinFunction("wait", cmdWait, 0, false),                      // wait() => error
		"set_path":        alloc.NewBuiltinFunction("set_path", cmdSetPath, 1, false),               // set_path(path string)
		"set_dir":         alloc.NewBuiltinFunction("set_dir", cmdSetDir, 1, false),                 // set_dir(dir string)
		"set_env":         alloc.NewBuiltinFunction("set_env", cmdSetEnv, 1, false),                 // set_env(env array(string))
		"process":         alloc.NewBuiltinFunction("process", cmdProcess, 0, false),                // process() => imap(process)
	}, true).(*value.Record)
}
