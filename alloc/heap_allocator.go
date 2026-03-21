package alloc

import (
	"time"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

type HeapAllocator struct {
	trueValue      core.Object
	falseValue     core.Object
	undefinedValue core.Object
}

func NewHeapAllocator() core.Allocator {
	return &HeapAllocator{
		trueValue:      value.NewStaticBool(true),
		falseValue:     value.NewStaticBool(false),
		undefinedValue: &value.Undefined{},
	}
}

func (a *HeapAllocator) Release(o core.Object) {
	// No-op, GC will take care of it
}

func (a *HeapAllocator) NewUndefined() core.Object {
	return a.undefinedValue
}

func (a *HeapAllocator) NewBool(v bool) core.Object {
	if v {
		return a.trueValue
	}
	return a.falseValue
}

func (a *HeapAllocator) NewInt(v int64) core.Object {
	o := &value.Int{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewFloat(v float64) core.Object {
	o := &value.Float{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewChar(v rune) core.Object {
	o := &value.Char{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewString(v string) core.Object {
	o := &value.String{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewBytes(v []byte) core.Object {
	o := &value.Bytes{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewTime(v time.Time) core.Object {
	o := &value.Time{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewStringIterator(v []rune) core.Iterator {
	o := &value.StringIterator{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewBytesIterator(v []byte) core.Iterator {
	o := &value.BytesIterator{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewMapIterator(v map[string]core.Object) core.Iterator {
	o := &value.MapIterator{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewArrayIterator(v []core.Object) core.Iterator {
	o := &value.ArrayIterator{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewError(v core.Object) core.Object {
	o := &value.Error{}
	o.Set(v)
	return o
}

func (a *HeapAllocator) NewMap(val map[string]core.Object, immutable bool) core.Object {
	o := &value.Map{}
	o.Set(val, immutable)
	return o
}

func (a *HeapAllocator) NewRecord(val map[string]core.Object, immutable bool) core.Object {
	o := &value.Record{}
	o.Set(val, immutable)
	return o
}

func (a *HeapAllocator) NewArray(val []core.Object, immutable bool) core.Object {
	o := &value.Array{}
	o.Set(val, immutable)
	return o
}

func (a *HeapAllocator) NewBuiltinFunction(name string, val core.NativeFunc, arity int, variadic bool) core.Object {
	o := &value.BuiltinFunction{}
	o.Set(name, val, arity, variadic)
	return o
}
