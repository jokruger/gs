package stdlib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

var osModule = map[string]gs.Object{
	"platform":            &gs.String{Value: runtime.GOOS},
	"arch":                &gs.String{Value: runtime.GOARCH},
	"o_rdonly":            &gs.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &gs.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &gs.Int{Value: int64(os.O_RDWR)},
	"o_append":            &gs.Int{Value: int64(os.O_APPEND)},
	"o_create":            &gs.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &gs.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &gs.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &gs.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &gs.Int{Value: int64(os.ModeDir)},
	"mode_append":         &gs.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &gs.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &gs.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &gs.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &gs.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &gs.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &gs.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &gs.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &gs.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &gs.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &gs.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &gs.Int{Value: int64(os.ModeType)},
	"mode_perm":           &gs.Int{Value: int64(os.ModePerm)},
	"path_separator":      &gs.Char{Value: os.PathSeparator},
	"path_list_separator": &gs.Char{Value: os.PathListSeparator},
	"dev_null":            &gs.String{Value: os.DevNull},
	"seek_set":            &gs.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &gs.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &gs.Int{Value: int64(io.SeekEnd)},
	"args": &gs.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &gs.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &gs.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &gs.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &gs.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &gs.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &gs.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &gs.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &gs.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &gs.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &gs.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &gs.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &gs.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &gs.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &gs.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &gs.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &gs.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &gs.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &gs.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &gs.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &gs.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &gs.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &gs.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &gs.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &gs.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &gs.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &gs.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &gs.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &gs.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &gs.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &gs.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &gs.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &gs.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &gs.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &gs.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &gs.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &gs.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &gs.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &gs.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	fname, ok := gs.ToString(args[0])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > gs.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &gs.Bytes{Value: bytes}, nil
}

func osStat(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	fname, ok := gs.ToString(args[0])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &gs.ImmutableMap{
		Value: map[string]gs.Object{
			"name":  &gs.String{Value: stat.Name()},
			"mtime": &gs.Time{Value: stat.ModTime()},
			"size":  &gs.Int{Value: stat.Size()},
			"mode":  &gs.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = gs.TrueValue
	} else {
		fstat.Value["directory"] = gs.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...gs.Object) (gs.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...gs.Object) (gs.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...gs.Object) (gs.Object, error) {
	if len(args) != 3 {
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
	i2, ok := gs.ToInt(args[1])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := gs.ToInt(args[2])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...gs.Object) (gs.Object, error) {
	if len(args) != 0 {
		return nil, gse.ErrWrongNumArguments
	}
	arr := &gs.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > gs.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		arr.Value = append(arr.Value, &gs.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *gs.UserFunction {
	return &gs.UserFunction{
		Name: name,
		Value: func(args ...gs.Object) (gs.Object, error) {
			if len(args) != 2 {
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
			i2, ok := gs.ToInt64(args[1])
			if !ok {
				return nil, gse.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...gs.Object) (gs.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return gs.FalseValue, nil
	}
	if len(res) > gs.MaxStringLen {
		return nil, gse.ErrStringLimit
	}
	return &gs.String{Value: res}, nil
}

func osExpandEnv(args ...gs.Object) (gs.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > gs.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > gs.MaxStringLen {
		return nil, gse.ErrStringLimit
	}
	return &gs.String{Value: s}, nil
}

func osExec(args ...gs.Object) (gs.Object, error) {
	if len(args) == 0 {
		return nil, gse.ErrWrongNumArguments
	}
	name, ok := gs.ToString(args[0])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := gs.ToString(arg)
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...gs.Object) (gs.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	i1, ok := gs.ToInt(args[0])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...gs.Object) (gs.Object, error) {
	if len(args) != 4 {
		return nil, gse.ErrWrongNumArguments
	}
	name, ok := gs.ToString(args[0])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *gs.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *gs.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := gs.ToString(args[2])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *gs.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *gs.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []gs.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*gs.String)
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
