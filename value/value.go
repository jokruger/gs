package value

const (
	// TrueString is a string representation of the boolean value true.
	TrueString = "true"

	// FalseString is a string representation of the boolean value false.
	FalseString = "false"
)

var (
	// TrueValue is the singleton instance representing the boolean value true.
	TrueValue Object = &Bool{value: true}

	// FalseValue is the singleton instance representing the boolean value false.
	FalseValue Object = &Bool{value: false}

	// UndefinedValue is the singleton instance representing the undefined value.
	UndefinedValue Object = &Undefined{}
)

// CallableFunc is a function signature for the callable functions.
type CallableFunc = func(args ...Object) (ret Object, err error)
