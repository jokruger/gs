package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var hexModule = map[string]core.Object{
	"encode": value.NewStaticBuiltinFunction("encode", hexEncodeToString, 1, false),
	"decode": value.NewStaticBuiltinFunction("decode", hexDecodeString, 1, false),
}

func hexDecodeString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("hex.decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("hex.decode", "first", "string(compatible)", args[0])
	}
	res, err := hex.DecodeString(s1)
	if err != nil {
		return wrapError(alloc, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return nil, core.NewBytesLimitError("hex.decode")
	}
	return alloc.NewBytes(res), nil
}

func hexEncodeToString(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("hex.encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("hex.encode", "first", "bytes(compatible)", args[0])
	}
	res := hex.EncodeToString(y1)
	return alloc.NewString(res), nil
}
