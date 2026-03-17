package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type StringIterator struct {
	v []rune
	i int
	l int
}

func NewStringIterator(v []rune) *StringIterator {
	o := &StringIterator{}
	o.Set(v)
	return o
}

func (o *StringIterator) Set(v []rune) {
	o.v = v
	o.i = 0
	o.l = len(v)
}

func (o *StringIterator) Next() bool {
	o.i++
	return o.i <= o.l
}

func (o *StringIterator) Key() core.Object {
	return NewInt(int64(o.i - 1))
}

func (o *StringIterator) Value() core.Object {
	return NewChar(o.v[o.i-1])
}

func (o *StringIterator) TypeName() string {
	return "string-iterator"
}

func (o *StringIterator) String() string {
	return "<string-iterator>"
}

func (o *StringIterator) Interface() any {
	return o
}

func (o *StringIterator) Arity() int {
	return 0
}

func (o *StringIterator) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *StringIterator) Equals(core.Object) bool {
	return false
}

func (o *StringIterator) Copy() core.Object {
	t := NewStringIterator(o.v)
	t.i = o.i
	return t
}

func (o *StringIterator) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *StringIterator) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *StringIterator) Iterate() core.Iterator {
	return nil
}

func (o *StringIterator) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *StringIterator) IsFalsy() bool {
	return true
}

func (o *StringIterator) IsIterable() bool {
	return false
}

func (o *StringIterator) IsCallable() bool {
	return false
}

func (o *StringIterator) IsImmutable() bool {
	return false
}

func (o *StringIterator) IsVariadic() bool {
	return false
}

func (o *StringIterator) AsString() (string, bool) {
	return "", false
}

func (o *StringIterator) AsInt() (int64, bool) {
	return 0, false
}

func (o *StringIterator) AsFloat() (float64, bool) {
	return 0, false
}

func (o *StringIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}

func (o *StringIterator) AsRune() (rune, bool) {
	return 0, false
}

func (o *StringIterator) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *StringIterator) AsTime() (time.Time, bool) {
	return time.Time{}, false
}
