package core

import (
	"time"

	"github.com/jokruger/gs/token"
)

// Object represents an object in the VM.
type Object interface {
	TypeName() string // return the name of the type
	String() string   // return the string representation of the type's value
	Interface() any   // return the value of the type as an empty interface
	Arity() int       // return the number of positional arguments (or minimum number of arguments if variadic) for callable objects

	BinaryOp(token.Token, Object) (Object, error) // return the result of a binary operation with another object
	Equals(Object) bool                           // return whether the value of the type is equal to the value of another object
	Copy() Object                                 // return a copy of the type (and its value)
	IndexGet(Object) (Object, error)              // return the result of indexing the object with the given index
	IndexSet(idx, val Object) error               // return the result of setting the value of the object at the given index
	Iterate() Iterator                            // return an Iterator for the type
	Call(VM, ...Object) (Object, error)           // return the result of calling the object with the given arguments

	IsFalsy() bool     // return whether the value of the type is equivalent to false in a boolean context
	IsIterable() bool  // return whether the type is iterable (i.e. can be used in a for loop)
	IsCallable() bool  // return whether the type is callable (i.e. can be called like a function)
	IsImmutable() bool // return whether the type is immutable (i.e. cannot be modified after creation)
	IsVariadic() bool  // return whether the callable type is variadic (i.e. can accept a variable number of arguments)

	AsString() (string, bool)    // return the string value and whether the conversion was successful
	AsInt() (int64, bool)        // return the int value and whether the conversion was successful
	AsFloat() (float64, bool)    // return the float value and whether the conversion was successful
	AsBool() (bool, bool)        // return the bool value and whether the conversion was successful
	AsRune() (rune, bool)        // return the rune value and whether the conversion was successful
	AsByteSlice() ([]byte, bool) // return the byte slice value and whether the conversion was successful
	AsTime() (time.Time, bool)   // return the time value and whether the conversion was successful
}
