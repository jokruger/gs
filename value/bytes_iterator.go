package value

type BytesIterator struct {
	ObjectImpl
	v []byte
	i int
	l int
}

func (i *BytesIterator) TypeName() string {
	return "bytes-iterator"
}

func (i *BytesIterator) String() string {
	return "<bytes-iterator>"
}

func (i *BytesIterator) Equals(Object) bool {
	return false
}

func (i *BytesIterator) Copy() Object {
	return &BytesIterator{v: i.v, i: i.i, l: i.l}
}

func (i *BytesIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *BytesIterator) Key() Object {
	return &Int{Value: int64(i.i - 1)}
}

func (i *BytesIterator) Value() Object {
	return &Int{Value: int64(i.v[i.i-1])}
}

func (o *BytesIterator) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
