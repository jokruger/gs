package stdlib_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/jokruger/gs/core"
	gse "github.com/jokruger/gs/error"
	"github.com/jokruger/gs/stdlib"
	"github.com/jokruger/gs/tests/require"
	"github.com/jokruger/gs/value"
)

func TestFuncAIRSsE(t *testing.T) {
	uf := stdlib.FuncAIRSsE(func(a int) ([]string, error) {
		return []string{"foo", "bar"}, nil
	})
	ret, err := funcCall(uf, &value.Int{Value: 10})
	require.NoError(t, err)
	require.Equal(t, array(&value.String{Value: "foo"},
		&value.String{Value: "bar"}), ret)
	uf = stdlib.FuncAIRSsE(func(a int) ([]string, error) {
		return nil, errors.New("some error")
	})
	ret, err = funcCall(uf, &value.Int{Value: 10})
	require.NoError(t, err)
	require.Equal(t,
		&value.Error{Value: &value.String{Value: "some error"}}, ret)
	_, err = funcCall(uf)
	require.Equal(t, gse.ErrWrongNumArguments, err)
}

func TestFuncASRIE(t *testing.T) {
	uf := stdlib.FuncASRIE(func(a string) (int, error) { return 5, nil })
	ret, err := funcCall(uf, &value.String{Value: "foo"})
	require.NoError(t, err)
	require.Equal(t, &value.Int{Value: 5}, ret)
	uf = stdlib.FuncASRIE(func(a string) (int, error) {
		return 0, errors.New("some error")
	})
	ret, err = funcCall(uf, &value.String{Value: "foo"})
	require.NoError(t, err)
	require.Equal(t,
		&value.Error{Value: &value.String{Value: "some error"}}, ret)
	_, err = funcCall(uf)
	require.Equal(t, gse.ErrWrongNumArguments, err)
}

func TestFuncAIRS(t *testing.T) {
	uf := stdlib.FuncAIRS(func(a int) string { return strconv.Itoa(a) })
	ret, err := funcCall(uf, &value.Int{Value: 55})
	require.NoError(t, err)
	require.Equal(t, &value.String{Value: "55"}, ret)
	_, err = funcCall(uf)
	require.Equal(t, gse.ErrWrongNumArguments, err)
}

func funcCall(fn core.NativeFunc, args ...core.Object) (core.Object, error) {
	userFunc := &value.BuiltinFunction{Value: fn}
	return userFunc.Call(nil, args...)
}

func array(elements ...core.Object) *value.Array {
	return &value.Array{Value: elements}
}
