package core

import "time"

type Allocator interface {
	Release(Object)

	NewUndefined() Object
	NewBool(bool) Object
	NewInt(int64) Object
	NewFloat(float64) Object
	NewChar(rune) Object
	NewString(string) Object
	NewBytes([]byte) Object
	NewTime(time.Time) Object
	NewStringIterator([]rune) Iterator
	NewBytesIterator([]byte) Iterator
	NewMapIterator(map[string]Object) Iterator
	NewArrayIterator([]Object) Iterator
	NewError(Object) Object
	NewMap(val map[string]Object, immutable bool) Object
	NewRecord(val map[string]Object, immutable bool) Object
	NewArray(val []Object, immutable bool) Object
	NewBuiltinFunction(name string, val NativeFunc, arity int, variadic bool) Object
}
