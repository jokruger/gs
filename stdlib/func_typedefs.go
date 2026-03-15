package stdlib

import (
	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

// FuncASRIE transform a function of 'func(string) (int, error)' signature
// into CallableFunc type.
func FuncASRIE(fn func(string) (int, error)) core.NativeFunc {
	return func(args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, gse.ErrWrongNumArguments
		}
		s1, ok := args[0].AsString()
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		return &value.Int{Value: int64(res)}, nil
	}
}

// FuncASRYE transform a function of 'func(string) ([]byte, error)' signature
// into CallableFunc type.
func FuncASRYE(fn func(string) ([]byte, error)) core.NativeFunc {
	return func(args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, gse.ErrWrongNumArguments
		}
		s1, ok := args[0].AsString()
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > core.MaxBytesLen {
			return nil, gse.ErrBytesLimit
		}
		return &value.Bytes{Value: res}, nil
	}
}

// FuncAIRSsE transform a function of 'func(int) ([]string, error)' signature
// into CallableFunc type.
func FuncAIRSsE(fn func(int) ([]string, error)) core.NativeFunc {
	return func(args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, gse.ErrWrongNumArguments
		}
		i1, ok := args[0].AsInt()
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(int(i1))
		if err != nil {
			return wrapError(err), nil
		}
		arr := &value.Array{}
		for _, r := range res {
			if len(r) > core.MaxStringLen {
				return nil, gse.ErrStringLimit
			}
			arr.Value = append(arr.Value, &value.String{Value: r})
		}
		return arr, nil
	}
}

// FuncAIRS transform a function of 'func(int) string' signature into
// CallableFunc type.
func FuncAIRS(fn func(int) string) core.NativeFunc {
	return func(args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, gse.ErrWrongNumArguments
		}
		i1, ok := args[0].AsInt()
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s := fn(int(i1))
		if len(s) > core.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		return &value.String{Value: s}, nil
	}
}
