package value

type ArrayIterator struct {
	ObjectImpl
	v []Object
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

func (i *ArrayIterator) Equals(Object) bool {
	return false
}

func (i *ArrayIterator) Copy() Object {
	return &ArrayIterator{v: i.v, i: i.i, l: i.l}
}

func (i *ArrayIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *ArrayIterator) Key() Object {
	return &Int{Value: int64(i.i - 1)}
}

func (i *ArrayIterator) Value() Object {
	return i.v[i.i-1]
}

func (o *ArrayIterator) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
