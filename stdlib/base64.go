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
		Value: b64DecodeString,
	},
	"raw_encode": &value.BuiltinFunction{
		Value: b64RawEncodeToString,
	},
	"raw_decode": &value.BuiltinFunction{
		Value: b64RawDecodeString,
	},
	"url_encode": &value.BuiltinFunction{
		Value: b64URLEncodeToString,
	},
	"url_decode": &value.BuiltinFunction{
		Value: b64URLDecodeString,
	},
	"raw_url_encode": &value.BuiltinFunction{
		Value: b64RawURLEncodeToString,
	},
	"raw_url_decode": &value.BuiltinFunction{
		Value: b64RawURLDecodeString,
	},
}

func b64RawURLDecodeString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := base64.RawURLEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &value.Bytes{Value: res}, nil
}

func b64URLDecodeString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := base64.URLEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &value.Bytes{Value: res}, nil
}

func b64RawDecodeString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := base64.RawStdEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &value.Bytes{Value: res}, nil
}

func b64DecodeString(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := base64.StdEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, gse.ErrBytesLimit
	}
	return &value.Bytes{Value: res}, nil
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
