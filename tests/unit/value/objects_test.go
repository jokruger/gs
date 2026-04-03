package value

import (
	"testing"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/parser"
	mock "github.com/jokruger/gs/tests"
	"github.com/jokruger/gs/tests/require"
	"github.com/jokruger/gs/token"
	"github.com/jokruger/gs/value"
	_ "github.com/jokruger/gs/vm"
)

var vm = mock.Vm
var alloc = mock.Alloc

func TestObject_Value(t *testing.T) {
	var v core.Value
	var x core.Value
	var bs []byte
	var err error

	v.SetKind(core.V_BOOL)
	require.Equal(t, core.V_BOOL, v.Kind())
	v.SetBool(true)
	require.Equal(t, true, v.Bool())
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_BOOL, x.Kind())
	require.Equal(t, true, x.Bool())
	v.SetBool(false)
	require.Equal(t, false, v.Bool())
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_BOOL, x.Kind())
	require.Equal(t, false, x.Bool())

	v.SetKind(core.V_CHAR)
	require.Equal(t, core.V_CHAR, v.Kind())
	v.SetChar('A')
	require.Equal(t, 'A', v.Char())
	v.SetChar('B')
	require.Equal(t, 'B', v.Char())
	v.SetChar('₴')
	require.Equal(t, '₴', v.Char())
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_CHAR, x.Kind())
	require.Equal(t, '₴', x.Char())

	v.SetKind(core.V_INT)
	require.Equal(t, core.V_INT, v.Kind())
	v.SetInt(123)
	require.Equal(t, int64(123), v.Int())
	v.SetInt(-456)
	require.Equal(t, int64(-456), v.Int())
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_INT, x.Kind())
	require.Equal(t, int64(-456), x.Int())

	v.SetKind(core.V_FLOAT)
	require.Equal(t, core.V_FLOAT, v.Kind())
	v.SetFloat(3.14)
	require.Equal(t, 3.14, v.Float())
	v.SetFloat(-2.71828)
	require.Equal(t, -2.71828, v.Float())
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_FLOAT, x.Kind())
	require.Equal(t, -2.71828, x.Float())

	v.SetInt(0)
	v.SetKind(core.V_OBJECT)
	require.Equal(t, core.V_OBJECT, v.Kind())
	o := alloc.NewString("hello")
	v.SetObject(o)
	s, _ := v.AsString()
	require.Equal(t, "hello", s)
	bs, err = v.GobEncode()
	require.NoError(t, err)
	err = x.GobDecode(bs)
	require.NoError(t, err)
	require.Equal(t, core.V_OBJECT, x.Kind())
	s, _ = x.AsString()
	require.Equal(t, "hello", s)
}

func TestObject_TypeName(t *testing.T) {
	var o core.Value
	var i core.Iterator
	var obj core.Object

	o = core.NewInt(0)
	require.Equal(t, "int", o.TypeName())

	o = core.NewFloat(0)
	require.Equal(t, "float", o.TypeName())

	o = core.NewChar(0)
	require.Equal(t, "char", o.TypeName())

	o = alloc.NewStringValue("")
	require.Equal(t, "string", o.TypeName())

	o = core.NewBool(false)
	require.Equal(t, "bool", o.TypeName())

	o = alloc.NewArrayValue(nil, false)
	require.Equal(t, "array", o.TypeName())

	o = alloc.NewRecordValue(nil, false)
	require.Equal(t, "record", o.TypeName())

	i = alloc.NewArrayIterator(nil)
	require.Equal(t, "array-iterator", i.TypeName())

	i = alloc.NewStringIterator(nil)
	require.Equal(t, "string-iterator", i.TypeName())

	i = alloc.NewMapIterator(nil)
	require.Equal(t, "map-iterator", i.TypeName())

	o = alloc.NewBuiltinFunctionValue("fn", nil, 0, false)
	require.Equal(t, "<builtin-function:fn/0>", o.TypeName())

	obj = &value.CompiledFunction{}
	require.Equal(t, "<compiled-function/0>", obj.TypeName())

	o = core.NewUndefined()
	require.Equal(t, "undefined", o.TypeName())

	o = alloc.NewErrorValue(core.NewUndefined())
	require.Equal(t, "error", o.TypeName())

	o = alloc.NewBytesValue(nil)
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o core.Value
	var obj core.Object

	o = core.NewInt(0)
	require.True(t, o.IsFalse())

	o = core.NewInt(1)
	require.False(t, o.IsFalse())

	o = core.NewFloat(0)
	require.False(t, o.IsFalse())

	o = core.NewFloat(1)
	require.False(t, o.IsFalse())

	o = core.NewChar(' ')
	require.False(t, o.IsFalse())

	o = core.NewChar('T')
	require.False(t, o.IsFalse())

	o = alloc.NewStringValue("")
	require.True(t, o.IsFalse())

	o = alloc.NewStringValue(" ")
	require.False(t, o.IsFalse())

	o = alloc.NewArrayValue(nil, false)
	require.True(t, o.IsFalse())

	o = alloc.NewArrayValue([]core.Value{core.NewUndefined()}, false)
	require.False(t, o.IsFalse())

	o = alloc.NewRecordValue(nil, false)
	require.True(t, o.IsFalse())

	o = alloc.NewRecordValue(map[string]core.Value{"a": core.NewUndefined()}, false)
	require.False(t, o.IsFalse())

	obj = alloc.NewStringIterator(nil)
	require.True(t, obj.IsFalse())

	obj = alloc.NewArrayIterator(nil)
	require.True(t, obj.IsFalse())

	obj = alloc.NewMapIterator(nil)
	require.True(t, obj.IsFalse())

	obj = alloc.NewBuiltinFunction("fn", nil, 0, false)
	require.False(t, obj.IsFalse())

	obj = &value.CompiledFunction{}
	require.False(t, obj.IsFalse())

	o = core.NewUndefined()
	require.True(t, o.IsFalse())

	o = alloc.NewErrorValue(core.NewUndefined())
	require.True(t, o.IsFalse())

	o = alloc.NewBytesValue(nil)
	require.True(t, o.IsFalse())

	o = alloc.NewBytesValue([]byte{1, 2})
	require.False(t, o.IsFalse())
}

