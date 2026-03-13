package stdlib

import (
	"encoding/hex"

	"github.com/jokruger/gs"
)

var hexModule = map[string]gs.Object{
	"encode": &gs.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &gs.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
