package value

import (
	"fmt"
	"time"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/token"
)

type Time struct {
	Object
	value time.Time
}

func (o *Time) GobDecode(b []byte) error {
	var t time.Time
	if err := t.GobDecode(b); err != nil {
		return err
	}
	o.Set(t)
	return nil
}

func (o *Time) GobEncode() ([]byte, error) {
	return o.value.GobEncode()
}

func (o *Time) Set(t time.Time) {
	o.value = t
}

func (o *Time) Value() time.Time {
	return o.value
}

func (o *Time) TypeName() string {
	return "time"
}

func (o *Time) String() string {
	return fmt.Sprintf("time(%q)", o.value.String())
}

func (o *Time) Interface() any {
	return o.value
}

func (o *Time) BinaryOp(vm core.VM, op token.Token, rhs core.Value) (core.Value, error) {
	alloc := vm.Allocator()

	if rhs.IsInt() {
		r := rhs.Int()
		switch op {
		case token.Add: // time + int => time
			return alloc.NewTimeValue(o.value.Add(time.Duration(r))), nil
		case token.Sub: // time - int => time
			return alloc.NewTimeValue(o.value.Add(time.Duration(-r))), nil
		}
	}

	v, ok := rhs.AsTime()
	if !ok {
		return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
	}

	switch op {
	case token.Sub: // time - time => int (duration)
		return core.NewInt(int64(o.value.Sub(v))), nil
	case token.Less: // time < time => bool
		return core.NewBool(o.value.Before(v)), nil
	case token.Greater:
		return core.NewBool(o.value.After(v)), nil
	case token.LessEq:
		return core.NewBool(o.value.Equal(v) || o.value.Before(v)), nil
	case token.GreaterEq:
		return core.NewBool(o.value.Equal(v) || o.value.After(v)), nil
	}

	return core.NewUndefined(), core.NewInvalidBinaryOperatorError(op.String(), o.TypeName(), rhs.TypeName())
}

func (o *Time) Equals(x core.Value) bool {
	t, ok := x.AsTime()
	if !ok {
		return false
	}
	return o.value.Equal(t)
}

func (o *Time) Copy(alloc core.Allocator) core.Value {
	return alloc.NewTimeValue(o.value)
}

func (o *Time) Method(vm core.VM, name string, args ...core.Value) (core.Value, error) {
	return core.NewUndefined(), core.NewInvalidMethodError(name, o.TypeName())
}

func (o *Time) Access(vm core.VM, index core.Value, op core.Opcode) (core.Value, error) {
	k, ok := index.AsString()
	if !ok {
		return core.NewUndefined(), core.NewInvalidIndexTypeError("map access", "string", index.TypeName())
	}

	alloc := vm.Allocator()
	switch k {
	case "time":
		return core.NewObject(o, false), nil

	case "bool":
		return core.NewBool(o.IsTrue()), nil

	case "int":
		return core.NewInt(o.value.Unix()), nil

	case "string":
		return alloc.NewStringValue(o.value.String()), nil

	case "year":
		return core.NewInt(int64(o.value.Year())), nil

	case "month":
		return core.NewInt(int64(o.value.Month())), nil

	case "day":
		return core.NewInt(int64(o.value.Day())), nil

	case "hour":
		return core.NewInt(int64(o.value.Hour())), nil

	case "minute":
		return core.NewInt(int64(o.value.Minute())), nil

	case "second":
		return core.NewInt(int64(o.value.Second())), nil

	case "nanosecond":
		return core.NewInt(int64(o.value.Nanosecond())), nil

	case "unix":
		return core.NewInt(o.value.Unix()), nil

	case "unix_nano":
		return core.NewInt(o.value.UnixNano()), nil

	case "week_day":
		return core.NewInt(int64(o.value.Weekday())), nil

	case "year_day":
		return core.NewInt(int64(o.value.YearDay())), nil

	case "month_name":
		return alloc.NewStringValue(o.value.Month().String()), nil

	case "week_day_name":
		return alloc.NewStringValue(o.value.Weekday().String()), nil

	case "utc":
		return alloc.NewTimeValue(o.value.UTC()), nil

	case "local":
		return alloc.NewTimeValue(o.value.Local()), nil

	case "date_str":
		return alloc.NewStringValue(o.value.Format(time.DateOnly)), nil

	case "time_str":
		return alloc.NewStringValue(o.value.Format(time.TimeOnly)), nil

	case "date_time_str":
		return alloc.NewStringValue(o.value.Format(time.DateTime)), nil

	case "zone_offset":
		_, offset := o.value.Zone()
		return core.NewInt(int64(offset)), nil

	case "zone_name":
		name, _ := o.value.Zone()
		return alloc.NewStringValue(name), nil

	default:
		return core.NewUndefined(), core.NewInvalidSelectorError(o.TypeName(), k)
	}
}

func (o *Time) Assign(core.Value, core.Value) error {
	return core.NewNotAssignableError(o.TypeName())
}

func (o *Time) IsTime() bool {
	return true
}

func (o *Time) IsTrue() bool {
	return !o.IsFalse()
}

func (o *Time) IsFalse() bool {
	return o.value.IsZero()
}

func (o *Time) AsString() (string, bool) {
	return o.value.String(), true
}

func (o *Time) AsInt() (int64, bool) {
	return o.value.Unix(), true
}

func (o *Time) AsBool() (bool, bool) {
	return !o.IsFalse(), true
}

func (o *Time) AsTime() (time.Time, bool) {
	return o.value, true
}
