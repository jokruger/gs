package stdlib

import (
	"math"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var mathModule = map[string]core.Object{
	"e":                      value.NewStaticFloat(math.E),
	"pi":                     value.NewStaticFloat(math.Pi),
	"phi":                    value.NewStaticFloat(math.Phi),
	"sqrt2":                  value.NewStaticFloat(math.Sqrt2),
	"sqrtE":                  value.NewStaticFloat(math.SqrtE),
	"sqrtPi":                 value.NewStaticFloat(math.SqrtPi),
	"sqrtPhi":                value.NewStaticFloat(math.SqrtPhi),
	"ln2":                    value.NewStaticFloat(math.Ln2),
	"log2E":                  value.NewStaticFloat(math.Log2E),
	"ln10":                   value.NewStaticFloat(math.Ln10),
	"log10E":                 value.NewStaticFloat(math.Log10E),
	"maxFloat32":             value.NewStaticFloat(math.MaxFloat32),
	"smallestNonzeroFloat32": value.NewStaticFloat(math.SmallestNonzeroFloat32),
	"maxFloat64":             value.NewStaticFloat(math.MaxFloat64),
	"smallestNonzeroFloat64": value.NewStaticFloat(math.SmallestNonzeroFloat64),
	"maxInt":                 value.NewStaticInt(math.MaxInt),
	"minInt":                 value.NewStaticInt(math.MinInt),
	"maxInt8":                value.NewStaticInt(math.MaxInt8),
	"minInt8":                value.NewStaticInt(math.MinInt8),
	"maxInt16":               value.NewStaticInt(math.MaxInt16),
	"minInt16":               value.NewStaticInt(math.MinInt16),
	"maxInt32":               value.NewStaticInt(math.MaxInt32),
	"minInt32":               value.NewStaticInt(math.MinInt32),
	"maxInt64":               value.NewStaticInt(math.MaxInt64),
	"minInt64":               value.NewStaticInt(math.MinInt64),

	"abs":       value.NewStaticBuiltinFunction("abs", mathAbs, 1, false),
	"acos":      value.NewStaticBuiltinFunction("acos", mathAcos, 1, false),
	"acosh":     value.NewStaticBuiltinFunction("acosh", mathAcosh, 1, false),
	"asin":      value.NewStaticBuiltinFunction("asin", mathAsin, 1, false),
	"asinh":     value.NewStaticBuiltinFunction("asinh", mathAsinh, 1, false),
	"atan":      value.NewStaticBuiltinFunction("atan", mathAtan, 1, false),
	"atan2":     value.NewStaticBuiltinFunction("atan2", mathAtan2, 2, false),
	"atanh":     value.NewStaticBuiltinFunction("atanh", mathAtanh, 1, false),
	"cbrt":      value.NewStaticBuiltinFunction("cbrt", mathCbrt, 1, false),
	"ceil":      value.NewStaticBuiltinFunction("ceil", mathCeil, 1, false),
	"copysign":  value.NewStaticBuiltinFunction("copysign", mathCopysign, 2, false),
	"cos":       value.NewStaticBuiltinFunction("cos", mathCos, 1, false),
	"cosh":      value.NewStaticBuiltinFunction("cosh", mathCosh, 1, false),
	"dim":       value.NewStaticBuiltinFunction("dim", mathDim, 2, false),
	"erf":       value.NewStaticBuiltinFunction("erf", mathErf, 1, false),
	"erfc":      value.NewStaticBuiltinFunction("erfc", mathErfc, 1, false),
	"exp":       value.NewStaticBuiltinFunction("exp", mathExp, 1, false),
	"exp2":      value.NewStaticBuiltinFunction("exp2", mathExp2, 1, false),
	"expm1":     value.NewStaticBuiltinFunction("expm1", mathExpm1, 1, false),
	"floor":     value.NewStaticBuiltinFunction("floor", mathFloor, 1, false),
	"gamma":     value.NewStaticBuiltinFunction("gamma", mathGamma, 1, false),
	"hypot":     value.NewStaticBuiltinFunction("hypot", mathHypot, 2, false),
	"ilogb":     value.NewStaticBuiltinFunction("ilogb", mathIlogb, 1, false),
	"inf":       value.NewStaticBuiltinFunction("inf", mathInf, 1, false),
	"is_inf":    value.NewStaticBuiltinFunction("is_inf", mathIsInf, 2, false),
	"is_nan":    value.NewStaticBuiltinFunction("is_nan", mathIsNaN, 1, false),
	"j0":        value.NewStaticBuiltinFunction("j0", mathJ0, 1, false),
	"j1":        value.NewStaticBuiltinFunction("j1", mathJ1, 1, false),
	"jn":        value.NewStaticBuiltinFunction("jn", mathJn, 2, false),
	"ldexp":     value.NewStaticBuiltinFunction("ldexp", mathLdexp, 2, false),
	"log":       value.NewStaticBuiltinFunction("log", mathLog, 1, false),
	"log10":     value.NewStaticBuiltinFunction("log10", mathLog10, 1, false),
	"log1p":     value.NewStaticBuiltinFunction("log1p", mathLog1p, 1, false),
	"log2":      value.NewStaticBuiltinFunction("log2", mathLog2, 1, false),
	"logb":      value.NewStaticBuiltinFunction("logb", mathLogb, 1, false),
	"max":       value.NewStaticBuiltinFunction("max", mathMax, 2, false),
	"min":       value.NewStaticBuiltinFunction("min", mathMin, 2, false),
	"mod":       value.NewStaticBuiltinFunction("mod", mathMod, 2, false),
	"nan":       value.NewStaticBuiltinFunction("nan", mathNaN, 0, false),
	"nextafter": value.NewStaticBuiltinFunction("nextafter", mathNextafter, 2, false),
	"pow":       value.NewStaticBuiltinFunction("pow", mathPow, 2, false),
	"pow10":     value.NewStaticBuiltinFunction("pow10", mathPow10, 1, false),
	"remainder": value.NewStaticBuiltinFunction("remainder", mathRemainder, 2, false),
	"signbit":   value.NewStaticBuiltinFunction("signbit", mathSignbit, 1, false),
	"sin":       value.NewStaticBuiltinFunction("sin", mathSin, 1, false),
	"sinh":      value.NewStaticBuiltinFunction("sinh", mathSinh, 1, false),
	"sqrt":      value.NewStaticBuiltinFunction("sqrt", mathSqrt, 1, false),
	"tan":       value.NewStaticBuiltinFunction("tan", mathTan, 1, false),
	"tanh":      value.NewStaticBuiltinFunction("tanh", mathTanh, 1, false),
	"trunc":     value.NewStaticBuiltinFunction("trunc", mathTrunc, 1, false),
	"y0":        value.NewStaticBuiltinFunction("y0", mathY0, 1, false),
	"y1":        value.NewStaticBuiltinFunction("y1", mathY1, 1, false),
	"yn":        value.NewStaticBuiltinFunction("yn", mathYn, 2, false),
}

func mathSignbit(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.signbit", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.signbit", "first", "float(compatible)", args[0])
	}
	return alloc.NewBool(math.Signbit(f1)), nil
}

