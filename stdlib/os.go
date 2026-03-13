package stdlib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

var osModule = map[string]gst.Object{
	"platform":            &gst.String{Value: runtime.GOOS},
	"arch":                &gst.String{Value: runtime.GOARCH},
	"o_rdonly":            &gst.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &gst.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &gst.Int{Value: int64(os.O_RDWR)},
	"o_append":            &gst.Int{Value: int64(os.O_APPEND)},
	"o_create":            &gst.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &gst.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &gst.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &gst.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &gst.Int{Value: int64(os.ModeDir)},
	"mode_append":         &gst.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &gst.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &gst.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &gst.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &gst.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &gst.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &gst.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &gst.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &gst.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &gst.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &gst.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &gst.Int{Value: int64(os.ModeType)},
	"mode_perm":           &gst.Int{Value: int64(os.ModePerm)},
	"path_separator":      &gst.Char{Value: os.PathSeparator},
	"path_list_separator": &gst.Char{Value: os.PathListSeparator},
	"dev_null":            &gst.String{Value: os.DevNull},
	"seek_set":            &gst.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &gst.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &gst.Int{Value: int64(io.SeekEnd)},
	"args": &gst.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &gst.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &gst.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &gst.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &gst.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &gst.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &gst.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &gst.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &gst.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &gst.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &gst.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &gst.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &gst.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &gst.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &gst.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &gst.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &gst.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &gst.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &gst.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &gst.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &gst.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &gst.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &gst.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &gst.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &gst.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &gst.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &gst.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &gst.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &gst.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &gst.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &gst.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &gst.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &gst.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &gst.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &gst.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &gst.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &gst.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &gst.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &gst.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	fname, ok := args[0].ToString()
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
	if len(bytes) > gst.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &gst.Bytes{Value: bytes}, nil
}

func osStat(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	fname, ok := args[0].ToString()
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
	fstat := &gst.ImmutableMap{
		Value: map[string]gst.Object{
			"name":  &gst.String{Value: stat.Name()},
			"mtime": &gst.Time{Value: stat.ModTime()},
			"size":  &gst.Int{Value: stat.Size()},
			"mode":  &gst.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = gst.TrueValue
	} else {
		fstat.Value["directory"] = gst.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...gst.Object) (gst.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...gst.Object) (gst.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...gst.Object) (gst.Object, error) {
	if len(args) != 3 {
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
	i2, ok := args[1].ToInt()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := args[2].ToInt()
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

func osArgs(args ...gst.Object) (gst.Object, error) {
	if len(args) != 0 {
		return nil, gse.ErrWrongNumArguments
	}
	arr := &gst.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > gst.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		arr.Value = append(arr.Value, &gst.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *gst.UserFunction {
	return &gst.UserFunction{
		Name: name,
		Value: func(args ...gst.Object) (gst.Object, error) {
			if len(args) != 2 {
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
			i2, ok := args[1].ToInt64()
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

func osLookupEnv(args ...gst.Object) (gst.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return gst.FalseValue, nil
	}
	if len(res) > gst.MaxStringLen {
		return nil, gse.ErrStringLimit
	}
	return &gst.String{Value: res}, nil
}

func osExpandEnv(args ...gst.Object) (gst.Object, error) {
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
		if vlen > gst.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > gst.MaxStringLen {
		return nil, gse.ErrStringLimit
	}
	return &gst.String{Value: s}, nil
}

func osExec(args ...gst.Object) (gst.Object, error) {
	if len(args) == 0 {
		return nil, gse.ErrWrongNumArguments
	}
	name, ok := args[0].ToString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := arg.ToString()
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

func osFindProcess(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	i1, ok := args[0].ToInt()
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

func osStartProcess(args ...gst.Object) (gst.Object, error) {
	if len(args) != 4 {
		return nil, gse.ErrWrongNumArguments
	}
	name, ok := args[0].ToString()
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
	case *gst.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *gst.ImmutableArray:
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

	dir, ok := args[2].ToString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *gst.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *gst.ImmutableArray:
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

func stringArray(arr []gst.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*gst.String)
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
