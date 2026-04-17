package core

import (
	"fmt"
	"unsafe"
)

// ValuePtrValue creates new boxed value pointer value.
func ValuePtrValue(p *Value) Value {
	return Value{
		Ptr:  unsafe.Pointer(p),
		Type: VT_VALUE_PTR,
	}
}

/* ValuePtr type methods */

func valuePtrTypeName(v Value) string {
	return fmt.Sprintf("<value_ptr:%s>", v.TypeName())
}
