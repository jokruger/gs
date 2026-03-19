package vm

import (
	"github.com/jokruger/gs/core"
	"github.com/jokruger/gs/value"
)

type Module struct {
	Attrs map[string]core.Object
}

// Import returns an immutable record for the module.
func (m *Module) Import(moduleName string) (any, error) {
	return m.AsImmutableRecord(moduleName), nil
}

// AsImmutableRecord converts builtin module into an immutable record.
func (m *Module) AsImmutableRecord(moduleName string) *value.Record {
	attrs := make(map[string]core.Object, len(m.Attrs))
	for k, v := range m.Attrs {
		attrs[k] = v.Copy()
	}
	attrs["__module_name__"] = value.NewString(moduleName)
	return value.NewRecord(attrs, true)
}
