package stdlib

import (
	"math/rand"

	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/value"
)

var randModule = map[string]value.Object{
	"int": &value.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &value.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &value.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &value.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &value.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &value.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &value.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &value.UserFunction{
		Name: "read",
		Value: func(args ...value.Object) (ret value.Object, err error) {
			if len(args) != 1 {
				return nil, gse.ErrWrongNumArguments
			}
			y1, ok := args[0].(*value.Bytes)
			if !ok {
				return nil, gse.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &value.Int{Value: int64(res)}, nil
		},
	},
	"rand": &value.UserFunction{
		Name: "rand",
		Value: func(args ...value.Object) (value.Object, error) {
			if len(args) != 1 {
				return nil, gse.ErrWrongNumArguments
			}
			i1, ok := args[0].ToInt64()
			if !ok {
				return nil, gse.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "int(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *value.ImmutableMap {
	return &value.ImmutableMap{
		Value: map[string]value.Object{
			"int": &value.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &value.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &value.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &value.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &value.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &value.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &value.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &value.UserFunction{
				Name: "read",
				Value: func(args ...value.Object) (
					ret value.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					y1, ok := args[0].(*value.Bytes)
					if !ok {
						return nil, gse.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &value.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
