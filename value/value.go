package value

import "github.com/jokruger/gs/core"

const (
	// TrueString is a string representation of the boolean value true.
	TrueString = "true"

	// FalseString is a string representation of the boolean value false.
	FalseString = "false"
)

var (
	// TrueValue is the singleton instance representing the boolean value true.
	TrueValue core.Object = &Bool{value: true}

	// FalseValue is the singleton instance representing the boolean value false.
	FalseValue core.Object = &Bool{value: false}

	// UndefinedValue is the singleton instance representing the undefined value.
	UndefinedValue core.Object = &Undefined{}
)
