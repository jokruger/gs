package core

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/jokruger/gs/token"
)

// The minimum required fields for Value are ptr, d64 and kind. This allow to store primitive types such as int, float, rune; and heap allocated objects.
// Due to padding, the size of such structure will be 24 bytes on 64-bit architectures. So we can add some d32, d16 and d8 extra fields for free.
type Value struct {
	Type uint8
	Data uint64
	Ptr  unsafe.Pointer
}

func (v *Value) Set(val Value) {
	*v = val
}

func (v Value) EncodeJSON() ([]byte, error) {
	b, err := ValueTypes[v.Type].TypeEncodeJSON(v)
	if err != nil {
		return nil, fmt.Errorf("json encoding failed for type %s: %w", v.TypeName(), err)
	}
	return b, nil
}

func (v Value) EncodeBinary() ([]byte, error) {
	b, err := ValueTypes[v.Type].TypeEncodeBinary(v)
	if err != nil {
		return nil, fmt.Errorf("binary encoding failed for type %s: %w", v.TypeName(), err)
	}
	return append([]byte{v.Type}, b...), nil
}

func (v Value) GobEncode() ([]byte, error) {
	return v.EncodeBinary()
}

func (v *Value) DecodeBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("binary decoding failed (type header): expected at least 1 byte for type, got %d", len(data))
	}

	var t Value
	t.Type = data[0]
	if err := ValueTypes[t.Type].TypeDecodeBinary(&t, data[1:]); err != nil {
		return fmt.Errorf("binary decoding failed for type %d: %w", t.Type, err)
	}
	*v = t

	return nil
}

func (v *Value) GobDecode(data []byte) error {
	return v.DecodeBinary(data)
}

func (v Value) Next() bool {
	return ValueTypes[v.Type].TypeNext(v)
}

func (v Value) Key(alloc Allocator) Value {
	return ValueTypes[v.Type].TypeKey(v, alloc)
}

func (v Value) Value(alloc Allocator) Value {
	return ValueTypes[v.Type].TypeValue(v, alloc)
}

func (v Value) TypeName() string {
	return ValueTypes[v.Type].TypeName(v)
}

func (v Value) String() string {
	return ValueTypes[v.Type].TypeString(v)
}

func (v Value) Interface() any {
	return ValueTypes[v.Type].TypeInterface(v)
}

func (v Value) Arity() int8 {
	return ValueTypes[v.Type].TypeArity(v)
}

func (v Value) IsUserDefined() bool {
	return v.Type >= VT_USER_DEFINED
}

func (v Value) IsTrue() bool {
	return ValueTypes[v.Type].TypeIsTrue(v)
}

func (v Value) IsIterable() bool {
	return ValueTypes[v.Type].TypeIsIterable(v)
}

func (v Value) IsCallable() bool {
	return ValueTypes[v.Type].TypeIsCallable(v)
}

func (v Value) IsVariadic() bool {
	return ValueTypes[v.Type].TypeIsVariadic(v)
}

func (v Value) IsImmutable() bool {
	return ValueTypes[v.Type].TypeIsImmutable(v)
}

func (v Value) Contains(e Value) bool {
	return ValueTypes[v.Type].TypeContains(v, e)
}

func (v Value) AsBool() (bool, bool) {
	return ValueTypes[v.Type].TypeAsBool(v)
}

func (v Value) AsChar() (rune, bool) {
	return ValueTypes[v.Type].TypeAsChar(v)
}

func (v Value) AsInt() (int64, bool) {
	return ValueTypes[v.Type].TypeAsInt(v)
}

func (v Value) AsFloat() (float64, bool) {
	return ValueTypes[v.Type].TypeAsFloat(v)
}

func (v Value) AsTime() (time.Time, bool) {
	return ValueTypes[v.Type].TypeAsTime(v)
}

func (v Value) AsString() (string, bool) {
	return ValueTypes[v.Type].TypeAsString(v)
}

func (v Value) AsBytes() ([]byte, bool) {
	return ValueTypes[v.Type].TypeAsBytes(v)
}

func (v Value) BinaryOp(a Allocator, op token.Token, rhs Value) (Value, error) {
	return ValueTypes[v.Type].TypeBinaryOp(v, a, op, rhs)
}

func (v Value) Equal(rhs Value) bool {
	return ValueTypes[v.Type].TypeEqual(v, rhs)
}

func (v *Value) Copy(alloc Allocator) Value {
	return ValueTypes[v.Type].TypeCopy(*v, alloc)
}

func (v Value) MethodCall(vm VM, name string, args []Value) (Value, error) {
	return ValueTypes[v.Type].TypeMethodCall(v, vm, name, args)
}

func (v Value) Access(vm VM, index Value, mode Opcode) (Value, error) {
	return ValueTypes[v.Type].TypeAccess(v, vm.Allocator(), index, mode)
}

func (v Value) Assign(idx Value, val Value) error {
	return ValueTypes[v.Type].TypeAssign(v, idx, val)
}

func (v Value) Iterator(alloc Allocator) Value {
	return ValueTypes[v.Type].TypeIterator(v, alloc)
}

func (v Value) Call(vm VM, args []Value) (Value, error) {
	return ValueTypes[v.Type].TypeCall(v, vm, args)
}

func (v Value) Len() int64 {
	return ValueTypes[v.Type].TypeLen(v)
}

func (v Value) Append(a Allocator, args []Value) (Value, error) {
	return ValueTypes[v.Type].TypeAppend(v, a, args)
}

func (v Value) Delete(key Value) (Value, error) {
	return ValueTypes[v.Type].TypeDelete(v, key)
}
