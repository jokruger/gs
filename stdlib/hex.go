package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var hexModule = map[string]core.Object{
	"encode": &value.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &value.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
