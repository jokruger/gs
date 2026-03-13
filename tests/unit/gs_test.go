package gs_test

import (
	"strings"
	"testing"
	"time"

	"github.com/jokruger/gs"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/tests/require"
	gst "github.com/jokruger/gs/types"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			gs.MakeInstruction(parser.OpConstant, 1),
			gs.MakeInstruction(parser.OpConstant, 2),
			gs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1    
0003 CONST   2    
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			gs.MakeInstruction(parser.OpBinaryOp, 11),
			gs.MakeInstruction(parser.OpConstant, 2),
			gs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 CONST   2    
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			gs.MakeInstruction(parser.OpBinaryOp, 11),
			gs.MakeInstruction(parser.OpGetLocal, 1),
			gs.MakeInstruction(parser.OpConstant, 2),
			gs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 GETL    1    
0004 CONST   2    
0007 CONST   65535`)
}

func TestMakeInstruction(t *testing.T) {
	makeInstruction(t, []byte{parser.OpConstant, 0, 0},
		parser.OpConstant, 0)
	makeInstruction(t, []byte{parser.OpConstant, 0, 1},
		parser.OpConstant, 1)
	makeInstruction(t, []byte{parser.OpConstant, 255, 254},
		parser.OpConstant, 65534)
	makeInstruction(t, []byte{parser.OpPop}, parser.OpPop)
	makeInstruction(t, []byte{parser.OpTrue}, parser.OpTrue)
	makeInstruction(t, []byte{parser.OpFalse}, parser.OpFalse)
}

func TestNumObjects(t *testing.T) {
	testCountObjects(t, &gst.Array{}, 1)
	testCountObjects(t, &gst.Array{Value: []gst.Object{
		&gst.Int{Value: 1},
		&gst.Int{Value: 2},
		&gst.Array{Value: []gst.Object{
			&gst.Int{Value: 3},
			&gst.Int{Value: 4},
			&gst.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, gst.TrueValue, 1)
	testCountObjects(t, gst.FalseValue, 1)
	testCountObjects(t, &gst.BuiltinFunction{}, 1)
	testCountObjects(t, &gst.Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &gst.Char{Value: '가'}, 1)
	testCountObjects(t, &gst.CompiledFunction{}, 1)
	testCountObjects(t, &gst.Error{Value: &gst.Int{Value: 5}}, 2)
	testCountObjects(t, &gst.Float{Value: 19.84}, 1)
	testCountObjects(t, &gst.ImmutableArray{Value: []gst.Object{
		&gst.Int{Value: 1},
		&gst.Int{Value: 2},
		&gst.ImmutableArray{Value: []gst.Object{
			&gst.Int{Value: 3},
			&gst.Int{Value: 4},
			&gst.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &gst.ImmutableMap{
		Value: map[string]gst.Object{
			"k1": &gst.Int{Value: 1},
			"k2": &gst.Int{Value: 2},
			"k3": &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 3},
				&gst.Int{Value: 4},
				&gst.Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &gst.Int{Value: 1984}, 1)
	testCountObjects(t, &gst.Map{Value: map[string]gst.Object{
		"k1": &gst.Int{Value: 1},
		"k2": &gst.Int{Value: 2},
		"k3": &gst.Array{Value: []gst.Object{
			&gst.Int{Value: 3},
			&gst.Int{Value: 4},
			&gst.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &gst.String{Value: "foo bar"}, 1)
	testCountObjects(t, &gst.Time{Value: time.Now()}, 1)
	testCountObjects(t, gst.UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o gst.Object, expected int) {
	require.Equal(t, expected, gs.CountObjects(o))
}

func assertInstructionString(
	t *testing.T,
	instructions [][]byte,
	expected string,
) {
	concatted := make([]byte, 0)
	for _, e := range instructions {
		concatted = append(concatted, e...)
	}
	require.Equal(t, expected, strings.Join(
		gs.FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := gs.MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
