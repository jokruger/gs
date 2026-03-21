package unit

import (
	"bytes"
	"testing"
	"time"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/parser"
	"github.com/jokruger/gs/tests/require"
	"github.com/jokruger/gs/value"
	"github.com/jokruger/gs/vm"
)

type srcfile struct {
	name string
	size int
}

func TestBytecodeEmpty(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))
}

func TestBytecodeConstUndefined(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewUndefined(),
	)))
}

func TestBytecodeConstBool(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewBool(true),
		alloc.NewBool(false),
	)))
}

func TestBytecodeConstChar(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewChar('a'),
		alloc.NewChar('b'),
		alloc.NewChar('c'),
	)))
}

func TestBytecodeConstInt(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewInt(1),
		alloc.NewInt(2),
		alloc.NewInt(3),
		alloc.NewInt(1234567890),
	)))
}

func TestBytecodeConstFloat(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewFloat(0.123),
		alloc.NewFloat(123456.789),
	)))
}

func TestBytecodeConstString(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewString(""),
		alloc.NewString("foo"),
		alloc.NewString("foo bar"),
	)))
}

func TestBytecodeConstBytes(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewBytes([]byte{}),
		alloc.NewBytes([]byte{1, 2, 3}),
		alloc.NewBytes([]byte("foo bar")),
	)))
}

func TestBytecodeConstTime(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewTime(time.Unix(0, 0)),
		alloc.NewTime(time.Unix(1234567890, 123456789)),
	)))
}

func TestBytecodeConstArray(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewArray([]core.Object{
			alloc.NewInt(1),
			alloc.NewFloat(2.0),
			alloc.NewChar('3'),
			alloc.NewString("four"),
		}, true),
	)))

	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewArray([]core.Object{
			alloc.NewInt(1),
			alloc.NewFloat(2.0),
			alloc.NewChar('3'),
			alloc.NewString("four"),
		}, false),
	)))
}

func TestBytecodeConstMap(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewRecord(map[string]core.Object{
			"a": alloc.NewInt(1),
			"b": alloc.NewFloat(2.0),
			"c": alloc.NewChar('3'),
			"d": alloc.NewString("four"),
		}, true),
	)))

	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewRecord(map[string]core.Object{
			"a": alloc.NewInt(1),
			"b": alloc.NewFloat(2.0),
			"c": alloc.NewChar('3'),
			"d": alloc.NewString("four"),
		}, false),
	)))
}

