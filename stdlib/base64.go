package stdlib

import (
	"encoding/base64"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var base64Module = map[string]core.Object{
	"encode": &value.BuiltinFunction{
		Value: FuncAYRS(base64.StdEncoding.EncodeToString),
	},
	"decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &value.BuiltinFunction{
		Value: FuncAYRS(base64.RawStdEncoding.EncodeToString),
	},
	"raw_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &value.BuiltinFunction{
		Value: FuncAYRS(base64.URLEncoding.EncodeToString),
	},
	"url_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &value.BuiltinFunction{
		Value: FuncAYRS(base64.RawURLEncoding.EncodeToString),
	},
	"raw_url_decode": &value.BuiltinFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}