func mathIsNaN(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.is_nan", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.is_nan", "first", "float(compatible)", args[0])
	}
	return alloc.NewBool(math.IsNaN(f1)), nil
}

func mathIsInf(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.is_inf", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.is_inf", "first", "float(compatible)", args[0])
	}
	i2, ok := args[1].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.is_inf", "second", "int(compatible)", args[1])
	}
	return alloc.NewBool(math.IsInf(f1, int(i2))), nil
}

func mathLdexp(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.ldexp", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.ldexp", "first", "float(compatible)", args[0])
	}
	i2, ok := args[1].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.ldexp", "second", "int(compatible)", args[1])
	}
	return alloc.NewFloat(math.Ldexp(f1, int(i2))), nil
}

func mathYn(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.yn", "2", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.yn", "first", "int(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.yn", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Yn(int(i1), f2)), nil
}

func mathJn(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.jn", "2", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.jn", "first", "int(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.jn", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Jn(int(i1), f2)), nil
}

func mathIlogb(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.ilogb", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.ilogb", "first", "float(compatible)", args[0])
	}
	return alloc.NewInt(int64(math.Ilogb(f1))), nil
}

func mathPow10(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.pow10", "1", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.pow10", "first", "int(compatible)", args[0])
	}
	return alloc.NewFloat(math.Pow10(int(i1))), nil
}

func mathInf(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.inf", "1", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.inf", "first", "int(compatible)", args[0])
	}
	return alloc.NewFloat(math.Inf(int(i1))), nil
}

