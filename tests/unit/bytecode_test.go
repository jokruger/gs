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
		core.NewUndefined(),
	)))
}

func TestBytecodeConstBool(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		core.NewBool(true),
		core.NewBool(false),
	)))
}

func TestBytecodeConstChar(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		core.NewChar('a'),
		core.NewChar('b'),
		core.NewChar('c'),
	)))
}

func TestBytecodeConstInt(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		core.NewInt(1),
		core.NewInt(2),
		core.NewInt(3),
		core.NewInt(1234567890),
	)))
}

func TestBytecodeConstFloat(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		core.NewFloat(0.123),
		core.NewFloat(123456.789),
	)))
}

func TestBytecodeConstString(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewStringValue(""),
		alloc.NewStringValue("foo"),
		alloc.NewStringValue("foo bar"),
	)))
}

func TestBytecodeConstBytes(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewBytesValue([]byte{}),
		alloc.NewBytesValue([]byte{1, 2, 3}),
		alloc.NewBytesValue([]byte("foo bar")),
	)))
}

func TestBytecodeConstTime(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewTimeValue(time.Unix(0, 0)),
		alloc.NewTimeValue(time.Unix(1234567890, 123456789)),
	)))
}

func TestBytecodeConstArray(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
			core.NewFloat(2.0),
			core.NewChar('3'),
			alloc.NewStringValue("four"),
		}, true),
	)))

	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewArrayValue([]core.Value{
			core.NewInt(1),
			core.NewFloat(2.0),
			core.NewChar('3'),
			alloc.NewStringValue("four"),
		}, false),
	)))
}

func TestBytecodeConstMap(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewRecordValue(map[string]core.Value{
			"a": core.NewInt(1),
			"b": core.NewFloat(2.0),
			"c": core.NewChar('3'),
			"d": alloc.NewStringValue("four"),
		}, true),
	)))

	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewRecordValue(map[string]core.Value{
			"a": core.NewInt(1),
			"b": core.NewFloat(2.0),
			"c": core.NewChar('3'),
			"d": alloc.NewStringValue("four"),
		}, false),
	)))
}

func TestBytecodeConstError(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray(
		alloc.NewErrorValue(alloc.NewStringValue("some error")),
	)))
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			core.NewChar('y'),
			core.NewFloat(93.11),
			compiledFunction(1, 0,
				vm.MakeInstruction(parser.OpConstant, 3),
				vm.MakeInstruction(parser.OpSetLocal, 0),
				vm.MakeInstruction(parser.OpGetGlobal, 0),
				vm.MakeInstruction(parser.OpGetFree, 0)),
			core.NewFloat(39.2),
			core.NewInt(192),
			alloc.NewStringValue("bar"),
		)))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			vm.MakeInstruction(parser.OpConstant, 0),
			vm.MakeInstruction(parser.OpSetGlobal, 0),
			vm.MakeInstruction(parser.OpConstant, 6),
			vm.MakeInstruction(parser.OpPop)),
		objectsArray(
			core.NewInt(55),
			core.NewInt(66),
			core.NewInt(77),
			core.NewInt(88),
			alloc.NewRecordValue(map[string]core.Value{
				"array": alloc.NewArrayValue([]core.Value{
					core.NewInt(1),
					core.NewInt(2),
					core.NewInt(3),
					core.NewBool(true),
					core.NewBool(false),
					core.NewUndefined(),
				}, true),
				"true":  core.NewBool(true),
				"false": core.NewBool(false),
				"bytes": alloc.NewBytesValue(make([]byte, 16)),
				"char":  core.NewChar('Y'),
				"error": alloc.NewErrorValue(alloc.NewStringValue("some error")),
				"float": core.NewFloat(-19.84),
				"immutable_array": alloc.NewArrayValue([]core.Value{
					core.NewInt(1),
					core.NewInt(2),
					core.NewInt(3),
					core.NewBool(true),
					core.NewBool(false),
					core.NewUndefined(),
				}, true),
				"immutable_map": alloc.NewRecordValue(map[string]core.Value{
					"a": core.NewInt(1),
					"b": core.NewInt(2),
					"c": core.NewInt(3),
					"d": core.NewBool(true),
					"e": core.NewBool(false),
					"f": core.NewUndefined(),
				}, true),
				"int": core.NewInt(91),
				"map": alloc.NewRecordValue(map[string]core.Value{
					"a": core.NewInt(1),
					"b": core.NewInt(2),
					"c": core.NewInt(3),
					"d": core.NewBool(true),
					"e": core.NewBool(false),
					"f": core.NewUndefined(),
				}, false),
				"string":    alloc.NewStringValue("foo bar"),
				"time":      alloc.NewTimeValue(time.Now()),
				"undefined": core.NewUndefined(),
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
				core.NewChar('y'),
				core.NewFloat(93.11),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				core.NewFloat(39.2),
				core.NewInt(192),
				alloc.NewStringValue("bar"))),
		bytecode(
			concatInsts(), objectsArray(
				core.NewChar('y'),
				core.NewFloat(93.11),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				core.NewFloat(39.2),
				core.NewInt(192),
				alloc.NewStringValue("bar"))))

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
				core.NewInt(1),
				core.NewFloat(2.0),
				core.NewChar('3'),
				alloc.NewStringValue("four"),
				compiledFunction(1, 0,
					vm.MakeInstruction(parser.OpConstant, 3),
					vm.MakeInstruction(parser.OpConstant, 7),
					vm.MakeInstruction(parser.OpSetLocal, 0),
					vm.MakeInstruction(parser.OpGetGlobal, 0),
					vm.MakeInstruction(parser.OpGetFree, 0)),
				core.NewInt(1),
				core.NewFloat(2.0),
				core.NewChar('3'),
				alloc.NewStringValue("four"))),
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
				core.NewInt(1),
				core.NewFloat(2.0),
				core.NewChar('3'),
				alloc.NewStringValue("four"),
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
				core.NewInt(1),
				core.NewInt(2),
				core.NewInt(3),
				core.NewInt(1),
				core.NewInt(3))),
		bytecode(
			concatInsts(
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 1),
				vm.MakeInstruction(parser.OpConstant, 2),
				vm.MakeInstruction(parser.OpConstant, 0),
				vm.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				core.NewInt(1),
				core.NewInt(2),
				core.NewInt(3))))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			core.NewInt(55),
			core.NewInt(66),
			core.NewInt(77),
			core.NewInt(88),
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

func bytecodeFileSet(instructions []byte, constants []core.Value, fileSet *parser.SourceFileSet) *vm.Bytecode {
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
