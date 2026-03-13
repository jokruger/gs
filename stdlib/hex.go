package stdlib

import (
	"encoding/hex"

	gst "github.com/jokruger/gs/types"
)

var hexModule = map[string]gst.Object{
	"encode": &gst.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &gst.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