func TestObject_String(t *testing.T) {
	var o core.Value
	var obj core.Object

	o = core.NewInt(0)
	require.Equal(t, "0", o.String())

	o = core.NewInt(1)
	require.Equal(t, "1", o.String())

	o = core.NewFloat(0)
	require.Equal(t, "0", o.String())

	o = core.NewFloat(1)
	require.Equal(t, "1", o.String())

	o = core.NewChar(' ')
	require.Equal(t, "' '", o.String())

	o = core.NewChar('T')
	require.Equal(t, "'T'", o.String())

	o = alloc.NewStringValue("")
	require.Equal(t, `""`, o.String())

	o = alloc.NewStringValue(" ")
	require.Equal(t, `" "`, o.String())

	o = alloc.NewArrayValue(nil, false)
	require.Equal(t, "[]", o.String())

	o = alloc.NewRecordValue(nil, false)
	require.Equal(t, "{}", o.String())

	o = alloc.NewErrorValue(core.NewUndefined())
	require.Equal(t, "error(undefined)", o.String())

	o = alloc.NewErrorValue(alloc.NewStringValue("error 1"))
	require.Equal(t, `error("error 1")`, o.String())

	obj = alloc.NewStringIterator(nil)
	require.Equal(t, "<string-iterator>", obj.String())

	obj = alloc.NewArrayIterator(nil)
	require.Equal(t, "<array-iterator>", obj.String())

	obj = alloc.NewMapIterator(nil)
	require.Equal(t, "<map-iterator>", obj.String())

	o = core.NewUndefined()
	require.Equal(t, "undefined", o.String())

	o = alloc.NewBytesValue(nil)
	require.Equal(t, "bytes([])", o.String())

	o = alloc.NewBytesValue([]byte("foo"))
	require.Equal(t, "bytes([102, 111, 111])", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o core.Value
	var obj core.Object

	o = core.NewChar(0)
	_, err := o.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	o = core.NewBool(false)
	_, err = o.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	o = alloc.NewRecordValue(nil, false)
	_, err = o.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	obj = alloc.NewArrayIterator(nil)
	_, err = obj.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	obj = alloc.NewStringIterator(nil)
	_, err = obj.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	obj = alloc.NewMapIterator(nil)
	_, err = obj.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	obj = alloc.NewBuiltinFunction("fn", nil, 0, false)
	_, err = obj.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	obj = &value.CompiledFunction{}
	_, err = obj.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	o = core.NewUndefined()
	_, err = o.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)

	o = alloc.NewErrorValue(core.NewUndefined())
	_, err = o.BinaryOp(vm, token.Add, core.NewUndefined())
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, alloc.NewArrayValue(nil, false), token.Add,
		alloc.NewArrayValue(nil, false), alloc.NewArrayValue(nil, false))
	testBinaryOp(t, alloc.NewArrayValue(nil, false), token.Add,
		alloc.NewArrayValue([]core.Value{}, false), alloc.NewArrayValue(nil, false))
	testBinaryOp(t, alloc.NewArrayValue([]core.Value{}, false), token.Add,
		alloc.NewArrayValue(nil, false), alloc.NewArrayValue([]core.Value{}, false))
	testBinaryOp(t, alloc.NewArrayValue([]core.Value{}, false), token.Add,
		alloc.NewArrayValue([]core.Value{}, false),
		alloc.NewArrayValue([]core.Value{}, false))
	testBinaryOp(t, alloc.NewArrayValue(nil, false), token.Add,
		alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
		}, false), alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
		}, false))
	testBinaryOp(t, alloc.NewArrayValue(nil, false), token.Add,
		alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
			core.NewInt(2),
			core.NewInt(3),
		}, false), alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
			core.NewInt(2),
			core.NewInt(3),
		}, false))
	testBinaryOp(t, alloc.NewArrayValue([]core.Value{
		core.NewInt(1),
		core.NewInt(2),
		core.NewInt(3),
	}, false), token.Add, alloc.NewArrayValue(nil, false),
		alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
			core.NewInt(2),
			core.NewInt(3),
		}, false))
	testBinaryOp(t, alloc.NewArrayValue([]core.Value{
		core.NewInt(1),
		core.NewInt(2),
		core.NewInt(3),
	}, false), token.Add, alloc.NewArrayValue([]core.Value{
		core.NewInt(4),
		core.NewInt(5),
		core.NewInt(6),
	}, false), alloc.NewArrayValue([]core.Value{
		core.NewInt(1),
		core.NewInt(2),
		core.NewInt(3),
		core.NewInt(4),
		core.NewInt(5),
		core.NewInt(6),
	}, false))
}

