package stdlib

import (
	"bytes"
	gojson "encoding/json"

	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/stdlib/json"
	gst "github.com/jokruger/gs/types"
)

var jsonModule = map[string]gst.Object{
	"decode": &gst.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &gst.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &gst.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &gst.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gst.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &gst.Error{
				Value: &gst.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *gst.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &gst.Error{
				Value: &gst.String{Value: err.Error()},
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

func jsonEncode(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &gst.Error{Value: &gst.String{Value: err.Error()}}, nil
	}

	return &gst.Bytes{Value: b}, nil
}

func jsonIndent(args ...gst.Object) (ret gst.Object, err error) {
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
	case *gst.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &gst.Error{
				Value: &gst.String{Value: err.Error()},
			}, nil
		}
		return &gst.Bytes{Value: dst.Bytes()}, nil
	case *gst.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &gst.Error{
				Value: &gst.String{Value: err.Error()},
			}, nil
		}
		return &gst.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gst.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &gst.Bytes{Value: dst.Bytes()}, nil
	case *gst.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &gst.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
