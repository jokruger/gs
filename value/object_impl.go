package value

import (
	"time"

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

func (o *ObjectImpl) BinaryOp(token.Token, Object) (Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *ObjectImpl) Copy() Object {
	return nil
}

func (o *ObjectImpl) IsFalsy() bool {
	return false
}

func (o *ObjectImpl) Equals(x Object) bool {
	return o == x
}

func (o *ObjectImpl) IndexGet(Object) (res Object, err error) {
	return nil, gse.ErrNotIndexable
}

func (o *ObjectImpl) IndexSet(Object, Object) (err error) {
	return gse.ErrNotIndexAssignable
}

func (o *ObjectImpl) Iterate() Iterator {
	return nil
}

func (o *ObjectImpl) CanIterate() bool {
	return false
}

func (o *ObjectImpl) Call(...Object) (ret Object, err error) {
	return nil, nil
}

func (o *ObjectImpl) CanCall() bool {
	return false
}

func (o *ObjectImpl) ToString() (string, bool) {
	return "", false
}

func (o *ObjectImpl) ToInt() (int, bool) {
	return 0, false
}

func (o *ObjectImpl) ToInt64() (int64, bool) {
	return 0, false
}

func (o *ObjectImpl) ToFloat64() (float64, bool) {
	return 0, false
}

func (o *ObjectImpl) ToBool() (bool, bool) {
	return false, false
}

func (o *ObjectImpl) ToRune() (rune, bool) {
	return 0, false
}

func (o *ObjectImpl) ToByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *ObjectImpl) ToTime() (time.Time, bool) {
	return time.Time{}, false
}
