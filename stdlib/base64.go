package stdlib

import (
	"encoding/base64"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

var base64Module = map[string]core.Object{
	"encode": &value.BuiltinFunction{
		Value: b64EncodeToString,
	},
	"decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &value.BuiltinFunction{
		Value: b64RawEncodeToString,
	},
	"raw_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &value.BuiltinFunction{
		Value: b64URLEncodeToString,
	},
	"url_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &value.BuiltinFunction{
		Value: b64RawURLEncodeToString,
	},
	"raw_url_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}

func b64RawURLEncodeToString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	y1, ok := args[0].AsByteSlice()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res := base64.RawURLEncoding.EncodeToString(y1)
	return &value.String{Value: res}, nil
}

func b64URLEncodeToString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	y1, ok := args[0].AsByteSlice()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res := base64.URLEncoding.EncodeToString(y1)
	return &value.String{Value: res}, nil
}

func b64RawEncodeToString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	y1, ok := args[0].AsByteSlice()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res := base64.RawStdEncoding.EncodeToString(y1)
	return &value.String{Value: res}, nil
}

func b64EncodeToString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	y1, ok := args[0].AsByteSlice()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res := base64.StdEncoding.EncodeToString(y1)
	return &value.String{Value: res}, nil
}
