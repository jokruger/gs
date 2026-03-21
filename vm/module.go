package vm

import (
	"github.com/jokruger/gs/core"
)

type Module struct {
	Attrs map[string]core.Object
}

// Import returns an immutable record for the module.
func (m *Module) Import(alloc core.Allocator, moduleName string) (any, error) {
	return m.AsImmutableRecord(alloc, moduleName), nil
}

// AsImmutableRecord converts builtin module into an immutable record.
func (m *Module) AsImmutableRecord(alloc core.Allocator, moduleName string) core.Object {
	attrs := make(map[string]core.Object, len(m.Attrs))
	for k, v := range m.Attrs {
		attrs[k] = v.Copy(alloc)
	}
	attrs["__module_name__"] = alloc.NewString(moduleName)
	return alloc.NewRecord(attrs, true)
}
