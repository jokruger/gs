package core

import (
	"errors"
	"fmt"
)

var (
	ErrLogicError            = errors.New("logic error")
	ErrStackOverflow         = errors.New("stack overflow")
	ErrObjectAllocLimit      = errors.New("object allocation limit exceeded")
	ErrBytesLimit            = errors.New("bytes size limit exceeded")
	ErrStringLimit           = errors.New("string size limit exceeded")
	ErrDecodeBinarySize      = errors.New("invalid binary size")
	ErrBinaryNotSupported    = errors.New("binary serialization not supported")
	ErrInvalidArgumentType   = errors.New("invalid argument type")
	ErrIndexOutOfBounds      = errors.New("index out of bounds")
	ErrWrongNumArguments     = errors.New("wrong number of arguments")
	ErrInvalidAccessMode     = errors.New("invalid access mode")
	ErrNotAccessible         = errors.New("object is not accessible")
	ErrNotAssignable         = errors.New("object is not assignable")
	ErrNotCallable           = errors.New("object is not callable")
	ErrInvalidIndexType      = errors.New("invalid index type")
	ErrInvalidSelector       = errors.New("invalid selector")
	ErrNotImplemented        = errors.New("not implemented")
	ErrInvalidBinaryOperator = errors.New("invalid binary operator")
	ErrInvalidValueKind      = errors.New("invalid value kind")
	ErrInvalidMethodError    = errors.New("invalid method error")
)

func NewLogicError(context string) error {
	return fmt.Errorf("%w: %s", ErrLogicError, context)
}

func NewStackOverflowError(context string) error {
	return fmt.Errorf("%w: %s", ErrStackOverflow, context)
}

func NewObjectAllocLimitError(context string) error {
	return fmt.Errorf("%w: %s", ErrObjectAllocLimit, context)
}

func NewBytesLimitError(context string) error {
	return fmt.Errorf("%w: %s", ErrBytesLimit, context)
}

func NewStringLimitError(context string) error {
	return fmt.Errorf("%w: %s", ErrStringLimit, context)
}

func NewDecodeBinarySizeError(valType string, expected int, got int) error {
	return fmt.Errorf("%w: type %s expects %d bytes, got %d", ErrDecodeBinarySize, valType, expected, got)
}

func NewBinaryNotSupportedError(valType string) error {
	return fmt.Errorf("%w: type %s", ErrBinaryNotSupported, valType)
}

func NewInvalidArgumentTypeError(context string, name string, expected string, got string) error {
	return fmt.Errorf("%w: (%s) argument %s expects type %s, got %s", ErrInvalidArgumentType, context, name, expected, got)
}

func NewIndexOutOfBoundsError(context string, idx int, size int) error {
	return fmt.Errorf("%w: (%s) %d out of range [0, %d]", ErrIndexOutOfBounds, context, idx, size)
}

func NewWrongNumArgumentsError(context string, expected string, got int) error {
	return fmt.Errorf("%w: (%s) expected %s argument(s), got %d", ErrWrongNumArguments, context, expected, got)
}

func NewInvalidAccessModeError(dt string, mode string) error {
	return fmt.Errorf("%w: type %s does not support %s access", ErrInvalidAccessMode, dt, mode)
}

func NewNotAccessibleError(valType string) error {
	return fmt.Errorf("%w: type %s does not support indexing or field access", ErrNotAccessible, valType)
}

func NewNotAssignableError(valType string) error {
	return fmt.Errorf("%w: type %s does not support assignment via indexing or field access", ErrNotAssignable, valType)
}

func NewNotCallableError(valType string) error {
	return fmt.Errorf("%w: type %s does not support function call", ErrNotCallable, valType)
}

func NewInvalidIndexTypeError(context string, expected string, got string) error {
	return fmt.Errorf("%w: (%s) expected %s, got %s", ErrInvalidIndexType, context, expected, got)
}

func NewInvalidSelectorError(valType string, sel string) error {
	return fmt.Errorf("%w: type %s has no property or method %s", ErrInvalidSelector, valType, sel)
}

func NewNotImplementedError(feature string) error {
	return fmt.Errorf("%w: %s", ErrNotImplemented, feature)
}

func NewInvalidBinaryOperatorError(op string, left string, right string) error {
	return fmt.Errorf("%w: %s %s %s", ErrInvalidBinaryOperator, left, op, right)
}

func NewInvalidValueKindError(kind ValueKind) error {
	return fmt.Errorf("%w: %d", ErrInvalidValueKind, kind)
}

func NewInvalidMethodError(method string, valType string) error {
	return fmt.Errorf("%w: type %s has no method %s", ErrInvalidMethodError, valType, method)
}
