package stdlib

import (
	"fmt"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

var fmtModule = map[string]gs.Object{
	"print":   &gs.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &gs.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &gs.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &gs.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...gs.Object) (ret gs.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...gs.Object) (ret gs.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gse.ErrWrongNumArguments
	}

	format, ok := args[0].(*gs.String)
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := gs.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...gs.Object) (ret gs.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...gs.Object) (ret gs.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gse.ErrWrongNumArguments
	}

	format, ok := args[0].(*gs.String)
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := gs.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &gs.String{Value: s}, nil
}

func getPrintArgs(args ...gs.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := gs.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > gs.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