func mathAbs(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.abs", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.abs", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Abs(f1)), nil
}

func mathAcos(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.acos", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.acos", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Acos(f1)), nil
}

func mathAcosh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.acosh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.acosh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Acosh(f1)), nil
}

func mathAsin(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.asin", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.asin", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Asin(f1)), nil
}

func mathAsinh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.asinh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.asinh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Asinh(f1)), nil
}

func mathAtan(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.atan", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.atan", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Atan(f1)), nil
}

func mathAtanh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.atanh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.atanh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Atanh(f1)), nil
}

func mathCbrt(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.cbrt", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.cbrt", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Cbrt(f1)), nil
}

func mathCeil(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.ceil", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.ceil", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Ceil(f1)), nil
}

func mathCos(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.cos", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.cos", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Cos(f1)), nil
}

func mathCosh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.cosh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.cosh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Cosh(f1)), nil
}

func mathErf(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.erf", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.erf", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Erf(f1)), nil
}

func mathErfc(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.erfc", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.erfc", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Erfc(f1)), nil
}

func mathExp(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.exp", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.exp", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Exp(f1)), nil
}

func mathExp2(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.exp2", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.exp2", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Exp2(f1)), nil
}

func mathExpm1(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.expm1", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.expm1", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Expm1(f1)), nil
}

func mathFloor(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.floor", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.floor", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Floor(f1)), nil
}

func mathGamma(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.gamma", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.gamma", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Gamma(f1)), nil
}

func mathJ0(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.j0", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.j0", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.J0(f1)), nil
}

func mathJ1(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.j1", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.j1", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.J1(f1)), nil
}

func mathLog(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.log", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.log", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Log(f1)), nil
}

func mathLog10(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.log10", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.log10", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Log10(f1)), nil
}

func mathLog1p(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.log1p", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.log1p", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Log1p(f1)), nil
}

func mathLog2(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.log2", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.log2", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Log2(f1)), nil
}

func mathLogb(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.logb", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.logb", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Logb(f1)), nil
}

func mathSin(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.sin", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.sin", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Sin(f1)), nil
}

func mathSinh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.sinh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.sinh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Sinh(f1)), nil
}

func mathSqrt(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.sqrt", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.sqrt", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Sqrt(f1)), nil
}

func mathTan(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.tan", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.tan", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Tan(f1)), nil
}

func mathTanh(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.tanh", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.tanh", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Tanh(f1)), nil
}

func mathTrunc(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.trunc", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.trunc", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Trunc(f1)), nil
}

func mathY0(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.y0", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.y0", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Y0(f1)), nil
}

func mathY1(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("math.y1", "1", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.y1", "first", "float(compatible)", args[0])
	}
	return alloc.NewFloat(math.Y1(f1)), nil
}

func mathAtan2(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.atan2", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.atan2", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.atan2", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Atan2(f1, f2)), nil
}

func mathCopysign(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.copysign", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.copysign", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.copysign", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Copysign(f1, f2)), nil
}

func mathDim(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.dim", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.dim", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.dim", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Dim(f1, f2)), nil
}

func mathHypot(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.hypot", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.hypot", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.hypot", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Hypot(f1, f2)), nil
}

func mathMax(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.max", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.max", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.max", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Max(f1, f2)), nil
}

func mathMin(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.min", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.min", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.min", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Min(f1, f2)), nil
}

func mathMod(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.mod", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.mod", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.mod", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Mod(f1, f2)), nil
}

func mathNextafter(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.nextafter", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.nextafter", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.nextafter", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Nextafter(f1, f2)), nil
}

func mathPow(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.pow", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.pow", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.pow", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Pow(f1, f2)), nil
}

func mathRemainder(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 2 {
		return nil, core.NewWrongNumArgumentsError("math.remainder", "2", len(args))
	}
	f1, ok := args[0].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.remainder", "first", "float(compatible)", args[0])
	}
	f2, ok := args[1].AsFloat()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("math.remainder", "second", "float(compatible)", args[1])
	}
	return alloc.NewFloat(math.Remainder(f1, f2)), nil
}

func mathNaN(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 0 {
		return nil, core.NewWrongNumArgumentsError("math.nan", "0", len(args))
	}
	return alloc.NewFloat(math.NaN()), nil
}
