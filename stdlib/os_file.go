package stdlib

import (
	"os"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

func makeOSFile(file *os.File) *gst.ImmutableMap {
	return &gst.ImmutableMap{
		Value: map[string]gst.Object{
			// chdir() => true/error
			"chdir": &gst.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &gst.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &gst.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &gst.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &gst.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &gst.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &gst.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &gst.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &gst.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &gst.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &gst.UserFunction{
				Name: "seek",
				Value: func(args ...gst.Object) (gst.Object, error) {
					if len(args) != 2 {
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
					i2, ok := args[1].ToInt()
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &gst.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &gst.UserFunction{
				Name: "stat",
				Value: func(args ...gst.Object) (gst.Object, error) {
					if len(args) != 0 {
						return nil, gse.ErrWrongNumArguments
					}
					return osStat(&gst.String{Value: file.Name()})
				},
			},
		},
	}
}
