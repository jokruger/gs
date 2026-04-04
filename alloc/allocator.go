package alloc

import (
	"time"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/iter"
	"github.com/jokruger/gs/value"
)

type Allocator struct {
}

func New() core.Allocator {
	return &Allocator{}
}

func (a *Allocator) ReleaseObject(o core.Object) {
	// No-op, GC will take care of it
}

func (a *Allocator) ReleaseIterator(i core.Iterator) {
	// No-op, GC will take care of it
}

func (a *Allocator) NewString(v string) core.Object {
	o := &value.String{}
	o.Set(v)
	return o
}

func (a *Allocator) NewBytes(v []byte) core.Object {
	o := &value.Bytes{}
	o.Set(v)
	return o
}

func (a *Allocator) NewTime(v time.Time) core.Object {
	o := &value.Time{}
	o.Set(v)
	return o
}

func (a *Allocator) NewStringIterator(v []rune) core.Iterator {
	o := &iter.StringIterator{}
	o.Set(v)
	return o
}

func (a *Allocator) NewBytesIterator(v []byte) core.Iterator {
	o := &iter.BytesIterator{}
	o.Set(v)
	return o
}

func (a *Allocator) NewMapIterator(v map[string]core.Value) core.Iterator {
	o := &iter.MapIterator{}
	o.Set(v)
	return o
}

func (a *Allocator) NewArrayIterator(v []core.Value) core.Iterator {
	o := &iter.ArrayIterator{}
	o.Set(v)
	return o
}

func (a *Allocator) NewError(v core.Value) core.Object {
	o := &value.Error{}
	o.Set(v)
	return o
}

func (a *Allocator) NewMap(val map[string]core.Value, immutable bool) core.Object {
	o := &value.Map{}
	o.Set(val, immutable)
	return o
}

func (a *Allocator) NewRecord(val map[string]core.Value, immutable bool) core.Object {
	o := &value.Record{}
	o.Set(val, immutable)
	return o
}

func (a *Allocator) NewArray(val []core.Value, immutable bool) core.Object {
	o := &value.Array{}
	o.Set(val, immutable)
	return o
}

func (a *Allocator) NewBuiltinFunction(name string, val core.NativeFunc, arity int, variadic bool) core.Object {
	o := &value.BuiltinFunction{}
	o.Set(name, val, arity, variadic)
	return o
}

func (a *Allocator) NewStringValue(v string) core.Value {
	return core.NewObject(a.NewString(v), false)
}

func (a *Allocator) NewBytesValue(v []byte) core.Value {
	return core.NewObject(a.NewBytes(v), false)
}

func (a *Allocator) NewTimeValue(v time.Time) core.Value {
	return core.NewObject(a.NewTime(v), false)
}

func (a *Allocator) NewErrorValue(v core.Value) core.Value {
	return core.NewObject(a.NewError(v), false)
}

func (a *Allocator) NewMapValue(v map[string]core.Value, immutable bool) core.Value {
	return core.NewObject(a.NewMap(v, immutable), false)
}

func (a *Allocator) NewRecordValue(v map[string]core.Value, immutable bool) core.Value {
	return core.NewObject(a.NewRecord(v, immutable), false)
}

func (a *Allocator) NewArrayValue(v []core.Value, immutable bool) core.Value {
	return core.NewObject(a.NewArray(v, immutable), false)
}

func (a *Allocator) NewBuiltinFunctionValue(name string, val core.NativeFunc, arity int, variadic bool) core.Value {
	return core.NewObject(a.NewBuiltinFunction(name, val, arity, variadic), false)
}
