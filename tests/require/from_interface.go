package require

import (
	"fmt"
	"time"

	"github.com/jokruger/gs/core"
)

func FromInterface(alloc core.Allocator, v any) (core.Object, error) {
	switch v := v.(type) {
	case nil:
		return alloc.NewUndefined(), nil
	case string:
		return alloc.NewString(v), nil
	case int64:
		return alloc.NewInt(v), nil
	case int:
		return alloc.NewInt(int64(v)), nil
	case bool:
		return alloc.NewBool(v), nil
	case rune:
		return alloc.NewChar(v), nil
	case byte:
		return alloc.NewChar(rune(v)), nil
	case float64:
		return alloc.NewFloat(v), nil
	case []byte:
		return alloc.NewBytes(v), nil
	case error:
		return alloc.NewError(alloc.NewString(v.Error())), nil
	case map[string]core.Object:
		return alloc.NewRecord(v, false), nil
	case map[string]any:
		kv := make(map[string]core.Object)
		for vk, vv := range v {
			vo, err := FromInterface(alloc, vv)
			if err != nil {
				return nil, err
			}
			kv[vk] = vo
		}
		return alloc.NewRecord(kv, false), nil
	case []core.Object:
		return alloc.NewArray(v, false), nil
	case []any:
		arr := make([]core.Object, len(v))
		for i, e := range v {
			vo, err := FromInterface(alloc, e)
			if err != nil {
				return nil, err
			}
			arr[i] = vo
		}
		return alloc.NewArray(arr, false), nil
	case time.Time:
		return alloc.NewTime(v), nil
	case core.Object:
		return v, nil
	case core.NativeFunc:
		return alloc.NewBuiltinFunction("anonymous", v, 0, true), nil
	}
	return nil, fmt.Errorf("cannot convert to object: %T", v)
}
