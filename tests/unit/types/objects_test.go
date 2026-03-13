package types_test

import (
	"testing"

	"github.com/jokruger/gs/tests/require"
	"github.com/jokruger/gs/token"
	gst "github.com/jokruger/gs/types"
)

func TestObject_TypeName(t *testing.T) {
	var o gst.Object = &gst.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &gst.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &gst.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &gst.String{}
	require.Equal(t, "string", o.TypeName())
	o = &gst.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &gst.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &gst.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &gst.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &gst.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &gst.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &gst.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &gst.UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &gst.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &gst.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &gst.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &gst.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o gst.Object = &gst.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &gst.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &gst.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &gst.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &gst.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &gst.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &gst.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &gst.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &gst.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &gst.Array{Value: []gst.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &gst.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &gst.Map{Value: map[string]gst.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &gst.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &gst.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &gst.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &gst.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &gst.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &gst.Undefined{}
	require.True(t, o.IsFalsy())
	o = &gst.Error{}
	require.True(t, o.IsFalsy())
	o = &gst.Bytes{}
	require.True(t, o.IsFalsy())
	o = &gst.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o gst.Object = &gst.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &gst.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &gst.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &gst.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &gst.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &gst.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &gst.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &gst.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &gst.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &gst.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &gst.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &gst.Error{Value: &gst.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &gst.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &gst.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &gst.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &gst.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &gst.Bytes{}
	require.Equal(t, "", o.String())
	o = &gst.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o gst.Object = &gst.Char{}
	_, err := o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.Bool{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.Map{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.StringIterator{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.MapIterator{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.Undefined{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
	o = &gst.Error{}
	_, err = o.BinaryOp(token.Add, gst.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &gst.Array{Value: nil}, token.Add,
		&gst.Array{Value: nil}, &gst.Array{Value: nil})
	testBinaryOp(t, &gst.Array{Value: nil}, token.Add,
		&gst.Array{Value: []gst.Object{}}, &gst.Array{Value: nil})
	testBinaryOp(t, &gst.Array{Value: []gst.Object{}}, token.Add,
		&gst.Array{Value: nil}, &gst.Array{Value: []gst.Object{}})
	testBinaryOp(t, &gst.Array{Value: []gst.Object{}}, token.Add,
		&gst.Array{Value: []gst.Object{}},
		&gst.Array{Value: []gst.Object{}})
	testBinaryOp(t, &gst.Array{Value: nil}, token.Add,
		&gst.Array{Value: []gst.Object{
			&gst.Int{Value: 1},
		}}, &gst.Array{Value: []gst.Object{
			&gst.Int{Value: 1},
		}})
	testBinaryOp(t, &gst.Array{Value: nil}, token.Add,
		&gst.Array{Value: []gst.Object{
			&gst.Int{Value: 1},
			&gst.Int{Value: 2},
			&gst.Int{Value: 3},
		}}, &gst.Array{Value: []gst.Object{
			&gst.Int{Value: 1},
			&gst.Int{Value: 2},
			&gst.Int{Value: 3},
		}})
	testBinaryOp(t, &gst.Array{Value: []gst.Object{
		&gst.Int{Value: 1},
		&gst.Int{Value: 2},
		&gst.Int{Value: 3},
	}}, token.Add, &gst.Array{Value: nil},
		&gst.Array{Value: []gst.Object{
			&gst.Int{Value: 1},
			&gst.Int{Value: 2},
			&gst.Int{Value: 3},
		}})
	testBinaryOp(t, &gst.Array{Value: []gst.Object{
		&gst.Int{Value: 1},
		&gst.Int{Value: 2},
		&gst.Int{Value: 3},
	}}, token.Add, &gst.Array{Value: []gst.Object{
		&gst.Int{Value: 4},
		&gst.Int{Value: 5},
		&gst.Int{Value: 6},
	}}, &gst.Array{Value: []gst.Object{
		&gst.Int{Value: 1},
		&gst.Int{Value: 2},
		&gst.Int{Value: 3},
		&gst.Int{Value: 4},
		&gst.Int{Value: 5},
		&gst.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &gst.Error{Value: &gst.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &gst.Error{Value: &gst.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.Add,
				&gst.Float{Value: r}, &gst.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.Sub,
				&gst.Float{Value: r}, &gst.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.Mul,
				&gst.Float{Value: r}, &gst.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &gst.Float{Value: l}, token.Quo,
					&gst.Float{Value: r}, &gst.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.Less,
				&gst.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.Greater,
				&gst.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.LessEq,
				&gst.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &gst.Float{Value: l}, token.GreaterEq,
				&gst.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.Add,
				&gst.Int{Value: r}, &gst.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.Sub,
				&gst.Int{Value: r}, &gst.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.Mul,
				&gst.Int{Value: r}, &gst.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &gst.Float{Value: l}, token.Quo,
					&gst.Int{Value: r},
					&gst.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.Less,
				&gst.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.Greater,
				&gst.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.LessEq,
				&gst.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Float{Value: l}, token.GreaterEq,
				&gst.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.Add,
				&gst.Int{Value: r}, &gst.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.Sub,
				&gst.Int{Value: r}, &gst.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.Mul,
				&gst.Int{Value: r}, &gst.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &gst.Int{Value: l}, token.Quo,
					&gst.Int{Value: r}, &gst.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &gst.Int{Value: l}, token.Rem,
					&gst.Int{Value: r}, &gst.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.And, &gst.Int{Value: 0},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.And, &gst.Int{Value: 0},
		&gst.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.And, &gst.Int{Value: 1},
		&gst.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.And, &gst.Int{Value: 1},
		&gst.Int{Value: int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.And, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.And, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: int64(0xffffffff)}, token.And,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1984}, token.And,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &gst.Int{Value: -1984}, token.And,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Or, &gst.Int{Value: 0},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Or, &gst.Int{Value: 0},
		&gst.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Or, &gst.Int{Value: 1},
		&gst.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Or, &gst.Int{Value: 1},
		&gst.Int{Value: int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Or, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Or, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: int64(0xffffffff)}, token.Or,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1984}, token.Or,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: -1984}, token.Or,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Xor, &gst.Int{Value: 0},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Xor, &gst.Int{Value: 0},
		&gst.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Xor, &gst.Int{Value: 1},
		&gst.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Xor, &gst.Int{Value: 1},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.Xor, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.Xor, &gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: int64(0xffffffff)}, token.Xor,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1984}, token.Xor,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: -1984}, token.Xor,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.AndNot, &gst.Int{Value: 0},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.AndNot, &gst.Int{Value: 0},
		&gst.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.AndNot,
		&gst.Int{Value: 1}, &gst.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.AndNot, &gst.Int{Value: 1},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 0}, token.AndNot,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: 1}, token.AndNot,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: int64(0xffffffff)}, token.AndNot,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(0)})
	testBinaryOp(t,
		&gst.Int{Value: 1984}, token.AndNot,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&gst.Int{Value: -1984}, token.AndNot,
		&gst.Int{Value: int64(0xffffffff)},
		&gst.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&gst.Int{Value: 0}, token.Shl, &gst.Int{Value: s},
			&gst.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: 1}, token.Shl, &gst.Int{Value: s},
			&gst.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: 2}, token.Shl, &gst.Int{Value: s},
			&gst.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: -1}, token.Shl, &gst.Int{Value: s},
			&gst.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: -2}, token.Shl, &gst.Int{Value: s},
			&gst.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: int64(0xffffffff)}, token.Shl,
			&gst.Int{Value: s},
			&gst.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&gst.Int{Value: 0}, token.Shr, &gst.Int{Value: s},
			&gst.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: 1}, token.Shr, &gst.Int{Value: s},
			&gst.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: 2}, token.Shr, &gst.Int{Value: s},
			&gst.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: -1}, token.Shr, &gst.Int{Value: s},
			&gst.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: -2}, token.Shr, &gst.Int{Value: s},
			&gst.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&gst.Int{Value: int64(0xffffffff)}, token.Shr,
			&gst.Int{Value: s},
			&gst.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.Less,
				&gst.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.Greater,
				&gst.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.LessEq,
				&gst.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &gst.Int{Value: l}, token.GreaterEq,
				&gst.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.Add,
				&gst.Float{Value: r},
				&gst.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.Sub,
				&gst.Float{Value: r},
				&gst.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.Mul,
				&gst.Float{Value: r},
				&gst.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &gst.Int{Value: l}, token.Quo,
					&gst.Float{Value: r},
					&gst.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.Less,
				&gst.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.Greater,
				&gst.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.LessEq,
				&gst.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &gst.Int{Value: l}, token.GreaterEq,
				&gst.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &gst.Map{Value: make(map[string]gst.Object)}
	k := &gst.Int{Value: 1}
	v := &gst.String{Value: "abcdef"}
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
			testBinaryOp(t, &gst.String{Value: ls}, token.Add,
				&gst.String{Value: rs},
				&gst.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &gst.String{Value: ls}, token.Add,
				&gst.Char{Value: rc},
				&gst.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs gst.Object,
	op token.Token,
	rhs gst.Object,
	expected gst.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) gst.Object {
	if b {
		return gst.TrueValue
	}
	return gst.FalseValue
}
