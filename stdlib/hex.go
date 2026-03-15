package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

var hexModule = map[string]core.Object{
	"encode": &value.BuiltinFunction{Value: hexEncodeToString},
	"decode": &value.BuiltinFunction{Value: FuncASRYE(hex.DecodeString)},
}

func hexEncodeToString(args ...core.Object) (ret core.Object, err error) {
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
	res := hex.EncodeToString(y1)
	return &value.String{Value: res}, nil
}
