package stdlib

import (
	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

func wrapError(err error) core.Object {
	if err == nil {
		return value.TrueValue
	}
	return &value.Error{Value: &value.String{Value: err.Error()}}
}
