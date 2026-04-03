package value

import "github.com/jokruger/gs/core"

type ArrayIterator struct {
	Object
	v []core.Value
	i int
	l int
}

func (o *ArrayIterator) Set(v []core.Value) {
	o.v = v
	o.i = 0
	o.l = len(v)
}

func (o *ArrayIterator) Next() bool {
	o.i++
	return o.i <= o.l
}

func (o *ArrayIterator) Key(core.Allocator) core.Value {
	return core.NewInt(int64(o.i - 1))
}

func (o *ArrayIterator) Value(alloc core.Allocator) core.Value {
	return o.v[o.i-1].Copy(alloc)
}

func (o *ArrayIterator) TypeName() string {
	return "array-iterator"
}

func (o *ArrayIterator) String() string {
	return "<array-iterator>"
}

func (o *ArrayIterator) Copy(alloc core.Allocator) core.Value {
	t := alloc.NewArrayIterator(o.v).(*ArrayIterator)
	t.i = o.i
	return core.NewObject(t, false)
}

func (o *ArrayIterator) IsTrue() bool {
	return o.v != nil && o.i <= o.l
}

func (o *ArrayIterator) IsFalse() bool {
	return !o.IsTrue()
}

func (o *ArrayIterator) AsBool() (bool, bool) {
	return o.IsTrue(), true
}