func TestBytecodeConstError(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewError(alloc.NewString("some error")),
	)))
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			alloc.NewChar('y'),
			alloc.NewFloat(93.11),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpSetLocal, 0),
				vm.MakeInstruction(parser.OpGetGlobal, 0),
				vm.MakeInstruction(parser.OpGetFree, 0)),
			alloc.NewFloat(39.2),
			alloc.NewInt(192),
			alloc.NewString("bar"),
		)))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			vm.MakeInstruction(parser.OpConstant, 0),
			vm.MakeInstruction(parser.OpSetGlobal, 0),
			vm.MakeInstruction(parser.OpConstant, 6),
			vm.MakeInstruction(parser.OpPop)),
		objectsArray(
			alloc.NewInt(55),
			alloc.NewInt(66),
			alloc.NewInt(77),
			alloc.NewInt(88),
			alloc.NewRecord(map[string]core.Object{
				"array": alloc.NewArray([]core.Object{
					alloc.NewInt(1),
					alloc.NewInt(2),
					alloc.NewInt(3),
					alloc.NewBool(true),
					alloc.NewBool(false),
					alloc.NewUndefined(),
				}, true),
				"true":  alloc.NewBool(true),
				"false": alloc.NewBool(false),
				"bytes": alloc.NewBytes(make([]byte, 16)),
				"char":  alloc.NewChar('Y'),
				"error": alloc.NewError(alloc.NewString("some error")),
				"float": alloc.NewFloat(-19.84),
				"immutable_array": alloc.NewArray([]core.Object{
					alloc.NewInt(1),
					alloc.NewInt(2),
					alloc.NewInt(3),
					alloc.NewBool(true),
					alloc.NewBool(false),
					alloc.NewUndefined(),
				}, true),
				"immutable_map": alloc.NewRecord(map[string]core.Object{
					"a": alloc.NewInt(1),
					"b": alloc.NewInt(2),
					"c": alloc.NewInt(3),
					"d": alloc.NewBool(true),
					"e": alloc.NewBool(false),
					"f": alloc.NewUndefined(),
				}, true),
				"int": alloc.NewInt(91),
				"map": alloc.NewRecord(map[string]core.Object{
					"a": alloc.NewInt(1),
					"b": alloc.NewInt(2),
					"c": alloc.NewInt(3),
					"d": alloc.NewBool(true),
					"e": alloc.NewBool(false),
					"f": alloc.NewUndefined(),
				}, false),
				"string":    alloc.NewString("foo bar"),
				"time":      alloc.NewTime(time.Now()),
				"undefined": alloc.NewUndefined(),
			}, true),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpSetLocal, 0),
				vm.MakeInstruction(parser.OpGetGlobal, 0),
				vm.MakeInstruction(parser.OpGetFree, 0),
				vm.MakeInstruction(parser.OpBinaryOp, 11),
				vm.MakeInstruction(parser.OpGetFree, 1),
				vm.MakeInstruction(parser.OpBinaryOp, 11),
				vm.MakeInstruction(parser.OpGetLocal, 0),
				vm.MakeInstruction(parser.OpBinaryOp, 11),
				vm.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpSetLocal, 0),
				vm.MakeInstruction(parser.OpGetFree, 0),
				vm.MakeInstruction(parser.OpGetLocal, 0),
				vm.MakeInstruction(parser.OpClosure, 4, 2),
				vm.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpSetLocal, 0),
				vm.MakeInstruction(parser.OpGetLocal, 0),
				vm.MakeInstruction(parser.OpClosure, 5, 1),
				vm.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				alloc.NewChar('y'),
				alloc.NewFloat(93.11),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				alloc.NewFloat(39.2),
				alloc.NewInt(192),
				alloc.NewString("bar"))),
		bytecode(
			concatInsts(), objectsArray(
				alloc.NewChar('y'),
				alloc.NewFloat(93.11),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				alloc.NewFloat(39.2),
				alloc.NewInt(192),
				alloc.NewString("bar"))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpConstant, 4),
				vm.MakeInstruction(parser.OpConstant, 5),
				vm.MakeInstruction(parser.OpConstant, 6),
				vm.MakeInstruction(parser.OpConstant, 7),
				vm.MakeInstruction(parser.OpConstant, 8),
				vm.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				alloc.NewInt(1),
				alloc.NewFloat(2.0),
				alloc.NewChar('3'),
				alloc.NewString("four"),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpConstant, 7),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				alloc.NewInt(1),
				alloc.NewFloat(2.0),
				alloc.NewChar('3'),
				alloc.NewString("four"))),
		bytecode(
			concatInsts(
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpConstant, 4),
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				alloc.NewInt(1),
				alloc.NewFloat(2.0),
				alloc.NewChar('3'),
				alloc.NewString("four"),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpConstant, 2),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				alloc.NewInt(1),
				alloc.NewInt(2),
				alloc.NewInt(3),
				alloc.NewInt(1),
				alloc.NewInt(3))),
		bytecode(
			concatInsts(
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				alloc.NewInt(1),
				alloc.NewInt(2),
				alloc.NewInt(3))))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			alloc.NewInt(55),
			alloc.NewInt(66),
			alloc.NewInt(77),
			alloc.NewInt(88),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(instructions []byte, constants []core.Object, fileSet *parser.SourceFileSet) *vm.Bytecode {
	return &vm.Bytecode{
		FileSet:      fileSet,
		MainFunction: &value.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(t *testing.T, input, expected *vm.Bytecode) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *vm.Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &vm.Bytecode{}
	err = r.Decode(alloc, bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
