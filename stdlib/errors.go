package stdlib

import (
	"github.com/jokruger/gs/core"
)

func wrapError(alloc core.Allocator, err error) core.Object {
	if err == nil {
		return alloc.NewBool(true)
	}
	return alloc.NewError(alloc.NewString(err.Error()))
}
