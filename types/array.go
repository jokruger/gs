package types

import (
	"fmt"
	"strings"

	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/token"
)

type Array struct {
	ObjectImpl
	Value []Object
}

func (o *Array) TypeName() string {
	return "array"
}

func (o *Array) String() string {
	var elements []string
	for _, e := range o.Value {
		elements = append(elements, e.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

func (o *Array) BinaryOp(op token.Token, rhs Object) (Object, error) {
	if rhs, ok := rhs.(*Array); ok {
		switch op {
		case token.Add:
			if len(rhs.Value) == 0 {
				return o, nil
			}
			return &Array{Value: append(o.Value, rhs.Value...)}, nil
		}
	}
	return nil, gse.ErrInvalidOperator
}

func (o *Array) Copy() Object {
	var c []Object
	for _, elem := range o.Value {
		c = append(c, elem.Copy())
	}
	return &Array{Value: c}
}

func (o *Array) IsFalsy() bool {
	return len(o.Value) == 0
}

func (o *Array) Equals(x Object) bool {
	var xVal []Object
	switch x := x.(type) {
	case *Array:
		xVal = x.Value
	case *ImmutableArray:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for i, e := range o.Value {
		if !e.Equals(xVal[i]) {
			return false
		}
	}
	return true
}

func (o *Array) IndexGet(index Object) (res Object, err error) {
	intIdx, ok := index.(*Int)
	if !ok {
		err = gse.ErrInvalidIndexType
		return
	}
	idxVal := int(intIdx.Value)
	if idxVal < 0 || idxVal >= len(o.Value) {
		res = UndefinedValue
		return
	}
	res = o.Value[idxVal]
	return
}

func (o *Array) IndexSet(index, value Object) (err error) {
	intIdx, ok := index.ToInt()
	if !ok {
		err = gse.ErrInvalidIndexType
		return
	}
	if intIdx < 0 || intIdx >= len(o.Value) {
		err = gse.ErrIndexOutOfBounds
		return
	}
	o.Value[intIdx] = value
	return nil
}

func (o *Array) Iterate() Iterator {
	return &ArrayIterator{
		v: o.Value,
		l: len(o.Value),
	}
}

func (o *Array) CanIterate() bool {
	return true
}

func (o *Array) ToString() (string, bool) {
	return o.String(), true
}

func (o *Array) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
