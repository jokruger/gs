package stdlib

import "github.com/jokruger/gs/core"

// BuiltinModules are builtin type standard library modules.
var BuiltinModules = map[string]map[string]core.Object{
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
