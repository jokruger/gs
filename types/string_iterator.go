package types

type StringIterator struct {
	ObjectImpl
	v []rune
	i int
	l int
}

func (i *StringIterator) TypeName() string {
	return "string-iterator"
}

func (i *StringIterator) String() string {
	return "<string-iterator>"
}

func (i *StringIterator) IsFalsy() bool {
	return true
}

func (i *StringIterator) Equals(Object) bool {
	return false
}

func (i *StringIterator) Copy() Object {
	return &StringIterator{v: i.v, i: i.i, l: i.l}
}

func (i *StringIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *StringIterator) Key() Object {
	return &Int{Value: int64(i.i - 1)}
}

func (i *StringIterator) Value() Object {
	return &Char{Value: i.v[i.i-1]}
}

func (o *StringIterator) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
