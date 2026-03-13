package types

type MapIterator struct {
	ObjectImpl
	v map[string]Object
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

func (i *MapIterator) Equals(Object) bool {
	return false
}

func (i *MapIterator) Copy() Object {
	return &MapIterator{v: i.v, k: i.k, i: i.i, l: i.l}
}

func (i *MapIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *MapIterator) Key() Object {
	k := i.k[i.i-1]
	return &String{Value: k}
}

func (i *MapIterator) Value() Object {
	k := i.k[i.i-1]
	return i.v[k]
}

func (o *MapIterator) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
