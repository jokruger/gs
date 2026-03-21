package stdlib

import (
	"math/rand"

	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

var randModule = map[string]core.Object{
	"int":        value.NewStaticBuiltinFunction("int", randInt63, 0, false),
	"float":      value.NewStaticBuiltinFunction("float", randFloat64, 0, false),
	"intn":       value.NewStaticBuiltinFunction("intn", randInt63n, 1, false),
	"exp_float":  value.NewStaticBuiltinFunction("exp_float", randExpFloat64, 0, false),
	"norm_float": value.NewStaticBuiltinFunction("norm_float", randNormFloat64, 0, false),
	"perm":       value.NewStaticBuiltinFunction("perm", randPerm, 1, false),
	"seed":       value.NewStaticBuiltinFunction("seed", randSeed, 1, false),
	"read":       value.NewStaticBuiltinFunction("read", randRead, 1, false),
	"rand":       value.NewStaticBuiltinFunction("rand", randFunc, 1, false),
}

func randPerm(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("rand.perm", "1", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("rand.perm", "first", "int(compatible)", args[0])
	}
	res := rand.Perm(int(i1))
	arr := make([]core.Object, 0, len(res))
	for _, v := range res {
		arr = append(arr, alloc.NewInt(int64(v)))
	}
	return alloc.NewArray(arr, false), nil
}

func randNormFloat64(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 0 {
		return nil, core.NewWrongNumArgumentsError("rand.norm_float", "0", len(args))
	}
	return alloc.NewFloat(rand.NormFloat64()), nil
}

func randExpFloat64(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 0 {
		return nil, core.NewWrongNumArgumentsError("rand.exp_float", "0", len(args))
	}
	return alloc.NewFloat(rand.ExpFloat64()), nil
}

func randFloat64(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 0 {
		return nil, core.NewWrongNumArgumentsError("rand.float", "0", len(args))
	}
	return alloc.NewFloat(rand.Float64()), nil
}

func randSeed(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("rand.seed", "1", len(args))
	}

	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("rand.seed", "first", "int(compatible)", args[0])
	}
	rand.Seed(i1)
	return alloc.NewUndefined(), nil
}

func randInt63n(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("rand.intn", "1", len(args))
	}

	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("rand.intn", "first", "int(compatible)", args[0])
	}
	return alloc.NewInt(rand.Int63n(i1)), nil
}

func randRead(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("rand.read", "1", len(args))
	}
	y1, ok := args[0].(*value.Bytes)
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("rand.read", "first", "bytes", args[0])
	}
	res, err := rand.Read(y1.Value())
	if err != nil {
		ret = wrapError(alloc, err)
		return
	}
	return alloc.NewInt(int64(res)), nil
}

func randFunc(alloc core.Allocator, args ...core.Object) (core.Object, error) {
	if len(args) != 1 {
		return nil, core.NewWrongNumArgumentsError("rand.rand", "1", len(args))
	}
	i1, ok := args[0].AsInt()
	if !ok {
		return nil, core.NewInvalidArgumentTypeError("rand.rand", "first", "int(compatible)", args[0])
	}
	src := rand.NewSource(i1)
	return randRand(alloc, rand.New(src)), nil
}

func randInt63(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
	if len(args) != 0 {
		return nil, core.NewWrongNumArgumentsError("rand.int", "0", len(args))
	}
	return alloc.NewInt(rand.Int63()), nil
}

func randRand(alloc core.Allocator, r *rand.Rand) *value.Record {
	rInt63 := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.int", "0", len(args))
		}
		return alloc.NewInt(r.Int63()), nil
	}

	rRead := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.read", "1", len(args))
		}
		y1, ok := args[0].(*value.Bytes)
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("rand.rand.read", "first", "bytes", args[0])
		}
		res, err := r.Read(y1.Value())
		if err != nil {
			ret = wrapError(alloc, err)
			return
		}
		return alloc.NewInt(int64(res)), nil
	}

	rInt63n := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.intn", "1", len(args))
		}

		i1, ok := args[0].AsInt()
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("rand.rand.intn", "first", "int(compatible)", args[0])
		}
		return alloc.NewInt(r.Int63n(i1)), nil
	}

	rSeed := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.seed", "1", len(args))
		}

		i1, ok := args[0].AsInt()
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("rand.rand.seed", "first", "int(compatible)", args[0])
		}
		r.Seed(i1)
		return alloc.NewUndefined(), nil
	}

	rFloat64 := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.float", "0", len(args))
		}
		return alloc.NewFloat(r.Float64()), nil
	}

	rExpFloat64 := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.exp_float", "0", len(args))
		}
		return alloc.NewFloat(r.ExpFloat64()), nil
	}

	rNormFloat64 := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 0 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.norm_float", "0", len(args))
		}
		return alloc.NewFloat(r.NormFloat64()), nil
	}

	rPerm := func(alloc core.Allocator, args ...core.Object) (ret core.Object, err error) {
		if len(args) != 1 {
			return nil, core.NewWrongNumArgumentsError("rand.rand.perm", "1", len(args))
		}
		i1, ok := args[0].AsInt()
		if !ok {
			return nil, core.NewInvalidArgumentTypeError("rand.rand.perm", "first", "int(compatible)", args[0])
		}
		res := r.Perm(int(i1))
		arr := make([]core.Object, 0, len(res))
		for _, v := range res {
			arr = append(arr, alloc.NewInt(int64(v)))
		}
		return alloc.NewArray(arr, false), nil
	}

	return alloc.NewRecord(map[string]core.Object{
		"int":        alloc.NewBuiltinFunction("int", rInt63, 0, false),
		"float":      alloc.NewBuiltinFunction("float", rFloat64, 0, false),
		"intn":       alloc.NewBuiltinFunction("intn", rInt63n, 1, false),
		"exp_float":  alloc.NewBuiltinFunction("exp_float", rExpFloat64, 0, false),
		"norm_float": alloc.NewBuiltinFunction("norm_float", rNormFloat64, 0, false),
		"perm":       alloc.NewBuiltinFunction("perm", rPerm, 1, false),
		"seed":       alloc.NewBuiltinFunction("seed", rSeed, 1, false),
		"read":       alloc.NewBuiltinFunction("read", rRead, 1, false),
	}, true).(*value.Record)
}
