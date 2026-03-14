package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type ObjectImpl struct {
}

func (o *ObjectImpl) TypeName() string {
	panic(gse.ErrNotImplemented)
}

func (o *ObjectImpl) String() string {
	panic(gse.ErrNotImplemented)
}

func (o *ObjectImpl) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *ObjectImpl) Copy() core.Object {
	return nil
}

func (o *ObjectImpl) IsFalsy() bool {
	return false
}

func (o *ObjectImpl) Equals(x core.Object) bool {
	return o == x
}

func (o *ObjectImpl) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *ObjectImpl) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *ObjectImpl) Iterate() core.Iterator {
	return nil
}

func (o *ObjectImpl) IsIterable() bool {
	return false
}

func (o *ObjectImpl) Call(...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *ObjectImpl) IsCallable() bool {
	return false
}

func (o *ObjectImpl) AsString() (string, bool) {
	return "", false
}

func (o *ObjectImpl) AsInt() (int64, bool) {
	return 0, false
}

func (o *ObjectImpl) AsFloat() (float64, bool) {
	return 0, false
}

func (o *ObjectImpl) AsBool() (bool, bool) {
	return false, false
}

func (o *ObjectImpl) AsRune() (rune, bool) {
	return 0, false
}

func (o *ObjectImpl) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *ObjectImpl) AsTime() (time.Time, bool) {
	return time.Time{}, false
}

func (o *ObjectImpl) Interface() any {
	return o
}
