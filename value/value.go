package value

import (
	"time"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/token"
)

const (
	// TrueString is a string representation of the boolean value true.
	TrueString = "true"

	// FalseString is a string representation of the boolean value false.
	FalseString = "false"
)

var (
	// TrueValue is the singleton instance representing the boolean value true.
	TrueValue *Bool = &Bool{value: true}

	// FalseValue is the singleton instance representing the boolean value false.
	FalseValue *Bool = &Bool{value: false}

	// UndefinedValue is the singleton instance representing the undefined value.
	UndefinedValue *Undefined = &Undefined{}
)

func init() {
	TrueValue.Set(true)
	FalseValue.Set(false)
	UndefinedValue.Set()
}

/* === Object (Base) === */

type Object struct {
}

func (o *Object) TypeName() string {
	return "<object>"
}

func (o *Object) String() string {
	return o.TypeName()
}

func (o *Object) Interface() any {
	return o
}

func (o *Object) Arity() int {
	return 0
}

func (o *Object) BinaryOp(op token.Token, rhs core.Object) (core.Object, error) {
	return nil, core.InvalidBinaryOperator(op.String(), o, rhs)
}

func (o *Object) Equals(x core.Object) bool {
	return o == x
}

func (o *Object) Copy() core.Object {
	return o
}

func (o *Object) Access(core.Object, core.Opcode) (core.Object, error) {
	return nil, core.NotAccessible(o)
}

func (o *Object) Assign(core.Object, core.Object) error {
	return core.NotAssignable(o)
}

func (o *Object) Iterate() core.Iterator {
	return nil
}

func (o *Object) Call(core.VM, ...core.Object) (core.Object, error) {
	return nil, nil
}

func (o *Object) IsFalsy() bool {
	return o == nil
}

func (o *Object) IsIterable() bool {
	return false
}

func (o *Object) IsCallable() bool {
	return false
}

func (o *Object) IsImmutable() bool {
	return false
}

func (o *Object) IsVariadic() bool {
	return false
}

func (o *Object) AsString() (string, bool) {
	return "", false
}

func (o *Object) AsInt() (int64, bool) {
	return 0, false
}

func (o *Object) AsFloat() (float64, bool) {
	return 0, false
}

func (o *Object) AsBool() (bool, bool) {
	return false, false
}

func (o *Object) AsRune() (rune, bool) {
	return 0, false
}

func (o *Object) AsByteSlice() ([]byte, bool) {
	return nil, false
}

func (o *Object) AsTime() (time.Time, bool) {
	return time.Time{}, false
}
