package core

import (
	"fmt"
	"unsafe"
)

type RunesIterator struct {
	v []rune
	i int
}

func (i *RunesIterator) Set(v []rune) {
	i.v = v
	i.i = -1
}

func RunesIteratorValue(v *RunesIterator) Value {
	return Value{
		Ptr:  unsafe.Pointer(v),
		Type: VT_RUNES_ITERATOR,
	}
}

func NewRunesIteratorValue(v []rune) Value {
	i := &RunesIterator{}
	i.Set(v)
	return RunesIteratorValue(i)
}

func runesIteratorTypeName(v Value) string {
	return "runes-iterator"
}

func runesIteratorTypeString(v Value) string {
	i := (*RunesIterator)(v.Ptr)
	return fmt.Sprintf("RunesIterator{%d, %d}", i.i, len(i.v))
}

func runesIteratorTypeEqual(v Value, r Value) bool {
	if r.Type != VT_RUNES_ITERATOR {
		return false
	}
	a := (*RunesIterator)(v.Ptr)
	b := (*RunesIterator)(r.Ptr)
	return a == b
}

func runesIteratorTypeNext(v Value) bool {
	i := (*RunesIterator)(v.Ptr)
	i.i++
	return i.i < len(i.v)
}

func runesIteratorTypeKey(v Value, alloc Allocator) (Value, error) {
	i := (*RunesIterator)(v.Ptr)
	return IntValue(int64(i.i)), nil
}

func runesIteratorTypeValue(v Value, alloc Allocator) (Value, error) {
	i := (*RunesIterator)(v.Ptr)
	return RuneValue(i.v[i.i]), nil
}
