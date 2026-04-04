package core

// Iterator represents an iterator for underlying data type.
type Iterator interface {
	TypeName() string // returns the type name of the underlying data type
	String() string   // returns the string representation of the underlying data type

	Next() bool            // returns true if there are more elements to iterate
	Key(Allocator) Value   // returns the key or index value of the current element
	Value(Allocator) Value // returns the value of the current element
}
