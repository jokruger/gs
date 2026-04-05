package value

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/araddon/dateparse"
	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/internal/conv"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/token"
)

type String struct {
	Object
	value []rune
}

func NewStaticString(v string) core.Value {
	o := &String{}
	o.Set(v)
	return core.NewObject(o)
}

func (o *String) GobDecode(b []byte) error {
	o.Set(string(b))
	return nil
}

func (o *String) GobEncode() ([]byte, error) {
	return []byte(string(o.value)), nil
}

func (o *String) Set(s string) {
	o.value = []rune(s)
}

func (o *String) Value() string {
	return string(o.value)
}

func (o *String) Runes() []rune {
	return o.value
}

func (o *String) IsEmpty() bool {
	return len(o.value) == 0
}

func (o *String) Len() int {
	return len(o.value)
}

func (o *String) At(i int) rune {
	return o.value[i]
}

func (o *String) Get(i int) (rune, bool) {
	if i < 0 || i >= len(o.value) {
		return 0, false
	}
	return o.value[i], true
}

func (o *String) Substring(start, end int) string {
	return string(o.value[start:end])
}

func (o *String) Append(s string) {
	o.value = append(o.value, []rune(s)...)
}

func (o *String) TypeName() string {
	return "string"
}

func (o *String) String() string {
	return strconv.Quote(string(o.value))
}

func (o *String) Interface() any {
	return string(o.value)
}

func (o *String) BinaryOp(vm core.VM, op token.Token, rhs core.Value) (core.Value, error) {
	alloc := vm.Allocator()
	v, ok := rhs.AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
	}

	switch op {
	case token.Add:
		return alloc.NewStringValue(string(o.value) + v), nil
	case token.Less:
		return core.NewBool(string(o.value) < v), nil
	case token.LessEq:
		return core.NewBool(string(o.value) <= v), nil
	case token.Greater:
		return core.NewBool(string(o.value) > v), nil
	case token.GreaterEq:
		return core.NewBool(string(o.value) >= v), nil
	}

	return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
}

func (o *String) Equals(x core.Value) bool {
	t, ok := x.AsString()
	if !ok {
		return false
	}
	return string(o.value) == t
}

func (o *String) Copy(alloc core.Allocator) core.Value {
	return alloc.NewStringValue(string(o.value))
}

func (o *String) Method(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	switch name {
	case "trim":
		return o.fnTrim(vm, "string.trim", args...)
	default:
		return core.NewUndefined(), core.NewInvalidMethodError(name, o.TypeName())
	}
}

func (o *String) Access(vm core.VM, index core.Value, mode core.Opcode) (core.Value, error) {
	alloc := vm.Allocator()

	if mode == parser.OpIndex {
		i, ok := index.AsInt()
		if !ok {
			return core.NewUndefined(), core.NewInvalidIndexTypeError("string access", "int", index.TypeName())
		}
		if i < 0 || i >= int64(len(o.value)) {
			return core.NewUndefined(), nil
		}
		return core.NewChar(o.value[i]), nil
	}

	k, ok := index.AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidSelectorError(o.TypeName(), k)
	}

	switch k {
	case "string":
		return core.NewObject(o), nil

	case "array":
		arr := make([]core.Value, len(o.value))
		for i, r := range o.value {
			arr[i] = core.NewChar(r)
		}
		return alloc.NewArrayValue(arr, false), nil

	case "bool":
		b, _ := o.AsBool()
		return core.NewBool(b), nil

	case "bytes":
		return alloc.NewBytesValue([]byte(string(o.value))), nil

	case "char":
		if len(o.value) == 1 {
			return core.NewChar(o.value[0]), nil
		}
		return core.NewChar(0), nil

	case "float":
		f, _ := o.AsFloat()
		return core.NewFloat(f), nil

	case "int":
		i, _ := o.AsInt()
		return core.NewInt(i), nil

	case "time":
		t, _ := o.AsTime()
		return alloc.NewTimeValue(t), nil

	case "record":
		m := make(map[string]core.Value, len(o.value))
		for i, r := range o.value {
			m[strconv.Itoa(i)] = core.NewChar(r)
		}
		return alloc.NewRecordValue(m, false), nil

	case "empty":
		return core.NewBool(len(o.value) == 0), nil

	case "len":
		return core.NewInt(int64(len(o.value))), nil

	case "first":
		if len(o.value) == 0 {
			return core.NewUndefined(), nil
		}
		return core.NewChar(o.value[0]), nil

	case "last":
		if len(o.value) == 0 {
			return core.NewUndefined(), nil
		}
		return core.NewChar(o.value[len(o.value)-1]), nil

	case "lower":
		t := make([]rune, len(o.value))
		for i, r := range o.value {
			t[i] = unicode.ToLower(r)
		}
		return alloc.NewStringValue(string(t)), nil

	case "upper":
		t := make([]rune, len(o.value))
		for i, r := range o.value {
			t[i] = unicode.ToUpper(r)
		}
		return alloc.NewStringValue(string(t)), nil

	default:
		return core.NewUndefined(), core.NewInvalidSelectorError(o.TypeName(), k)
	}
}

func (o *String) Assign(core.Value, core.Value) error {
	return core.NewNotAssignableError(o.TypeName())
}

func (o *String) Iterate(alloc core.Allocator) core.Iterator {
	return alloc.NewStringIterator(o.value)
}

func (o *String) IsString() bool {
	return true
}

func (o *String) IsTrue() bool {
	return len(o.value) > 0
}

func (o *String) IsFalse() bool {
	return len(o.value) == 0
}

func (o *String) IsIterable() bool {
	return true
}

func (o *String) AsString() (string, bool) {
	return string(o.value), true
}

func (o *String) AsInt() (int64, bool) {
	i, err := strconv.ParseInt(string(o.value), 10, 64)
	if err == nil {
		return i, true
	}
	return 0, false
}

func (o *String) AsFloat() (float64, bool) {
	f, err := strconv.ParseFloat(string(o.value), 64)
	if err == nil {
		return f, true
	}
	return 0, false
}

func (o *String) AsBool() (bool, bool) {
	return conv.ParseBool(string(o.value))
}

func (o *String) AsChar() (rune, bool) {
	if len(o.value) == 1 {
		return o.value[0], true
	}
	return 0, false
}

func (o *String) AsBytes() ([]byte, bool) {
	return []byte(string(o.value)), true
}

func (o *String) AsTime() (time.Time, bool) {
	val, err := dateparse.ParseAny(string(o.value))
	if err != nil {
		return time.Time{}, false
	}
	return val, true
}

func (o *String) fnTrim(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	if len(args) > 1 {
		return core.NewUndefined(), core.NewWrongNumArgumentsError(name, "0 or 1", len(args))
	}

	if len(args) == 0 {
		return vm.Allocator().NewStringValue(strings.Trim(string(o.value), " \t\n")), nil
	}

	s, ok := args[0].AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidArgumentTypeError(name, "first", "string", args[0].TypeName())
	}

	return vm.Allocator().NewStringValue(strings.Trim(string(o.value), s)), nil
}
