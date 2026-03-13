package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/stdlib/json"
)

var jsonModule = map[string]gs.Object{
	"decode": &gs.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &gs.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &gs.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &gs.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gs.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &gs.Error{
				Value: &gs.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *gs.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &gs.Error{
				Value: &gs.String{Value: err.Error()},
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

func jsonEncode(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &gs.Error{Value: &gs.String{Value: err.Error()}}, nil
	}

	return &gs.Bytes{Value: b}, nil
}

func jsonIndent(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 3 {
		return nil, gse.ErrWrongNumArguments
	}

	prefix, ok := gs.ToString(args[1])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := gs.ToString(args[2])
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *gs.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &gs.Error{
				Value: &gs.String{Value: err.Error()},
			}, nil
		}
		return &gs.Bytes{Value: dst.Bytes()}, nil
	case *gs.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &gs.Error{
				Value: &gs.String{Value: err.Error()},
			}, nil
		}
		return &gs.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gs.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &gs.Bytes{Value: dst.Bytes()}, nil
	case *gs.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &gs.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