func TestError_Equals(t *testing.T) {
	err1 := alloc.NewErrorValue(alloc.NewStringValue("some error"))
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = alloc.NewErrorValue(alloc.NewStringValue("some error"))
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = alloc.NewErrorValue(alloc.NewStringValue("some error 2"))
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))

	bool1 := core.NewBool(true)
	bool2 := core.NewBool(true)
	bool3 := core.NewBool(false)

	char1 := core.NewChar('A')
	char2 := core.NewChar('A')
	char3 := core.NewChar('B')

	int1 := core.NewInt(123)
	int2 := core.NewInt(123)
	int3 := core.NewInt(456)

	float1 := core.NewFloat(3.14)
	float2 := core.NewFloat(3.14)
	float3 := core.NewFloat(2.71828)

	string1 := alloc.NewStringValue("hello")
	string2 := alloc.NewStringValue("hello")
	string3 := alloc.NewStringValue("world")

	bytes1 := alloc.NewBytesValue([]byte("foo"))
	bytes2 := alloc.NewBytesValue([]byte("foo"))
	bytes3 := alloc.NewBytesValue([]byte("bar"))

	array1 := alloc.NewArrayValue([]core.Value{core.NewInt(1), core.NewInt(2)}, false)
	array2 := alloc.NewArrayValue([]core.Value{core.NewInt(1), core.NewInt(2)}, false)
	array3 := alloc.NewArrayValue([]core.Value{core.NewInt(1), core.NewInt(3)}, false)

	map1 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(1)}, false)
	map2 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(1)}, false)
	map3 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(2)}, false)

	record1 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(1)}, false)
	record2 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(1)}, false)
	record3 := alloc.NewRecordValue(map[string]core.Value{"a": core.NewInt(2)}, false)

	// compare to undefined
	require.False(t, bool1.Equals(core.NewUndefined()))
	require.False(t, char1.Equals(core.NewUndefined()))
	require.False(t, int1.Equals(core.NewUndefined()))
	require.False(t, float1.Equals(core.NewUndefined()))
	require.False(t, string1.Equals(core.NewUndefined()))
	require.False(t, bytes1.Equals(core.NewUndefined()))
	require.False(t, array1.Equals(core.NewUndefined()))
	require.False(t, map1.Equals(core.NewUndefined()))
	require.False(t, record1.Equals(core.NewUndefined()))

	// compare to equal
	require.True(t, bool1.Equals(bool2))
	require.True(t, char1.Equals(char2))
	require.True(t, int1.Equals(int2))
	require.True(t, float1.Equals(float2))
	require.True(t, string1.Equals(string2))
	require.True(t, bytes1.Equals(bytes2))
	require.True(t, array1.Equals(array2))
	require.True(t, map1.Equals(map2))
	require.True(t, record1.Equals(record2))

	// compare to not equal
	require.False(t, bool1.Equals(bool3))
	require.False(t, char1.Equals(char3))
	require.False(t, int1.Equals(int3))
	require.False(t, float1.Equals(float3))
	require.False(t, string1.Equals(string3))
	require.False(t, bytes1.Equals(bytes3))
	require.False(t, array1.Equals(array3))
	require.False(t, map1.Equals(map3))
	require.False(t, record1.Equals(record3))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.Add,
				core.NewFloat(r), core.NewFloat(l+r))
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.Sub,
				core.NewFloat(r), core.NewFloat(l-r))
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.Mul,
				core.NewFloat(r), core.NewFloat(l*r))
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, core.NewFloat(l), token.Quo,
					core.NewFloat(r), core.NewFloat(l/r))
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.Less,
				core.NewFloat(r), boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.Greater,
				core.NewFloat(r), boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.LessEq,
				core.NewFloat(r), boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, core.NewFloat(l), token.GreaterEq,
				core.NewFloat(r), boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.Add,
				core.NewInt(r), core.NewFloat(l+float64(r)))
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.Sub,
				core.NewInt(r), core.NewFloat(l-float64(r)))
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.Mul,
				core.NewInt(r), core.NewFloat(l*float64(r)))
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, core.NewFloat(l), token.Quo,
					core.NewInt(r),
					core.NewFloat(l/float64(r)))
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.Less,
				core.NewInt(r), boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.Greater,
				core.NewInt(r), boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.LessEq,
				core.NewInt(r), boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewFloat(l), token.GreaterEq,
				core.NewInt(r), boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.Add,
				core.NewInt(r), core.NewInt(l+r))
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.Sub,
				core.NewInt(r), core.NewInt(l-r))
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.Mul,
				core.NewInt(r), core.NewInt(l*r))
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, core.NewInt(l), token.Quo,
					core.NewInt(r), core.NewInt(l/r))
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, core.NewInt(l), token.Rem,
					core.NewInt(r), core.NewInt(l%r))
			}
		}
	}

	// int & int
	testBinaryOp(t,
		core.NewInt(0), token.And, core.NewInt(0),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1), token.And, core.NewInt(0),
		core.NewInt(int64(1)&int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.And, core.NewInt(1),
		core.NewInt(int64(0)&int64(1)))
	testBinaryOp(t,
		core.NewInt(1), token.And, core.NewInt(1),
		core.NewInt(int64(1)))
	testBinaryOp(t,
		core.NewInt(0), token.And, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)&int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1), token.And, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1)&int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(int64(0xffffffff)), token.And,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1984), token.And,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1984)&int64(0xffffffff)))
	testBinaryOp(t, core.NewInt(-1984), token.And,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(-1984)&int64(0xffffffff)))

	// int | int
	testBinaryOp(t,
		core.NewInt(0), token.Or, core.NewInt(0),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1), token.Or, core.NewInt(0),
		core.NewInt(int64(1)|int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.Or, core.NewInt(1),
		core.NewInt(int64(0)|int64(1)))
	testBinaryOp(t,
		core.NewInt(1), token.Or, core.NewInt(1),
		core.NewInt(int64(1)))
	testBinaryOp(t,
		core.NewInt(0), token.Or, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)|int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1), token.Or, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1)|int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(int64(0xffffffff)), token.Or,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1984), token.Or,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1984)|int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(-1984), token.Or,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(-1984)|int64(0xffffffff)))

	// int ^ int
	testBinaryOp(t,
		core.NewInt(0), token.Xor, core.NewInt(0),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1), token.Xor, core.NewInt(0),
		core.NewInt(int64(1)^int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.Xor, core.NewInt(1),
		core.NewInt(int64(0)^int64(1)))
	testBinaryOp(t,
		core.NewInt(1), token.Xor, core.NewInt(1),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.Xor, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1), token.Xor, core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1)^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(int64(0xffffffff)), token.Xor,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1984), token.Xor,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1984)^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(-1984), token.Xor,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(-1984)^int64(0xffffffff)))

	// int &^ int
	testBinaryOp(t,
		core.NewInt(0), token.AndNot, core.NewInt(0),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1), token.AndNot, core.NewInt(0),
		core.NewInt(int64(1)&^int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.AndNot,
		core.NewInt(1), core.NewInt(int64(0)&^int64(1)))
	testBinaryOp(t,
		core.NewInt(1), token.AndNot, core.NewInt(1),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(0), token.AndNot,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)&^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(1), token.AndNot,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1)&^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(int64(0xffffffff)), token.AndNot,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(0)))
	testBinaryOp(t,
		core.NewInt(1984), token.AndNot,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(1984)&^int64(0xffffffff)))
	testBinaryOp(t,
		core.NewInt(-1984), token.AndNot,
		core.NewInt(int64(0xffffffff)),
		core.NewInt(int64(-1984)&^int64(0xffffffff)))

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			core.NewInt(0), token.Shl, core.NewInt(s),
			core.NewInt(int64(0)<<uint(s)))
		testBinaryOp(t,
			core.NewInt(1), token.Shl, core.NewInt(s),
			core.NewInt(int64(1)<<uint(s)))
		testBinaryOp(t,
			core.NewInt(2), token.Shl, core.NewInt(s),
			core.NewInt(int64(2)<<uint(s)))
		testBinaryOp(t,
			core.NewInt(-1), token.Shl, core.NewInt(s),
			core.NewInt(int64(-1)<<uint(s)))
		testBinaryOp(t,
			core.NewInt(-2), token.Shl, core.NewInt(s),
			core.NewInt(int64(-2)<<uint(s)))
		testBinaryOp(t,
			core.NewInt(int64(0xffffffff)), token.Shl,
			core.NewInt(s),
			core.NewInt(int64(0xffffffff)<<uint(s)))
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			core.NewInt(0), token.Shr, core.NewInt(s),
			core.NewInt(int64(0)>>uint(s)))
		testBinaryOp(t,
			core.NewInt(1), token.Shr, core.NewInt(s),
			core.NewInt(int64(1)>>uint(s)))
		testBinaryOp(t,
			core.NewInt(2), token.Shr, core.NewInt(s),
			core.NewInt(int64(2)>>uint(s)))
		testBinaryOp(t,
			core.NewInt(-1), token.Shr, core.NewInt(s),
			core.NewInt(int64(-1)>>uint(s)))
		testBinaryOp(t,
			core.NewInt(-2), token.Shr, core.NewInt(s),
			core.NewInt(int64(-2)>>uint(s)))
		testBinaryOp(t,
			core.NewInt(int64(0xffffffff)), token.Shr,
			core.NewInt(s),
			core.NewInt(int64(0xffffffff)>>uint(s)))
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.Less,
				core.NewInt(r), boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.Greater,
				core.NewInt(r), boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.LessEq,
				core.NewInt(r), boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, core.NewInt(l), token.GreaterEq,
				core.NewInt(r), boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.Add,
				core.NewFloat(r),
				core.NewFloat(float64(l)+r))
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.Sub,
				core.NewFloat(r),
				core.NewFloat(float64(l)-r))
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.Mul,
				core.NewFloat(r),
				core.NewFloat(float64(l)*r))
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, core.NewInt(l), token.Quo,
					core.NewFloat(r),
					core.NewFloat(float64(l)/r))
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.Less,
				core.NewFloat(r), boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.Greater,
				core.NewFloat(r), boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.LessEq,
				core.NewFloat(r), boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, core.NewInt(l), token.GreaterEq,
				core.NewFloat(r), boolValue(float64(l) >= r))
		}
	}
}

func TestRecord_Index(t *testing.T) {
	m := alloc.NewRecordValue(make(map[string]core.Value), false)
	k := core.NewInt(1)
	v := alloc.NewStringValue("abcdef")
	err := m.Assign(k, v)

	require.NoError(t, err)

	res, err := m.Access(vm, k, parser.OpIndex)
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
			testBinaryOp(t, alloc.NewStringValue(ls), token.Add,
				alloc.NewStringValue(rs),
				alloc.NewStringValue(ls+rs))

			rc := []rune(rstr)[r]
			testBinaryOp(t, alloc.NewStringValue(ls), token.Add,
				core.NewChar(rc),
				alloc.NewStringValue(ls+string(rc)))
		}
	}
}

func testBinaryOp(t *testing.T, lhs core.Value, op token.Token, rhs core.Value, expected core.Value) {
	t.Helper()
	actual, err := lhs.BinaryOp(vm, op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) core.Value {
	return core.NewBool(b)
}
