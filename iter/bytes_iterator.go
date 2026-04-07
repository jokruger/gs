package iter

import (
	"fmt"

	"github.com/jokruger/gs/core"
)

type BytesIterator struct {
	v []byte
	i int
	l int
}

func (i *BytesIterator) Set(v []byte) {
	i.v = v
	i.i = 0
	i.l = len(v)
}

func (i *BytesIterator) TypeName() string {
	return "bytes-iterator"
}

func (i *BytesIterator) String() string {
	return fmt.Sprintf("BytesIterator{%d/%d}", i.i, i.l)
}

func (i *BytesIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *BytesIterator) Key(core.Allocator) core.Value {
	return core.IntValue(int64(i.i - 1))
}

func (i *BytesIterator) Value(core.Allocator) core.Value {
	return core.IntValue(int64(i.v[i.i-1]))
}
