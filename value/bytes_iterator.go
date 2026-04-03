package value

import "github.com/jokruger/gs/core"

type BytesIterator struct {
	Object
	v []byte
	i int
	l int
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

func (o *BytesIterator) Key(core.Allocator) core.Value {
	return core.NewInt(int64(o.i - 1))
}

func (o *BytesIterator) Value(core.Allocator) core.Value {
	return core.NewInt(int64(o.v[o.i-1]))
}

func (o *BytesIterator) TypeName() string {
	return "bytes-iterator"
}

func (o *BytesIterator) String() string {
	return "<bytes-iterator>"
}

func (o *BytesIterator) Copy(alloc core.Allocator) core.Value {
	t := alloc.NewBytesIterator(o.v).(*BytesIterator)
	t.i = o.i
	return core.NewObject(t, false)
}

func (o *BytesIterator) IsTrue() bool {
	return o.v != nil && o.i <= o.l
}

func (o *BytesIterator) IsFalse() bool {
	return !o.IsTrue()
}

func (o *BytesIterator) AsBool() (bool, bool) {
	return o.IsTrue(), true
}
