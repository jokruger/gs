package value

import (
	"time"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type MapIterator struct {
	v map[string]core.Object
	k []string
	i int
	l int
}

func NewMapIterator(m map[string]core.Object) *MapIterator {
	o := &MapIterator{}
	o.Set(m)
	return o
}

func (o *MapIterator) Set(m map[string]core.Object) {
	o.v = m
	o.k = make([]string, 0, len(m))
	for k := range m {
		o.k = append(o.k, k)
	}
	o.i = 0
	o.l = len(o.k)
}

func (o *MapIterator) Next() bool {
	o.i++
	return o.i <= o.l
}

func (o *MapIterator) Key() core.Object {
	k := o.k[o.i-1]
	return NewString(k)
}

func (o *MapIterator) Value() core.Object {
	k := o.k[o.i-1]
	return o.v[k]
}

func (o *MapIterator) TypeName() string {
	return "map-iterator"
}

func (o *MapIterator) String() string {
	return "<map-iterator>"
}

func (o *MapIterator) Interface() any {
	return o
}

func (o *MapIterator) Arity() int {
	return 0
}

func (o *MapIterator) BinaryOp(token.Token, core.Object) (core.Object, error) {
	return nil, gse.ErrInvalidOperator
}

func (o *MapIterator) Equals(core.Object) bool {
	return false
}

func (o *MapIterator) Copy() core.Object {
	t := NewMapIterator(o.v)
	t.i = o.i
	return t
}

func (o *MapIterator) IndexGet(core.Object) (core.Object, error) {
	return nil, gse.ErrNotIndexable
}

func (o *MapIterator) IndexSet(core.Object, core.Object) error {
	return gse.ErrNotIndexAssignable
}

func (o *MapIterator) Iterate() core.Iterator {
	return nil
}

func (o *MapIterator) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *MapIterator) IsFalsy() bool {
	return true
}

func (o *MapIterator) IsIterable() bool {
	return false
}

func (o *MapIterator) IsCallable() bool {
	return false
}

func (o *MapIterator) IsImmutable() bool {
	return false
}

func (o *MapIterator) IsVariadic() bool {
	return false
}

func (o *MapIterator) AsString() (string, bool) {
	return "", false
}

func (o *MapIterator) AsInt() (int64, bool) {
	return 0, false
}

func (o *MapIterator) AsFloat() (float64, bool) {
	return 0, false
}

func (o *MapIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}

func (o *MapIterator) AsRune() (rune, bool) {
	return 0, false
}

func (o *MapIterator) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *MapIterator) AsTime() (time.Time, bool) {
	return time.Time{}, false
}
