package vm

import (
	"github.com/jokruger/gs/core"
)

type Module struct {
	Attrs map[string]core.Value
}

// Import returns an immutable record for the module.
func (m *Module) Import(alloc core.Allocator, moduleName string) (any, error) {
	return m.AsImmutableRecord(alloc, moduleName), nil
}

// AsImmutableRecord converts builtin module into an immutable record.
func (m *Module) AsImmutableRecord(alloc core.Allocator, moduleName string) core.Value {
	attrs := make(map[string]core.Value, len(m.Attrs))
	for k, v := range m.Attrs {
		attrs[k] = v.Copy(alloc)
	}
	attrs["__module_name__"] = alloc.NewStringValue(moduleName)
	return alloc.NewRecordValue(attrs, true)
}
