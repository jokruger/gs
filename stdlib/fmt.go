package stdlib

import (
	"fmt"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

var fmtModule = map[string]gst.Object{
	"print":   &gst.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &gst.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &gst.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &gst.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...gst.Object) (ret gst.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...gst.Object) (ret gst.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gse.ErrWrongNumArguments
	}

	format, ok := args[0].(*gst.String)
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

func fmtPrintln(args ...gst.Object) (ret gst.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...gst.Object) (ret gst.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gse.ErrWrongNumArguments
	}

	format, ok := args[0].(*gst.String)
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
	return &gst.String{Value: s}, nil
}

func getPrintArgs(args ...gst.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		// TODO: shell we check if arg cannot be converted to string?
		s, _ := arg.ToString()
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > gst.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
