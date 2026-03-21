package stdlib

import (
	"encoding/base64"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var base64Module = map[string]core.Object{
	"encode":         value.NewStaticBuiltinFunction("encode", b64EncodeToString, 1, false),
	"decode":         value.NewStaticBuiltinFunction("decode", b64DecodeString, 1, false),
	"raw_encode":     value.NewStaticBuiltinFunction("raw_encode", b64RawEncodeToString, 1, false),
	"raw_decode":     value.NewStaticBuiltinFunction("raw_decode", b64RawDecodeString, 1, false),
	"url_encode":     value.NewStaticBuiltinFunction("url_encode", b64URLEncodeToString, 1, false),
	"url_decode":     value.NewStaticBuiltinFunction("url_decode", b64URLDecodeString, 1, false),
	"raw_url_encode": value.NewStaticBuiltinFunction("raw_url_encode", b64RawURLEncodeToString, 1, false),
	"raw_url_decode": value.NewStaticBuiltinFunction("raw_url_decode", b64RawURLDecodeString, 1, false),
}

func b64RawURLDecodeString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.raw_url_decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.raw_url_decode", "first", "string(compatible)", args[0])
	}
	res, err := base64.RawURLEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(alloc, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, core.NewBytesLimitError("base64.raw_url_decode")
	}
	return alloc.NewBytes(res), nil
}

func b64URLDecodeString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.url_decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.url_decode", "first", "string(compatible)", args[0])
	}
	res, err := base64.URLEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(alloc, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, core.NewBytesLimitError("base64.url_decode")
	}
	return alloc.NewBytes(res), nil
}

func b64RawDecodeString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.raw_decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.raw_decode", "first", "string(compatible)", args[0])
	}
	res, err := base64.RawStdEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(alloc, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, core.NewBytesLimitError("base64.raw_decode")
	}
	return alloc.NewBytes(res), nil
}

func b64DecodeString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.decode", "first", "string(compatible)", args[0])
	}
	res, err := base64.StdEncoding.DecodeString(s1)
	if err != nil {
		return wrapError(alloc, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, core.NewBytesLimitError("base64.decode")
	}
	return alloc.NewBytes(res), nil
}

func b64RawURLEncodeToString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.raw_url_encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.raw_url_encode", "first", "bytes(compatible)", args[0])
	}
	res := base64.RawURLEncoding.EncodeToString(y1)
	return alloc.NewString(res), nil
}

func b64URLEncodeToString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.url_encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.url_encode", "first", "bytes(compatible)", args[0])
	}
	res := base64.URLEncoding.EncodeToString(y1)
	return alloc.NewString(res), nil
}

func b64RawEncodeToString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.raw_encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.raw_encode", "first", "bytes(compatible)", args[0])
	}
	res := base64.RawStdEncoding.EncodeToString(y1)
	return alloc.NewString(res), nil
}

func b64EncodeToString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("base64.encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("base64.encode", "first", "bytes(compatible)", args[0])
	}
	res := base64.StdEncoding.EncodeToString(y1)
	return alloc.NewString(res), nil
}
