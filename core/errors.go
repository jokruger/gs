package core

import (
	"errors"
	"fmt"
)

var (
	ErrStackOverflow    = errors.New("stack overflow")
	ErrObjectAllocLimit = errors.New("object allocation limit exceeded")
	ErrBytesLimit       = errors.New("exceeding bytes size limit")
)

func StackOverflow(context string) error {
	return fmt.Errorf("%w: %s", ErrStackOverflow, context)
}

func ObjectAllocLimit(context string) error {
	return fmt.Errorf("%w: %s", ErrObjectAllocLimit, context)
}

func BytesLimit(context string) error {
	return fmt.Errorf("%w: %s", ErrBytesLimit, context)
}
