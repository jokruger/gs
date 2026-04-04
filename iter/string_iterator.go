package iter

import (
	"fmt"

	"github.com/jokruger/gs/core"
)

type StringIterator struct {
	v []rune
	i int
	l int
}

func (i *StringIterator) Set(v []rune) {
	i.v = v
	i.i = 0
	i.l = len(v)
}

func (i *StringIterator) TypeName() string {
	return "string-iterator"
}

func (i *StringIterator) String() string {
	return fmt.Sprintf("StringIterator{%d/%d}", i.i, i.l)
}

func (i *StringIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *StringIterator) Key(core.Allocator) core.Value {
	return core.NewInt(int64(i.i - 1))
}

func (i *StringIterator) Value(core.Allocator) core.Value {
	return core.NewChar(i.v[i.i-1])
}
