package stdlib

import "github.com/jokruger/kavun/core"

// BuiltinModules are builtin type standard library modules.
var BuiltinModules = map[string]map[string]core.Value{
	"math":   mathModule,
	"os":     osModule,
	"text":   textModule,
	"times":  timesModule,
	"rand":   randModule,
	"fmt":    fmtModule,
	"json":   jsonModule,
	"base64": base64Module,
	"hex":    hexModule,
}
