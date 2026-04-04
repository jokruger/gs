package core

import "time"

type Allocator interface {
	ReleaseObject(Object)
	ReleaseIterator(Iterator)

	NewString(string) Object
	NewBytes([]byte) Object
	NewTime(time.Time) Object
	NewStringIterator([]rune) Iterator
	NewBytesIterator([]byte) Iterator
	NewMapIterator(map[string]Value) Iterator
	NewArrayIterator([]Value) Iterator
	NewError(Value) Object
	NewMap(val map[string]Value, immutable bool) Object
	NewRecord(val map[string]Value, immutable bool) Object
	NewArray(val []Value, immutable bool) Object
	NewBuiltinFunction(name string, val NativeFunc, arity int, variadic bool) Object

	NewStringValue(string) Value
	NewBytesValue([]byte) Value
	NewTimeValue(time.Time) Value
	NewErrorValue(Value) Value
	NewMapValue(val map[string]Value, immutable bool) Value
	NewRecordValue(val map[string]Value, immutable bool) Value
	NewArrayValue(val []Value, immutable bool) Value
	NewBuiltinFunctionValue(name string, val NativeFunc, arity int, variadic bool) Value
}
