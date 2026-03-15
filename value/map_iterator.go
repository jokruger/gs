package value

import "github.com/jokruger/gs/core"

type MapIterator struct {
	Object
	v map[string]core.Object
	k []string
	i int
	l int
}

func (i *MapIterator) TypeName() string {
	return "map-iterator"
}

func (i *MapIterator) String() string {
	return "<map-iterator>"
}

func (i *MapIterator) IsFalsy() bool {
	return true
}

func (i *MapIterator) Equals(core.Object) bool {
	return false
}

func (i *MapIterator) Copy() core.Object {
	return &MapIterator{v: i.v, k: i.k, i: i.i, l: i.l}
}

func (i *MapIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *MapIterator) Key() core.Object {
	k := i.k[i.i-1]
	return &String{Value: k}
}

func (i *MapIterator) Value() core.Object {
	k := i.k[i.i-1]
	return i.v[k]
}

func (o *MapIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}
