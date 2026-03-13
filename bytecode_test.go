package gs_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/jokruger/gs"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/require"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&gs.Char{Value: 'y'},
			&gs.Float{Value: 93.11},
			compiledFunction(1, 0,
				gs.MakeInstruction(parser.OpConstant, 3),
				gs.MakeInstruction(parser.OpSetLocal, 0),
				gs.MakeInstruction(parser.OpGetGlobal, 0),
				gs.MakeInstruction(parser.OpGetFree, 0)),
			&gs.Float{Value: 39.2},
			&gs.Int{Value: 192},
			&gs.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			gs.MakeInstruction(parser.OpConstant, 0),
			gs.MakeInstruction(parser.OpSetGlobal, 0),
			gs.MakeInstruction(parser.OpConstant, 6),
			gs.MakeInstruction(parser.OpPop)),
		objectsArray(
			&gs.Int{Value: 55},
			&gs.Int{Value: 66},
			&gs.Int{Value: 77},
			&gs.Int{Value: 88},
			&gs.ImmutableMap{
				Value: map[string]gs.Object{
					"array": &gs.ImmutableArray{
						Value: []gs.Object{
							&gs.Int{Value: 1},
							&gs.Int{Value: 2},
							&gs.Int{Value: 3},
							gs.TrueValue,
							gs.FalseValue,
							gs.UndefinedValue,
						},
					},
					"true":  gs.TrueValue,
					"false": gs.FalseValue,
					"bytes": &gs.Bytes{Value: make([]byte, 16)},
					"char":  &gs.Char{Value: 'Y'},
					"error": &gs.Error{Value: &gs.String{
						Value: "some error",
					}},
					"float": &gs.Float{Value: -19.84},
					"immutable_array": &gs.ImmutableArray{
						Value: []gs.Object{
							&gs.Int{Value: 1},
							&gs.Int{Value: 2},
							&gs.Int{Value: 3},
							gs.TrueValue,
							gs.FalseValue,
							gs.UndefinedValue,
						},
					},
					"immutable_map": &gs.ImmutableMap{
						Value: map[string]gs.Object{
							"a": &gs.Int{Value: 1},
							"b": &gs.Int{Value: 2},
							"c": &gs.Int{Value: 3},
							"d": gs.TrueValue,
							"e": gs.FalseValue,
							"f": gs.UndefinedValue,
						},
					},
					"int": &gs.Int{Value: 91},
					"map": &gs.Map{
						Value: map[string]gs.Object{
							"a": &gs.Int{Value: 1},
							"b": &gs.Int{Value: 2},
							"c": &gs.Int{Value: 3},
							"d": gs.TrueValue,
							"e": gs.FalseValue,
							"f": gs.UndefinedValue,
						},
					},
					"string":    &gs.String{Value: "foo bar"},
					"time":      &gs.Time{Value: time.Now()},
					"undefined": gs.UndefinedValue,
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
				&gs.Char{Value: 'y'},
				&gs.Float{Value: 93.11},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&gs.Float{Value: 39.2},
				&gs.Int{Value: 192},
				&gs.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&gs.Char{Value: 'y'},
				&gs.Float{Value: 93.11},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&gs.Float{Value: 39.2},
				&gs.Int{Value: 192},
				&gs.String{Value: "bar"})))

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
				&gs.Int{Value: 1},
				&gs.Float{Value: 2.0},
				&gs.Char{Value: '3'},
				&gs.String{Value: "four"},
				compiledFunction(1, 0,
					gs.MakeInstruction(parser.OpConstant, 3),
					gs.MakeInstruction(parser.OpConstant, 7),
					gs.MakeInstruction(parser.OpSetLocal, 0),
					gs.MakeInstruction(parser.OpGetGlobal, 0),
					gs.MakeInstruction(parser.OpGetFree, 0)),
				&gs.Int{Value: 1},
				&gs.Float{Value: 2.0},
				&gs.Char{Value: '3'},
				&gs.String{Value: "four"})),
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
				&gs.Int{Value: 1},
				&gs.Float{Value: 2.0},
				&gs.Char{Value: '3'},
				&gs.String{Value: "four"},
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
				&gs.Int{Value: 1},
				&gs.Int{Value: 2},
				&gs.Int{Value: 3},
				&gs.Int{Value: 1},
				&gs.Int{Value: 3})),
		bytecode(
			concatInsts(
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 1),
				gs.MakeInstruction(parser.OpConstant, 2),
				gs.MakeInstruction(parser.OpConstant, 0),
				gs.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&gs.Int{Value: 1},
				&gs.Int{Value: 2},
				&gs.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&gs.Int{Value: 55},
			&gs.Int{Value: 66},
			&gs.Int{Value: 77},
			&gs.Int{Value: 88},
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
	constants []gs.Object,
	fileSet *parser.SourceFileSet,
) *gs.Bytecode {
	return &gs.Bytecode{
		FileSet:      fileSet,
		MainFunction: &gs.CompiledFunction{Instructions: instructions},
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
