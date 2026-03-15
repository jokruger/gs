package stdlib

import (
	"math"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var mathModule = map[string]core.Object{
	"e":                      &value.Float{Value: math.E},
	"pi":                     &value.Float{Value: math.Pi},
	"phi":                    &value.Float{Value: math.Phi},
	"sqrt2":                  &value.Float{Value: math.Sqrt2},
	"sqrtE":                  &value.Float{Value: math.SqrtE},
	"sqrtPi":                 &value.Float{Value: math.SqrtPi},
	"sqrtPhi":                &value.Float{Value: math.SqrtPhi},
	"ln2":                    &value.Float{Value: math.Ln2},
	"log2E":                  &value.Float{Value: math.Log2E},
	"ln10":                   &value.Float{Value: math.Ln10},
	"log10E":                 &value.Float{Value: math.Log10E},
	"maxFloat32":             &value.Float{Value: math.MaxFloat32},
	"smallestNonzeroFloat32": &value.Float{Value: math.SmallestNonzeroFloat32},
	"maxFloat64":             &value.Float{Value: math.MaxFloat64},
	"smallestNonzeroFloat64": &value.Float{Value: math.SmallestNonzeroFloat64},
	"maxInt":                 &value.Int{Value: math.MaxInt},
	"minInt":                 &value.Int{Value: math.MinInt},
	"maxInt8":                &value.Int{Value: math.MaxInt8},
	"minInt8":                &value.Int{Value: math.MinInt8},
	"maxInt16":               &value.Int{Value: math.MaxInt16},
	"minInt16":               &value.Int{Value: math.MinInt16},
	"maxInt32":               &value.Int{Value: math.MaxInt32},
	"minInt32":               &value.Int{Value: math.MinInt32},
	"maxInt64":               &value.Int{Value: math.MaxInt64},
	"minInt64":               &value.Int{Value: math.MinInt64},
	"abs": &value.BuiltinFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &value.BuiltinFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &value.BuiltinFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &value.BuiltinFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &value.BuiltinFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &value.BuiltinFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &value.BuiltinFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &value.BuiltinFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &value.BuiltinFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &value.BuiltinFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &value.BuiltinFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &value.BuiltinFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &value.BuiltinFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &value.BuiltinFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &value.BuiltinFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &value.BuiltinFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &value.BuiltinFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &value.BuiltinFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &value.BuiltinFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &value.BuiltinFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &value.BuiltinFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &value.BuiltinFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &value.BuiltinFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &value.BuiltinFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &value.BuiltinFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &value.BuiltinFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &value.BuiltinFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &value.BuiltinFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &value.BuiltinFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &value.BuiltinFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &value.BuiltinFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &value.BuiltinFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &value.BuiltinFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &value.BuiltinFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &value.BuiltinFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &value.BuiltinFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &value.BuiltinFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &value.BuiltinFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &value.BuiltinFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &value.BuiltinFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &value.BuiltinFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &value.BuiltinFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &value.BuiltinFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &value.BuiltinFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &value.BuiltinFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &value.BuiltinFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &value.BuiltinFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &value.BuiltinFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &value.BuiltinFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &value.BuiltinFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &value.BuiltinFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &value.BuiltinFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &value.BuiltinFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
