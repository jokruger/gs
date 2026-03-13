package value

import (
	"fmt"
	"strings"

	gse "github.com/jokruger/gs/error"
)

type Map struct {
	ObjectImpl
	Value map[string]Object
}

func (o *Map) TypeName() string {
	return "map"
}

func (o *Map) String() string {
	var pairs []string
	for k, v := range o.Value {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

func (o *Map) Copy() Object {
	c := make(map[string]Object)
	for k, v := range o.Value {
		c[k] = v.Copy()
	}
	return &Map{Value: c}
}

func (o *Map) IsFalsy() bool {
	return len(o.Value) == 0
}

func (o *Map) Equals(x Object) bool {
	var xVal map[string]Object
	switch x := x.(type) {
	case *Map:
		xVal = x.Value
	case *ImmutableMap:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for k, v := range o.Value {
		tv := xVal[k]
		if !v.Equals(tv) {
			return false
		}
	}
	return true
}

func (o *Map) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.ToString()
	if !ok {
		err = gse.ErrInvalidIndexType
		return
	}
	res, ok = o.Value[strIdx]
	if !ok {
		res = UndefinedValue
	}
	return
}

func (o *Map) IndexSet(index, value Object) (err error) {
	strIdx, ok := index.ToString()
	if !ok {
		err = gse.ErrInvalidIndexType
		return
	}
	o.Value[strIdx] = value
	return nil
}

func (o *Map) Iterate() Iterator {
	var keys []string
	for k := range o.Value {
		keys = append(keys, k)
	}
	return &MapIterator{
		v: o.Value,
		k: keys,
		l: len(keys),
	}
}

func (o *Map) CanIterate() bool {
	return true
}

func (o *Map) ToString() (string, bool) {
	return o.String(), true
}

func (o *Map) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
