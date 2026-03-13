package gs

import (
	"errors"
	"fmt"
	"time"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

const (
	// GlobalsSize is the maximum number of global variables for a VM.
	GlobalsSize = 1024

	// StackSize is the maximum stack size for a VM.
	StackSize = 2048

	// MaxFrames is the maximum number of function frames for a VM.
	MaxFrames = 1024

	// SourceFileExtDefault is the default extension for source files.
	SourceFileExtDefault = ".gs"
)

// CountObjects returns the number of objects that a given object o contains.
// For scalar value types, it will always be 1. For compound value types,
// this will include its elements and all of their elements recursively.
func CountObjects(o gst.Object) (c int) {
	c = 1
	switch o := o.(type) {
	case *gst.Array:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *gst.ImmutableArray:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *gst.Map:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *gst.ImmutableMap:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *gst.Error:
		c += CountObjects(o.Value)
	}
	return
}

// ToInterface attempts to convert an object o to an interface{} value
func ToInterface(o gst.Object) (res interface{}) {
	switch o := o.(type) {
	case *gst.Int:
		res = o.Value
	case *gst.String:
		res = o.Value
	case *gst.Float:
		res = o.Value
	case *gst.Bool:
		res = o == gst.TrueValue
	case *gst.Char:
		res = o.Value
	case *gst.Bytes:
		res = o.Value
	case *gst.Array:
		res = make([]interface{}, len(o.Value))
		for i, val := range o.Value {
			res.([]interface{})[i] = ToInterface(val)
		}
	case *gst.ImmutableArray:
		res = make([]interface{}, len(o.Value))
		for i, val := range o.Value {
			res.([]interface{})[i] = ToInterface(val)
		}
	case *gst.Map:
		res = make(map[string]interface{})
		for key, v := range o.Value {
			res.(map[string]interface{})[key] = ToInterface(v)
		}
	case *gst.ImmutableMap:
		res = make(map[string]interface{})
		for key, v := range o.Value {
			res.(map[string]interface{})[key] = ToInterface(v)
		}
	case *gst.Time:
		res = o.Value
	case *gst.Error:
		res = errors.New(o.String())
	case *gst.Undefined:
		res = nil
	case gst.Object:
		return o
	}
	return
}

// FromInterface will attempt to convert an interface{} v to a Gs Object
func FromInterface(v interface{}) (gst.Object, error) {
	switch v := v.(type) {
	case nil:
		return gst.UndefinedValue, nil
	case string:
		if len(v) > gst.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		return &gst.String{Value: v}, nil
	case int64:
		return &gst.Int{Value: v}, nil
	case int:
		return &gst.Int{Value: int64(v)}, nil
	case bool:
		if v {
			return gst.TrueValue, nil
		}
		return gst.FalseValue, nil
	case rune:
		return &gst.Char{Value: v}, nil
	case byte:
		return &gst.Char{Value: rune(v)}, nil
	case float64:
		return &gst.Float{Value: v}, nil
	case []byte:
		if len(v) > gst.MaxBytesLen {
			return nil, gse.ErrBytesLimit
		}
		return &gst.Bytes{Value: v}, nil
	case error:
		return &gst.Error{Value: &gst.String{Value: v.Error()}}, nil
	case map[string]gst.Object:
		return &gst.Map{Value: v}, nil
	case map[string]interface{}:
		kv := make(map[string]gst.Object)
		for vk, vv := range v {
			vo, err := FromInterface(vv)
			if err != nil {
				return nil, err
			}
			kv[vk] = vo
		}
		return &gst.Map{Value: kv}, nil
	case []gst.Object:
		return &gst.Array{Value: v}, nil
	case []interface{}:
		arr := make([]gst.Object, len(v))
		for i, e := range v {
			vo, err := FromInterface(e)
			if err != nil {
				return nil, err
			}
			arr[i] = vo
		}
		return &gst.Array{Value: arr}, nil
	case time.Time:
		return &gst.Time{Value: v}, nil
	case gst.Object:
		return v, nil
	case gst.CallableFunc:
		return &gst.UserFunction{Value: v}, nil
	}
	return nil, fmt.Errorf("cannot convert to object: %T", v)
}
