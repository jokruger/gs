package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/stdlib/json"
	"github.com/jokruger/gs/value"
)

var jsonModule = map[string]core.Object{
	"decode": &value.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &value.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &value.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &value.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *value.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &value.Error{
				Value: &value.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *value.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &value.Error{
				Value: &value.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &value.Error{Value: &value.String{Value: err.Error()}}, nil
	}

	return &value.Bytes{Value: b}, nil
}

func jsonIndent(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 3 {
		return nil, gse.ErrWrongNumArguments
	}

	prefix, ok := args[1].ToString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := args[2].ToString()
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *value.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &value.Error{
				Value: &value.String{Value: err.Error()},
			}, nil
		}
		return &value.Bytes{Value: dst.Bytes()}, nil
	case *value.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &value.Error{
				Value: &value.String{Value: err.Error()},
			}, nil
		}
		return &value.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *value.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &value.Bytes{Value: dst.Bytes()}, nil
	case *value.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &value.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
