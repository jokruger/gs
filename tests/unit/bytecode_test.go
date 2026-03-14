package gs_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/jokruger/gs"
	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/tests/require"
	"github.com/jokruger/gs/value"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&value.Char{Value: 'y'},
			&value.Float{Value: 93.11},
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpSetLocal, 0),
				gs.MakeInstruction(parser.OpGetGlobal, 0),
				gs.MakeInstruction(parser.OpGetFree, 0)),
			&value.Float{Value: 39.2},
			&value.Int{Value: 192},
			&value.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			gs.MakeInstruction(parser.OpConstant, 0),
			gs.MakeInstruction(parser.OpSetGlobal, 0),
			gs.MakeInstruction(parser.OpConstant, 6),
			gs.MakeInstruction(parser.OpPop)),
		objectsArray(
			&value.Int{Value: 55},
			&value.Int{Value: 66},
			&value.Int{Value: 77},
			&value.Int{Value: 88},
			&value.ImmutableMap{
				Value: map[string]core.Object{
					"array": &value.ImmutableArray{
						Value: []core.Object{
							&value.Int{Value: 1},
							&value.Int{Value: 2},
							&value.Int{Value: 3},
							value.TrueValue,
							value.FalseValue,
							value.UndefinedValue,
						},
					},
					"true":  value.TrueValue,
					"false": value.FalseValue,
					"bytes": &value.Bytes{Value: make([]byte, 16)},
					"char":  &value.Char{Value: 'Y'},
					"error": &value.Error{Value: &value.String{
						Value: "some error",
					}},
					"float": &value.Float{Value: -19.84},
					"immutable_array": &value.ImmutableArray{
						Value: []core.Object{
							&value.Int{Value: 1},
							&value.Int{Value: 2},
							&value.Int{Value: 3},
							value.TrueValue,
							value.FalseValue,
							value.UndefinedValue,
						},
					},
					"immutable_map": &value.ImmutableMap{
						Value: map[string]core.Object{
							"a": &value.Int{Value: 1},
							"b": &value.Int{Value: 2},
							"c": &value.Int{Value: 3},
							"d": value.TrueValue,
							"e": value.FalseValue,
							"f": value.UndefinedValue,
						},
					},
					"int": &value.Int{Value: 91},
					"map": &value.Map{
						Value: map[string]core.Object{
							"a": &value.Int{Value: 1},
							"b": &value.Int{Value: 2},
							"c": &value.Int{Value: 3},
							"d": value.TrueValue,
							"e": value.FalseValue,
							"f": value.UndefinedValue,
						},
					},
					"string":    &value.String{Value: "foo bar"},
					"time":      &value.Time{Value: time.Now()},
					"undefined": value.UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpSetLocal, 0),
				gs.MakeInstruction(parser.OpGetGlobal, 0),
				gs.MakeInstruction(parser.OpGetFree, 0),
				gs.MakeInstruction(parser.OpBinaryOp, 11),
				gs.MakeInstruction(parser.OpGetFree, 1),
				gs.MakeInstruction(parser.OpBinaryOp, 11),
				gs.MakeInstruction(parser.OpGetLocal, 0),
				gs.MakeInstruction(parser.OpBinaryOp, 11),
				gs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpSetLocal, 0),
				gs.MakeInstruction(parser.OpGetFree, 0),
				gs.MakeInstruction(parser.OpGetLocal, 0),
				gs.MakeInstruction(parser.OpClosure, 4, 2),
				gs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpSetLocal, 0),
				gs.MakeInstruction(parser.OpGetLocal, 0),
				gs.MakeInstruction(parser.OpClosure, 5, 1),
				gs.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&value.Char{Value: 'y'},
				&value.Float{Value: 93.11},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&value.Float{Value: 39.2},
				&value.Int{Value: 192},
				&value.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&value.Char{Value: 'y'},
				&value.Float{Value: 93.11},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&value.Float{Value: 39.2},
				&value.Int{Value: 192},
				&value.String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpConstant, 4),
				gs.MakeInstruction(parser.OpConstant, 5),
				gs.MakeInstruction(parser.OpConstant, 6),
				gs.MakeInstruction(parser.OpConstant, 7),
				gs.MakeInstruction(parser.OpConstant, 8),
				gs.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&value.Int{Value: 1},
				&value.Float{Value: 2.0},
				&value.Char{Value: '3'},
				&value.String{Value: "four"},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpConstant, 7),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&value.Int{Value: 1},
				&value.Float{Value: 2.0},
				&value.Char{Value: '3'},
				&value.String{Value: "four"})),
		bytecode(
			concatInsts(
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpConstant, 4),
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&value.Int{Value: 1},
				&value.Float{Value: 2.0},
				&value.Char{Value: '3'},
				&value.String{Value: "four"},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpConstant, 2),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&value.Int{Value: 1},
				&value.Int{Value: 2},
				&value.Int{Value: 3},
				&value.Int{Value: 1},
				&value.Int{Value: 3})),
		bytecode(
			concatInsts(
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&value.Int{Value: 1},
				&value.Int{Value: 2},
				&value.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&value.Int{Value: 55},
			&value.Int{Value: 66},
			&value.Int{Value: 77},
			&value.Int{Value: 88},
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(
	instructions []byte,
	constants []core.Object,
	fileSet *parser.SourceFileSet,
) *gs.Bytecode {
	return &gs.Bytecode{
		FileSet:      fileSet,
		MainFunction: &value.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *gs.Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *gs.Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &gs.Bytecode{}
	err = r.Decode(bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
