package gs

import (
	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

var builtinFuncs = []*gst.BuiltinFunction{
	{Name: "len", Value: builtinLen},
	{Name: "copy", Value: builtinCopy},
	{Name: "append", Value: builtinAppend},
	{Name: "delete", Value: builtinDelete},
	{Name: "splice", Value: builtinSplice},
	{Name: "string", Value: builtinString},
	{Name: "int", Value: builtinInt},
	{Name: "bool", Value: builtinBool},
	{Name: "float", Value: builtinFloat},
	{Name: "char", Value: builtinChar},
	{Name: "bytes", Value: builtinBytes},
	{Name: "time", Value: builtinTime},
	{Name: "is_int", Value: builtinIsInt},
	{Name: "is_float", Value: builtinIsFloat},
	{Name: "is_string", Value: builtinIsString},
	{Name: "is_bool", Value: builtinIsBool},
	{Name: "is_char", Value: builtinIsChar},
	{Name: "is_bytes", Value: builtinIsBytes},
	{Name: "is_array", Value: builtinIsArray},
	{Name: "is_immutable_array", Value: builtinIsImmutableArray},
	{Name: "is_map", Value: builtinIsMap},
	{Name: "is_immutable_map", Value: builtinIsImmutableMap},
	{Name: "is_iterable", Value: builtinIsIterable},
	{Name: "is_time", Value: builtinIsTime},
	{Name: "is_error", Value: builtinIsError},
	{Name: "is_undefined", Value: builtinIsUndefined},
	{Name: "is_function", Value: builtinIsFunction},
	{Name: "is_callable", Value: builtinIsCallable},
	{Name: "type_name", Value: builtinTypeName},
	{Name: "format", Value: builtinFormat},
	{Name: "range", Value: builtinRange},
}

// GetAllBuiltinFunctions returns all builtin function objects.
func GetAllBuiltinFunctions() []*gst.BuiltinFunction {
	return append([]*gst.BuiltinFunction{}, builtinFuncs...)
}

func builtinTypeName(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	return &gst.String{Value: args[0].TypeName()}, nil
}

func builtinIsString(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.String); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsInt(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Int); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsFloat(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Float); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsBool(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Bool); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsChar(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Char); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsBytes(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Bytes); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsArray(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Array); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsImmutableArray(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.ImmutableArray); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsMap(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Map); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsImmutableMap(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.ImmutableMap); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsTime(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Time); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsError(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Error); ok {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsUndefined(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if args[0] == gst.UndefinedValue {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsFunction(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	switch args[0].(type) {
	case *gst.CompiledFunction:
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsCallable(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if args[0].CanCall() {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

func builtinIsIterable(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if args[0].CanIterate() {
		return gst.TrueValue, nil
	}
	return gst.FalseValue, nil
}

// len(obj object) => int
func builtinLen(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *gst.Array:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	case *gst.ImmutableArray:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	case *gst.String:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	case *gst.Bytes:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	case *gst.Map:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	case *gst.ImmutableMap:
		return &gst.Int{Value: int64(len(arg.Value))}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array/string/bytes/map",
			Found:    arg.TypeName(),
		}
	}
}

// range(start, stop[, step])
func builtinRange(args ...gst.Object) (gst.Object, error) {
	numArgs := len(args)
	if numArgs < 2 || numArgs > 3 {
		return nil, gse.ErrWrongNumArguments
	}
	var start, stop, step *gst.Int

	for i, arg := range args {
		v, ok := args[i].(*gst.Int)
		if !ok {
			var name string
			switch i {
			case 0:
				name = "start"
			case 1:
				name = "stop"
			case 2:
				name = "step"
			}

			return nil, gse.ErrInvalidArgumentType{
				Name:     name,
				Expected: "int",
				Found:    arg.TypeName(),
			}
		}
		if i == 2 && v.Value <= 0 {
			return nil, gse.ErrInvalidRangeStep
		}
		switch i {
		case 0:
			start = v
		case 1:
			stop = v
		case 2:
			step = v
		}
	}

	if step == nil {
		step = &gst.Int{Value: int64(1)}
	}

	return buildRange(start.Value, stop.Value, step.Value), nil
}

func buildRange(start, stop, step int64) *gst.Array {
	array := &gst.Array{}
	if start <= stop {
		for i := start; i < stop; i += step {
			array.Value = append(array.Value, &gst.Int{
				Value: i,
			})
		}
	} else {
		for i := start; i > stop; i -= step {
			array.Value = append(array.Value, &gst.Int{
				Value: i,
			})
		}
	}
	return array
}

func builtinFormat(args ...gst.Object) (gst.Object, error) {
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
	s, err := Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &gst.String{Value: s}, nil
}

func builtinCopy(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	return args[0].Copy(), nil
}

func builtinString(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.String); ok {
		return args[0], nil
	}
	v, ok := args[0].ToString()
	if ok {
		if len(v) > gst.MaxStringLen {
			return nil, gse.ErrStringLimit
		}
		return &gst.String{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

func builtinInt(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Int); ok {
		return args[0], nil
	}
	v, ok := args[0].ToInt64()
	if ok {
		return &gst.Int{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

func builtinFloat(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Float); ok {
		return args[0], nil
	}
	v, ok := args[0].ToFloat64()
	if ok {
		return &gst.Float{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

func builtinBool(args ...gst.Object) (gst.Object, error) {
	if len(args) != 1 {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Bool); ok {
		return args[0], nil
	}
	v, ok := args[0].ToBool()
	if ok {
		if v {
			return gst.TrueValue, nil
		}
		return gst.FalseValue, nil
	}
	return gst.UndefinedValue, nil
}

func builtinChar(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Char); ok {
		return args[0], nil
	}
	v, ok := args[0].ToRune()
	if ok {
		return &gst.Char{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

func builtinBytes(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}

	// bytes(N) => create a new bytes with given size N
	if n, ok := args[0].(*gst.Int); ok {
		if n.Value > int64(gst.MaxBytesLen) {
			return nil, gse.ErrBytesLimit
		}
		return &gst.Bytes{Value: make([]byte, int(n.Value))}, nil
	}
	v, ok := args[0].ToByteSlice()
	if ok {
		if len(v) > gst.MaxBytesLen {
			return nil, gse.ErrBytesLimit
		}
		return &gst.Bytes{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

func builtinTime(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, gse.ErrWrongNumArguments
	}
	if _, ok := args[0].(*gst.Time); ok {
		return args[0], nil
	}
	v, ok := args[0].ToTime()
	if ok {
		return &gst.Time{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return gst.UndefinedValue, nil
}

// append(arr, items...)
func builtinAppend(args ...gst.Object) (gst.Object, error) {
	if len(args) < 2 {
		return nil, gse.ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *gst.Array:
		return &gst.Array{Value: append(arg.Value, args[1:]...)}, nil
	case *gst.ImmutableArray:
		return &gst.Array{Value: append(arg.Value, args[1:]...)}, nil
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    arg.TypeName(),
		}
	}
}

// builtinDelete deletes Map keys
// usage: delete(map, "key")
// key must be a string
func builtinDelete(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, gse.ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *gst.Map:
		if key, ok := args[1].(*gst.String); ok {
			delete(arg.Value, key.Value)
			return gst.UndefinedValue, nil
		}
		return nil, gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	default:
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    arg.TypeName(),
		}
	}
}

// builtinSplice deletes and changes given Array, returns deleted items.
// usage:
// deleted_items := splice(array[,start[,delete_count[,item1[,item2[,...]]]])
func builtinSplice(args ...gst.Object) (gst.Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return nil, gse.ErrWrongNumArguments
	}

	array, ok := args[0].(*gst.Array)
	if !ok {
		return nil, gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	arrayLen := len(array.Value)

	var startIdx int
	if argsLen > 1 {
		arg1, ok := args[1].(*gst.Int)
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		startIdx = int(arg1.Value)
		if startIdx < 0 || startIdx > arrayLen {
			return nil, gse.ErrIndexOutOfBounds
		}
	}

	delCount := len(array.Value)
	if argsLen > 2 {
		arg2, ok := args[2].(*gst.Int)
		if !ok {
			return nil, gse.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		delCount = int(arg2.Value)
		if delCount < 0 {
			return nil, gse.ErrIndexOutOfBounds
		}
	}
	// if count of to be deleted items is bigger than expected, truncate it
	if startIdx+delCount > arrayLen {
		delCount = arrayLen - startIdx
	}
	// delete items
	endIdx := startIdx + delCount
	deleted := append([]gst.Object{}, array.Value[startIdx:endIdx]...)

	head := array.Value[:startIdx]
	var items []gst.Object
	if argsLen > 3 {
		items = make([]gst.Object, 0, argsLen-3)
		for i := 3; i < argsLen; i++ {
			items = append(items, args[i])
		}
	}
	items = append(items, array.Value[endIdx:]...)
	array.Value = append(head, items...)

	// return deleted items
	return &gst.Array{Value: deleted}, nil
}
