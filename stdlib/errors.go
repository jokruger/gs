package stdlib

import (
	"github.com/jokruger/gs"
)

func wrapError(err error) gs.Object {
	if err == nil {
		return gs.TrueValue
	}
	return &gs.Error{Value: &gs.String{Value: err.Error()}}
}
