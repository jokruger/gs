package stdlib

import (
	"fmt"

	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/errs"
	"github.com/jokruger/kavun/formatter"
)

var fmtModule = map[string]core.Value{
	"print":   core.NewBuiltinFunctionValue("print", fmtPrint, 0, true),
	"printf":  core.NewBuiltinFunctionValue("printf", fmtPrintf, 1, true),
	"println": core.NewBuiltinFunctionValue("println", fmtPrintln, 0, true),
	"sprintf": core.NewBuiltinFunctionValue("sprintf", fmtSprintf, 1, true),
}

func fmtPrint(vm core.VM, args []core.Value) (core.Value, error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return core.Undefined, err
	}
	_, _ = fmt.Print(printArgs...)
	return core.Undefined, nil
}

func fmtPrintln(vm core.VM, args []core.Value) (core.Value, error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return core.Undefined, err
	}
	_, _ = fmt.Println(printArgs...)
	return core.Undefined, nil
}

func getPrintArgs(args ...core.Value) ([]any, error) {
	printArgs := make([]any, 0, len(args))
	for _, arg := range args {
		switch arg.Type {
		case core.VT_UNDEFINED, core.VT_BYTES, core.VT_ARRAY, core.VT_RECORD, core.VT_DICT, core.VT_INT_RANGE:
			printArgs = append(printArgs, arg.String())

		default:
			s, ok := arg.AsString()
			if !ok {
				s = arg.String()
			}
			printArgs = append(printArgs, s)
		}
	}

	return printArgs, nil
}

func fmtPrintf(vm core.VM, args []core.Value) (core.Value, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return core.Undefined, errs.NewWrongNumArgumentsError("fmt.printf", "at least 1", numArgs)
	}

	format, ok := args[0].AsString()
	if !ok {
		return core.Undefined, errs.NewInvalidArgumentTypeError("fmt.printf", "format", "string", args[0].TypeName())
	}
	if numArgs == 1 {
		fmt.Print(format)
		return core.Undefined, nil
	}

	s, err := formatter.Format(format, args[1:]...)
	if err != nil {
		return core.Undefined, err
	}
	fmt.Print(s)
	return core.Undefined, nil
}

func fmtSprintf(vm core.VM, args []core.Value) (core.Value, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return core.Undefined, errs.NewWrongNumArgumentsError("fmt.sprintf", "at least 1", numArgs)
	}

	format, ok := args[0].AsString()
	if !ok {
		return core.Undefined, errs.NewInvalidArgumentTypeError("fmt.sprintf", "format", "string", args[0].TypeName())
	}
	if numArgs == 1 {
		return vm.Allocator().NewStringValue(format), nil
	}
	s, err := formatter.Format(format, args[1:]...)
	if err != nil {
		return core.Undefined, err
	}
	return vm.Allocator().NewStringValue(s), nil
}
