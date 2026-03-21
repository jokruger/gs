package core

// Iterator represents an iterator for underlying data type.
type Iterator interface {
	Object
	Next() bool             // returns true if there are more elements to iterate
	Key(Allocator) Object   // returns the key or index value of the current element
	Value(Allocator) Object // returns the value of the current element
}
