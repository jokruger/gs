package error

import (
	"errors"
)

var (
	ErrInvalidOperator  = errors.New("invalid operator")
	ErrNotImplemented   = errors.New("not implemented")
	ErrInvalidRangeStep = errors.New("range step must be greater than 0")
)
