package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var hexModule = map[string]core.Value{
	"encode": value.NewStaticBuiltinFunction("encode", hexEncodeToString, 1, false),
	"decode": value.NewStaticBuiltinFunction("decode", hexDecodeString, 1, false),
}

func hexDecodeString(vm core.VM, args ...core.Value) (ret core.Value, err error) {
	if len(args) != 1 {
		return core.NewUndefined(), core.NewWrongNumArgumentsError("hex.decode", "1", len(args))
	}
	s1, ok := args[0].AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidArgumentTypeError("hex.decode", "first", "string(compatible)", args[0].TypeName())
	}
	res, err := hex.DecodeString(s1)
	if err != nil {
		return wrapError(vm, err), nil
	}
	if len(res) > core.MaxBytesLen {
		return core.NewUndefined(), core.NewBytesLimitError("hex.decode")
	}
	return vm.Allocator().NewBytesValue(res), nil
}

func hexEncodeToString(vm core.VM, args ...core.Value) (ret core.Value, err error) {
	if len(args) != 1 {
		return core.NewUndefined(), core.NewWrongNumArgumentsError("hex.encode", "1", len(args))
	}
	y1, ok := args[0].AsBytes()
	if !ok {
		return core.NewUndefined(), core.NewInvalidArgumentTypeError("hex.encode", "first", "bytes(compatible)", args[0].TypeName())
	}
	res := hex.EncodeToString(y1)
	return vm.Allocator().NewStringValue(res), nil
}
