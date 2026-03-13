package stdlib

import (
	"os"

	"github.com/jokruger/gs"
)

func makeOSFile(file *os.File) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			// chdir() => true/error
			"chdir": &gs.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &gs.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &gs.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &gs.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &gs.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &gs.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &gs.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &gs.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &gs.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &gs.UserFunction{
				Name: "chmod",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 1 {
						return nil, gs.ErrWrongNumArguments
					}
					i1, ok := gs.ToInt64(args[0])
					if !ok {
						return nil, gs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &gs.UserFunction{
				Name: "seek",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 2 {
						return nil, gs.ErrWrongNumArguments
					}
					i1, ok := gs.ToInt64(args[0])
					if !ok {
						return nil, gs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					i2, ok := gs.ToInt(args[1])
					if !ok {
						return nil, gs.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &gs.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &gs.UserFunction{
				Name: "stat",
				Value: func(args ...gs.Object) (gs.Object, error) {
					if len(args) != 0 {
						return nil, gs.ErrWrongNumArguments
					}
					return osStat(&gs.String{Value: file.Name()})
				},
			},
		},
	}
}
