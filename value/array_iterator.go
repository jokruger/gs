package value

import "github.com/jokruger/gs/core"

type ArrayIterator struct {
	ObjectImpl
	v []core.Object
	i int
	l int
}

func (i *ArrayIterator) TypeName() string {
	return "array-iterator"
}

func (i *ArrayIterator) String() string {
	return "<array-iterator>"
}

func (i *ArrayIterator) IsFalsy() bool {
	return true
}

func (i *ArrayIterator) Equals(core.Object) bool {
	return false
}

func (i *ArrayIterator) Copy() core.Object {
	return &ArrayIterator{v: i.v, i: i.i, l: i.l}
}

func (i *ArrayIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *ArrayIterator) Key() core.Object {
	return &Int{Value: int64(i.i - 1)}
}

func (i *ArrayIterator) Value() core.Object {
	return i.v[i.i-1]
}

func (o *ArrayIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}
