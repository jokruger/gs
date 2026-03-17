package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type BytesIterator struct {
	v []byte
	i int
	l int
}

func NewBytesIterator(v []byte) *BytesIterator {
	o := &BytesIterator{}
	o.Set(v)
	return o
}

func (o *BytesIterator) Set(v []byte) {
	o.v = v
	o.i = 0
	o.l = len(v)
}

func (o *BytesIterator) Next() bool {
	o.i++
	return o.i <= o.l
}

func (o *BytesIterator) Key() core.Object {
	return NewInt(int64(o.i - 1))
}

func (o *BytesIterator) Value() core.Object {
	return NewInt(int64(o.v[o.i-1]))
}

func (o *BytesIterator) TypeName() string {
	return "bytes-iterator"
}

func (o *BytesIterator) String() string {
	return "<bytes-iterator>"
}

func (o *BytesIterator) Interface() any {
	return o
}

func (o *BytesIterator) Arity() int {
	return 0
}

func (o *BytesIterator) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *BytesIterator) Equals(core.Object) bool {
	return false
}

func (o *BytesIterator) Copy() core.Object {
	t := NewBytesIterator(o.v)
	t.i = o.i
	return t
}

func (o *BytesIterator) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *BytesIterator) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *BytesIterator) Iterate() core.Iterator {
	return nil
}

func (o *BytesIterator) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *BytesIterator) IsFalsy() bool {
	return o == nil
}

func (o *BytesIterator) IsIterable() bool {
	return false
}

func (o *BytesIterator) IsCallable() bool {
	return false
}

func (o *BytesIterator) IsImmutable() bool {
	return false
}

func (o *BytesIterator) IsVariadic() bool {
	return false
}

func (o *BytesIterator) AsString() (string, bool) {
	return "", false
}

func (o *BytesIterator) AsInt() (int64, bool) {
	return 0, false
}

func (o *BytesIterator) AsFloat() (float64, bool) {
	return 0, false
}

func (o *BytesIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}

func (o *BytesIterator) AsRune() (rune, bool) {
	return 0, false
}

func (o *BytesIterator) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *BytesIterator) AsTime() (time.Time, bool) {
	return time.Time{}, false
}
