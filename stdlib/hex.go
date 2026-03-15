package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var hexModule = map[string]core.Object{
	"encode": &value.BuiltinFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &value.BuiltinFunction{Value: FuncASRYE(hex.DecodeString)},
}
