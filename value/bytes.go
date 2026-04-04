package value

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/token"
)

type Bytes struct {
	Object
	value []byte
}

func (o *Bytes) GobDecode(b []byte) error {
	decoded := make([]byte, len(b))
	copy(decoded, b)
	o.Set(decoded)
	return nil
}

func (o *Bytes) GobEncode() ([]byte, error) {
	encoded := make([]byte, len(o.value))
	copy(encoded, o.value)
	return encoded, nil
}

func (o *Bytes) Set(v []byte) {
	o.value = v
	if o.value == nil {
		o.value = make([]byte, 0)
	}
}

func (o *Bytes) Value() []byte {
	return o.value
}

func (o *Bytes) IsEmpty() bool {
	return len(o.value) == 0
}

func (o *Bytes) Len() int {
	return len(o.value)
}

func (o *Bytes) Append(v []byte) {
	o.value = append(o.value, v...)
}

func (o *Bytes) At(i int) byte {
	return o.value[i]
}

func (o *Bytes) Clear() {
	o.value = o.value[:0]
}

func (o *Bytes) Slice(start, end int) []byte {
	return o.value[start:end]
}

func (o *Bytes) TypeName() string {
	return "bytes"
}

func (o *Bytes) String() string {
	es := make([]string, len(o.value))
	for i, b := range o.value {
		es[i] = fmt.Sprintf("%d", b)
	}
	return fmt.Sprintf("bytes([%s])", strings.Join(es, ", "))
}

func (o *Bytes) Interface() any {
	return o.value
}

func (o *Bytes) BinaryOp(vm core.VM, op token.Token, rhs core.Value) (core.Value, error) {
	alloc := vm.Allocator()
	v, ok := rhs.AsBytes()
	if !ok {
		return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
	}

	switch op {
	case token.Add:
		return alloc.NewBytesValue(append(o.value, v...)), nil
	}

	return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
}

func (o *Bytes) Equals(x core.Value) bool {
	t, ok := x.AsBytes()
	if !ok {
		return false
	}
	return bytes.Equal(o.value, t)
}

func (o *Bytes) Copy(alloc core.Allocator) core.Value {
	t := make([]byte, len(o.value))
	copy(t, o.value)
	return core.NewObject(alloc.NewBytes(t), false)
}

func (o *Bytes) Method(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	return core.NewUndefined(), core.NewInvalidMethodError(name, o.TypeName())
}

func (o *Bytes) Access(vm core.VM, index core.Value, mode core.Opcode) (core.Value, error) {
	alloc := vm.Allocator()

	if mode == parser.OpIndex {
		i, ok := index.AsInt()
		if !ok {
			return core.NewUndefined(), core.NewInvalidIndexTypeError("bytes index", "int", index.TypeName())
		}
		if i < 0 || i >= int64(len(o.value)) {
			return core.NewUndefined(), nil
		}
		return core.NewInt(int64(o.value[i])), nil
	}

	k, ok := index.AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidSelectorError(o.TypeName(), k)
	}

	switch k {
	case "bytes":
		return core.NewObject(o, false), nil

	case "array":
		arr := make([]core.Value, len(o.value))
		for i, b := range o.value {
			arr[i] = core.NewInt(int64(b))
		}
		return alloc.NewArrayValue(arr, false), nil

	case "record":
		m := make(map[string]core.Value, len(o.value))
		for i, b := range o.value {
			m[strconv.Itoa(i)] = core.NewInt(int64(b))
		}
		return alloc.NewMapValue(m, false), nil

	case "string":
		return alloc.NewStringValue(string(o.value)), nil

	case "empty":
		return core.NewBool(o.IsEmpty()), nil

	case "len":
		return core.NewInt(int64(o.Len())), nil

	case "first":
		if len(o.value) == 0 {
			return core.NewUndefined(), nil
		}
		return core.NewInt(int64(o.value[0])), nil

	case "last":
		if len(o.value) == 0 {
			return core.NewUndefined(), nil
		}
		return core.NewInt(int64(o.value[len(o.value)-1])), nil

	default:
		return core.NewUndefined(), core.NewInvalidSelectorError(o.TypeName(), k)
	}
}

func (o *Bytes) Assign(core.Value, core.Value) error {
	return core.NewNotAssignableError(o.TypeName())
}

func (o *Bytes) Iterate(alloc core.Allocator) core.Iterator {
	return alloc.NewBytesIterator(o.value)
}

func (o *Bytes) IsBytes() bool {
	return true
}

func (o *Bytes) IsTrue() bool {
	return len(o.value) > 0
}

func (o *Bytes) IsFalse() bool {
	return len(o.value) == 0
}

func (o *Bytes) IsIterable() bool {
	return true
}

func (o *Bytes) AsString() (string, bool) {
	return string(o.value), true
}

func (o *Bytes) AsBool() (bool, bool) {
	return o.IsTrue(), true
}

func (o *Bytes) AsBytes() ([]byte, bool) {
	return o.value, true
}
