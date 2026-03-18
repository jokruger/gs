package error

import "fmt"

type DecodeUnsupportedError struct {
	Type    string
	Context string
}

func (e *DecodeUnsupportedError) Error() string {
	if e.Context != "" {
		return fmt.Sprintf("decode %s from binary: unsupported (%s)", e.Type, e.Context)
	}
	return fmt.Sprintf("decode %s from binary: unsupported", e.Type)
}
