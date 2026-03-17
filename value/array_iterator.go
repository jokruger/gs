package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type ArrayIterator struct {
	v []core.Object
	i int
	l int
}

func NewArrayIterator(v []core.Object) *ArrayIterator {
	o := &ArrayIterator{}
	o.Set(v)
	return o
}

func (o *ArrayIterator) Set(v []core.Object) {
	o.v = v
	o.i = 0
	o.l = len(v)
}

func (o *ArrayIterator) Next() bool {
	o.i++
	return o.i <= o.l
}

func (o *ArrayIterator) Key() core.Object {
	return NewInt(int64(o.i - 1))
}

func (o *ArrayIterator) Value() core.Object {
	return o.v[o.i-1]
}

func (o *ArrayIterator) TypeName() string {
	return "array-iterator"
}

func (o *ArrayIterator) String() string {
	return "<array-iterator>"
}

func (o *ArrayIterator) Interface() any {
	return o
}

func (o *ArrayIterator) Arity() int {
	return 0
}

func (o *ArrayIterator) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *ArrayIterator) Equals(core.Object) bool {
	return false
}

func (o *ArrayIterator) Copy() core.Object {
	t := NewArrayIterator(o.v)
	t.i = o.i
	return t
}

func (o *ArrayIterator) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *ArrayIterator) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *ArrayIterator) Iterate() core.Iterator {
	return nil
}

func (o *ArrayIterator) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *ArrayIterator) IsFalsy() bool {
	return true
}

func (o *ArrayIterator) IsIterable() bool {
	return false
}

func (o *ArrayIterator) IsCallable() bool {
	return false
}

func (o *ArrayIterator) IsImmutable() bool {
	return false
}

func (o *ArrayIterator) IsVariadic() bool {
	return false
}

func (o *ArrayIterator) AsString() (string, bool) {
	return "", false
}

func (o *ArrayIterator) AsInt() (int64, bool) {
	return 0, false
}

func (o *ArrayIterator) AsFloat() (float64, bool) {
	return 0, false
}

func (o *ArrayIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}

func (o *ArrayIterator) AsRune() (rune, bool) {
	return 0, false
}

func (o *ArrayIterator) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *ArrayIterator) AsTime() (time.Time, bool) {
	return time.Time{}, false
}
