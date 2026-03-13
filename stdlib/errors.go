package stdlib

import (
	gst "github.com/jokruger/gs/types"
)

func wrapError(err error) gst.Object {
	if err == nil {
		return gst.TrueValue
	}
	return &gst.Error{Value: &gst.String{Value: err.Error()}}
}
