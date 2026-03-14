package stdlib

import (
	"os"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

func makeOSFile(file *os.File) *value.ImmutableMap {
	return &value.ImmutableMap{
		Value: map[string]core.Object{
			// chdir() => true/error
			"chdir": &value.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &value.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &value.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &value.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &value.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &value.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &value.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &value.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &value.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &value.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &value.UserFunction{
				Name: "seek",
				Value: func(args ...core.Object) (core.Object, error) {
					if len(args) != 2 {
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
					i2, ok := args[1].AsInt()
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, int(i2))
					if err != nil {
						return wrapError(err), nil
					}
					return &value.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &value.UserFunction{
				Name: "stat",
				Value: func(args ...core.Object) (core.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					return osStat(&value.String{Value: file.Name()})
				},
			},
		},
	}
}
