package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type Object struct {
}

func (o *Object) TypeName() string {
	panic(gse.ErrNotImplemented)
}

func (o *Object) String() string {
	panic(gse.ErrNotImplemented)
}

func (o *Object) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *Object) Copy() core.Object {
	return o
}

func (o *Object) IsFalsy() bool {
	return o == nil
}

func (o *Object) Equals(x core.Object) bool {
	return o == x
}

func (o *Object) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *Object) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *Object) Iterate() core.Iterator {
	return nil
}

func (o *Object) IsIterable() bool {
	return false
}

func (o *Object) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *Object) IsCallable() bool {
	return false
}

func (o *Object) AsString() (string, bool) {
	return "", false
}

func (o *Object) AsInt() (int64, bool) {
	return 0, false
}

func (o *Object) AsFloat() (float64, bool) {
	return 0, false
}

func (o *Object) AsBool() (bool, bool) {
	return false, false
}

func (o *Object) AsRune() (rune, bool) {
	return 0, false
}

func (o *Object) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *Object) AsTime() (time.Time, bool) {
	return time.Time{}, false
}

func (o *Object) Interface() any {
	return o
}
