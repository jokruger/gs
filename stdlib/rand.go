package stdlib

import (
	"math/rand"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

var randModule = map[string]gs.Object{
	"int": &gs.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &gs.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &gs.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &gs.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &gs.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &gs.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &gs.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &gs.UserFunction{
		Name: "read",
		Value: func(args ...gs.Object) (ret gs.Object, err error) {
			if len(args) != 1 {
				return nil, gse.ErrWrongNumArguments
			}
			y1, ok := args[0].(*gs.Bytes)
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
			return &gs.Int{Value: int64(res)}, nil
		},
	},
	"rand": &gs.UserFunction{
		Name: "rand",
		Value: func(args ...gs.Object) (gs.Object, error) {
			if len(args) != 1 {
				return nil, gse.ErrWrongNumArguments
			}
			i1, ok := gs.ToInt64(args[0])
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

func randRand(r *rand.Rand) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			"int": &gs.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &gs.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &gs.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &gs.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &gs.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &gs.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &gs.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &gs.UserFunction{
				Name: "read",
				Value: func(args ...gs.Object) (
					ret gs.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, gse.ErrWrongNumArguments
					}
					y1, ok := args[0].(*gs.Bytes)
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
					return &gs.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
