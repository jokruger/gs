package value

import "github.com/jokruger/gs/core"

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

func (i *StringIterator) Equals(core.Object) bool {
	return false
}

func (i *StringIterator) Copy() core.Object {
	return &StringIterator{v: i.v, i: i.i, l: i.l}
}

func (i *StringIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *StringIterator) Key() core.Object {
	return &Int{Value: int64(i.i - 1)}
}

func (i *StringIterator) Value() core.Object {
	return &Char{Value: i.v[i.i-1]}
}

func (o *StringIterator) AsBool() (bool, bool) {
	return !o.IsFalsy(), true
}
