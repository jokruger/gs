package core

import (
	"errors"
	"fmt"
)

var (
	ErrLogicError          = errors.New("logic error")
	ErrStackOverflow       = errors.New("stack overflow")
	ErrObjectAllocLimit    = errors.New("object allocation limit exceeded")
	ErrBytesLimit          = errors.New("bytes size limit exceeded")
	ErrStringLimit         = errors.New("string size limit exceeded")
	ErrDecodeBinarySize    = errors.New("invalid binary size")
	ErrBinaryNotSupported  = errors.New("binary serialization not supported")
	ErrInvalidArgumentType = errors.New("invalid argument type")
	ErrIndexOutOfBounds    = errors.New("index out of bounds")
	ErrWrongNumArguments   = errors.New("wrong number of arguments")
	ErrInvalidAccessMode   = errors.New("invalid access mode")
	ErrNotAccessible       = errors.New("object is not accessible")
	ErrNotAssignable       = errors.New("object is not assignable")
	ErrInvalidIndexType    = errors.New("invalid index type")
	ErrInvalidSelector     = errors.New("invalid selector")
)

func LogicError(context string) error {
	return fmt.Errorf("%w: %s", ErrLogicError, context)
}

func StackOverflow(context string) error {
	return fmt.Errorf("%w: %s", ErrStackOverflow, context)
}

func ObjectAllocLimit(context string) error {
	return fmt.Errorf("%w: %s", ErrObjectAllocLimit, context)
}

func BytesLimit(context string) error {
	return fmt.Errorf("%w: %s", ErrBytesLimit, context)
}

func StringLimit(context string) error {
	return fmt.Errorf("%w: %s", ErrStringLimit, context)
}

func DecodeBinarySize(obj Object, expected int, got int) error {
	return fmt.Errorf("%w: type %s expects %d bytes, got %d", ErrDecodeBinarySize, obj.TypeName(), expected, got)
}

func BinaryNotSupported(obj Object) error {
	return fmt.Errorf("%w: type %s", ErrBinaryNotSupported, obj.TypeName())
}

func InvalidArgumentType(context string, name string, expected string, got Object) error {
	return fmt.Errorf("%w: %s argument '%s' expects type %s, got %s", ErrInvalidArgumentType, context, name, expected, got.TypeName())
}

func IndexOutOfBounds(context string, idx int, size int) error {
	return fmt.Errorf("%w: %s: index %d out of range [0,%d)", ErrIndexOutOfBounds, context, idx, size)
}

func WrongNumArguments(context string, expected string, got int) error {
	return fmt.Errorf("%w: %s: expected %s argument(s), got %d", ErrWrongNumArguments, context, expected, got)
}

func InvalidAccessMode(dt string, mode string) error {
	return fmt.Errorf("%w: type %s does not support access mode '%s'", ErrInvalidAccessMode, dt, mode)
}

func NotAccessible(obj Object) error {
	return fmt.Errorf("%w: type %s does not support indexing or field access", ErrNotAccessible, obj.TypeName())
}

func NotAssignable(obj Object) error {
	return fmt.Errorf("%w: type %s does not support assignment via indexing or field access", ErrNotAssignable, obj.TypeName())
}

func InvalidIndexType(context string, expected string, got Object) error {
	return fmt.Errorf("%w: %s: expected %s, got %s", ErrInvalidIndexType, context, expected, got.TypeName())
}

func InvalidSelector(obj Object, sel string) error {
	return fmt.Errorf("%w: type %s has no selector '%s'", ErrInvalidSelector, obj.TypeName(), sel)
}
