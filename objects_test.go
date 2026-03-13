package gs_test

import (
	"testing"

	"github.com/jokruger/gs"
	"github.com/jokruger/gs/require"
	"github.com/jokruger/gs/token"
)

func TestObject_TypeName(t *testing.T) {
	var o gs.Object = &gs.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &gs.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &gs.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &gs.String{}
	require.Equal(t, "string", o.TypeName())
	o = &gs.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &gs.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &gs.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &gs.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &gs.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &gs.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &gs.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &gs.UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &gs.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &gs.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &gs.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &gs.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o gs.Object = &gs.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &gs.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &gs.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &gs.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &gs.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &gs.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &gs.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &gs.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &gs.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &gs.Array{Value: []gs.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &gs.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &gs.Map{Value: map[string]gs.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &gs.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &gs.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &gs.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &gs.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &gs.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &gs.Undefined{}
	require.True(t, o.IsFalsy())
	o = &gs.Error{}
	require.True(t, o.IsFalsy())
	o = &gs.Bytes{}
	require.True(t, o.IsFalsy())
	o = &gs.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o gs.Object = &gs.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &gs.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &gs.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &gs.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &gs.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &gs.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &gs.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &gs.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &gs.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &gs.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &gs.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &gs.Error{Value: &gs.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &gs.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &gs.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &gs.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &gs.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &gs.Bytes{}
	require.Equal(t, "", o.String())
	o = &gs.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o gs.Object = &gs.Char{}
	_, err := o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.Bool{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.Map{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.StringIterator{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.MapIterator{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.Undefined{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
	o = &gs.Error{}
	_, err = o.BinaryOp(token.Add, gs.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &gs.Array{Value: nil}, token.Add,
		&gs.Array{Value: nil}, &gs.Array{Value: nil})
	testBinaryOp(t, &gs.Array{Value: nil}, token.Add,
		&gs.Array{Value: []gs.Object{}}, &gs.Array{Value: nil})
	testBinaryOp(t, &gs.Array{Value: []gs.Object{}}, token.Add,
		&gs.Array{Value: nil}, &gs.Array{Value: []gs.Object{}})
	testBinaryOp(t, &gs.Array{Value: []gs.Object{}}, token.Add,
		&gs.Array{Value: []gs.Object{}},
		&gs.Array{Value: []gs.Object{}})
	testBinaryOp(t, &gs.Array{Value: nil}, token.Add,
		&gs.Array{Value: []gs.Object{
			&gs.Int{Value: 1},
		}}, &gs.Array{Value: []gs.Object{
			&gs.Int{Value: 1},
		}})
	testBinaryOp(t, &gs.Array{Value: nil}, token.Add,
		&gs.Array{Value: []gs.Object{
			&gs.Int{Value: 1},
			&gs.Int{Value: 2},
			&gs.Int{Value: 3},
		}}, &gs.Array{Value: []gs.Object{
			&gs.Int{Value: 1},
			&gs.Int{Value: 2},
			&gs.Int{Value: 3},
		}})
	testBinaryOp(t, &gs.Array{Value: []gs.Object{
		&gs.Int{Value: 1},
		&gs.Int{Value: 2},
		&gs.Int{Value: 3},
	}}, token.Add, &gs.Array{Value: nil},
		&gs.Array{Value: []gs.Object{
			&gs.Int{Value: 1},
			&gs.Int{Value: 2},
			&gs.Int{Value: 3},
		}})
	testBinaryOp(t, &gs.Array{Value: []gs.Object{
		&gs.Int{Value: 1},
		&gs.Int{Value: 2},
		&gs.Int{Value: 3},
	}}, token.Add, &gs.Array{Value: []gs.Object{
		&gs.Int{Value: 4},
		&gs.Int{Value: 5},
		&gs.Int{Value: 6},
	}}, &gs.Array{Value: []gs.Object{
		&gs.Int{Value: 1},
		&gs.Int{Value: 2},
		&gs.Int{Value: 3},
		&gs.Int{Value: 4},
		&gs.Int{Value: 5},
		&gs.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &gs.Error{Value: &gs.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &gs.Error{Value: &gs.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.Add,
				&gs.Float{Value: r}, &gs.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.Sub,
				&gs.Float{Value: r}, &gs.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.Mul,
				&gs.Float{Value: r}, &gs.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &gs.Float{Value: l}, token.Quo,
					&gs.Float{Value: r}, &gs.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.Less,
				&gs.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.Greater,
				&gs.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.LessEq,
				&gs.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gs.Float{Value: l}, token.GreaterEq,
				&gs.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.Add,
				&gs.Int{Value: r}, &gs.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.Sub,
				&gs.Int{Value: r}, &gs.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.Mul,
				&gs.Int{Value: r}, &gs.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &gs.Float{Value: l}, token.Quo,
					&gs.Int{Value: r},
					&gs.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.Less,
				&gs.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.Greater,
				&gs.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.LessEq,
				&gs.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Float{Value: l}, token.GreaterEq,
				&gs.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.Add,
				&gs.Int{Value: r}, &gs.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.Sub,
				&gs.Int{Value: r}, &gs.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.Mul,
				&gs.Int{Value: r}, &gs.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &gs.Int{Value: l}, token.Quo,
					&gs.Int{Value: r}, &gs.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &gs.Int{Value: l}, token.Rem,
					&gs.Int{Value: r}, &gs.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.And, &gs.Int{Value: 0},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.And, &gs.Int{Value: 0},
		&gs.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.And, &gs.Int{Value: 1},
		&gs.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.And, &gs.Int{Value: 1},
		&gs.Int{Value: int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.And, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.And, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: int64(0xffffffff)}, token.And,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1984}, token.And,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &gs.Int{Value: -1984}, token.And,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Or, &gs.Int{Value: 0},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Or, &gs.Int{Value: 0},
		&gs.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Or, &gs.Int{Value: 1},
		&gs.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Or, &gs.Int{Value: 1},
		&gs.Int{Value: int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Or, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Or, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: int64(0xffffffff)}, token.Or,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1984}, token.Or,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: -1984}, token.Or,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Xor, &gs.Int{Value: 0},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Xor, &gs.Int{Value: 0},
		&gs.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Xor, &gs.Int{Value: 1},
		&gs.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Xor, &gs.Int{Value: 1},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.Xor, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.Xor, &gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: int64(0xffffffff)}, token.Xor,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1984}, token.Xor,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: -1984}, token.Xor,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.AndNot, &gs.Int{Value: 0},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.AndNot, &gs.Int{Value: 0},
		&gs.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.AndNot,
		&gs.Int{Value: 1}, &gs.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.AndNot, &gs.Int{Value: 1},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 0}, token.AndNot,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: 1}, token.AndNot,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: int64(0xffffffff)}, token.AndNot,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(0)})
	testBinaryOp(t,
		&gs.Int{Value: 1984}, token.AndNot,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gs.Int{Value: -1984}, token.AndNot,
		&gs.Int{Value: int64(0xffffffff)},
		&gs.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&gs.Int{Value: 0}, token.Shl, &gs.Int{Value: s},
			&gs.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: 1}, token.Shl, &gs.Int{Value: s},
			&gs.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: 2}, token.Shl, &gs.Int{Value: s},
			&gs.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: -1}, token.Shl, &gs.Int{Value: s},
			&gs.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: -2}, token.Shl, &gs.Int{Value: s},
			&gs.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: int64(0xffffffff)}, token.Shl,
			&gs.Int{Value: s},
			&gs.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&gs.Int{Value: 0}, token.Shr, &gs.Int{Value: s},
			&gs.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: 1}, token.Shr, &gs.Int{Value: s},
			&gs.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: 2}, token.Shr, &gs.Int{Value: s},
			&gs.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: -1}, token.Shr, &gs.Int{Value: s},
			&gs.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: -2}, token.Shr, &gs.Int{Value: s},
			&gs.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&gs.Int{Value: int64(0xffffffff)}, token.Shr,
			&gs.Int{Value: s},
			&gs.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.Less,
				&gs.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.Greater,
				&gs.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.LessEq,
				&gs.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gs.Int{Value: l}, token.GreaterEq,
				&gs.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.Add,
				&gs.Float{Value: r},
				&gs.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.Sub,
				&gs.Float{Value: r},
				&gs.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.Mul,
				&gs.Float{Value: r},
				&gs.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &gs.Int{Value: l}, token.Quo,
					&gs.Float{Value: r},
					&gs.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.Less,
				&gs.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.Greater,
				&gs.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.LessEq,
				&gs.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gs.Int{Value: l}, token.GreaterEq,
				&gs.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &gs.Map{Value: make(map[string]gs.Object)}
	k := &gs.Int{Value: 1}
	v := &gs.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	require.NoError(t, err)

	res, err := m.IndexGet(k)
	require.NoError(t, err)
	require.Equal(t, v, res)
}

func TestString_BinaryOp(t *testing.T) {
	lstr := "abcde"
	rstr := "01234"
	for l := 0; l < len(lstr); l++ {
		for r := 0; r < len(rstr); r++ {
			ls := lstr[l:]
			rs := rstr[r:]
			testBinaryOp(t, &gs.String{Value: ls}, token.Add,
				&gs.String{Value: rs},
				&gs.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &gs.String{Value: ls}, token.Add,
				&gs.Char{Value: rc},
				&gs.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs gs.Object,
	op token.Token,
	rhs gs.Object,
	expected gs.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) gs.Object {
	if b {
		return gs.TrueValue
	}
	return gs.FalseValue
}
