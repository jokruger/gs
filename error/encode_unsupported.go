package error

import "fmt"

type EncodeUnsupportedError struct {
	Type    string
	Context string
}

func (e *EncodeUnsupportedError) Error() string {
	if e.Context != "" {
		return fmt.Sprintf("encode %s to binary: unsupported (%s)", e.Type, e.Context)
	}
	return fmt.Sprintf("encode %s to binary: unsupported", e.Type)
}
