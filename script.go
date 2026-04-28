package kavun

import (
	"context"
	"fmt"
	"maps"
	"path/filepath"
	"sync"

	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/parser"
	"github.com/jokruger/kavun/vm"
)

// Script can simplify compilation and execution of embedded scripts.
type Script struct {
	variables        map[string]*Variable
	modules          vm.ModuleGetter
	input            []byte
	maxConstObjects  int
	maxFrames        int
	maxStack         int
	assignmentMode   AssignmentMode
	importDir        string
	enableFileImport bool
}

// NewScript creates a Script instance with an input script.
func NewScript(input []byte) *Script {
	return &Script{
		variables:       make(map[string]*Variable),
		input:           input,
		maxConstObjects: -1,
		maxFrames:       -1,
		maxStack:        -1,
		assignmentMode:  AssignmentModeSmart,
	}
}

// Add adds a new variable or updates an existing variable to the script.
func (s *Script) Add(name string, val core.Value) {
	s.variables[name] = NewVariable(name, val)
}

// Remove removes (undefine) an existing variable for the script. It returns false if the variable name is not defined.
func (s *Script) Remove(name string) bool {
	if _, ok := s.variables[name]; !ok {
		return false
	}
	delete(s.variables, name)
	return true
}

// SetImports sets import modules.
func (s *Script) SetImports(modules vm.ModuleGetter) {
	s.modules = modules
}

// SetImportDir sets the initial import directory for script files.
func (s *Script) SetImportDir(dir string) error {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	s.importDir = dir
	return nil
}

// SetMaxConstObjects sets the maximum number of objects in the compiled constants.
func (s *Script) SetMaxConstObjects(n int) {
	s.maxConstObjects = n
}

// SetMaxFrames sets the maximum number of frames for the compiled script.
func (s *Script) SetMaxFrames(n int) {
	s.maxFrames = n
}

// SetMaxStack sets the maximum stack size for the compiled script.
func (s *Script) SetMaxStack(n int) {
	s.maxStack = n
}

// SetAssignmentMode sets how plain '=' handles unresolved identifiers during compilation.
func (s *Script) SetAssignmentMode(mode AssignmentMode) {
	s.assignmentMode = mode
}

// EnableFileImport enables or disables module loading from local files. Local file modules are disabled by default.
func (s *Script) EnableFileImport(enable bool) {
	s.enableFileImport = enable
}

// Compile compiles the script with all the defined variables, and, returns Compiled object.
// Receives compile-time and runtime arenas for better memory management. If arenas are not provided, it will create new ones internally.
func (s *Script) Compile(cta *core.Arena, rta *core.Arena) (*Compiled, error) {
	if cta == nil {
		cta = core.NewArena(nil)
	}
	if rta == nil {
		rta = core.NewArena(nil)
	}

	symbolTable, globals, err := s.prepCompile()
	if err != nil {
		return nil, err
	}

	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile("(main)", -1, len(s.input))
	p := parser.NewParser(srcFile, s.input, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	c := NewCompiler(cta, srcFile, symbolTable, nil, s.modules, nil)
	c.SetAssignmentMode(s.assignmentMode)
	c.EnableFileImport(s.enableFileImport)
	c.SetImportDir(s.importDir)
	if err := c.Compile(file); err != nil {
		return nil, err
	}

	// reduce globals size
	globals = globals[:symbolTable.MaxSymbols()+1]

	// global symbol names to indexes
	globalIndexes := make(map[string]int, len(globals))
	for _, name := range symbolTable.Names() {
		symbol, _, _ := symbolTable.Resolve(name, false)
		if symbol.Scope == vm.ScopeGlobal {
			globalIndexes[name] = symbol.Index
		}
	}

	// remove duplicates from constants
	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()

	// check the constant objects limit
	if s.maxConstObjects >= 0 {
		cnt := bytecode.CountObjects()
		if cnt > s.maxConstObjects {
			return nil, fmt.Errorf("exceeding constant objects limit: %d", cnt)
		}
	}

	return &Compiled{
		alloc:         rta,
		bytecode:      bytecode,
		globalIndexes: globalIndexes,
		globals:       globals,
		maxFrames:     s.maxFrames,
		maxStack:      s.maxStack,
	}, nil
}

// Run compiles and runs the script, and returns the compiled instance.
// Note: prefer to use Compile() and Run() separately for better performance and control over the execution.
func (s *Script) Run() (*Compiled, error) {
	compiled, err := s.Compile(nil, nil)
	if err != nil {
		return nil, err
	}
	if err := compiled.Run(); err != nil {
		return nil, err
	}
	return compiled, nil
}

func (s *Script) prepCompile() (symbolTable *vm.SymbolTable, globals []core.Value, err error) {
	names := make([]string, 0, len(s.variables))
	for name := range s.variables {
		names = append(names, name)
	}

	symbolTable = vm.NewSymbolTable()
	for idx, fn := range vm.BuiltinFuncs {
		// it is safe to cast type because we know that all values in vm.BuiltinFuncs are *value.BuiltinFunction objects
		symbolTable.DefineBuiltin(idx, (*core.BuiltinFunction)(fn.Ptr).Name)
	}

	globals = make([]core.Value, vm.GlobalsSize)
	for idx, name := range names {
		symbol := symbolTable.Define(name)
		if symbol.Index != idx {
			panic(fmt.Errorf("wrong symbol index: %d != %d", idx, symbol.Index))
		}
		globals[symbol.Index] = s.variables[name].Value()
	}
	return
}

// Compiled is a compiled instance of the user script. Use Script.Compile() to create Compiled object.
type Compiled struct {
	lock           sync.RWMutex
	alloc          *core.Arena
	bytecode       *vm.Bytecode
	globalIndexes  map[string]int // global symbol name to index
	globals        []core.Value
	globalsRuntime []core.Value // global variables during execution
	maxFrames      int
	maxStack       int
	vm             *vm.VM
}

// Run executes the compiled script in the virtual machine.
func (c *Compiled) Run() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if err := c.prepareRun(); err != nil {
		return err
	}

	return c.vm.Run()
}

// RunContext is like Run but includes a context.
func (c *Compiled) RunContext(ctx context.Context) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if err := c.prepareRun(); err != nil {
		return err
	}

	ch := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch e := r.(type) {
				case string:
					ch <- fmt.Errorf("%s", e)
				case error:
					ch <- e
				default:
					ch <- fmt.Errorf("unknown panic: %v", e)
				}
			}
		}()
		ch <- c.vm.Run()
	}()

	select {
	case <-ctx.Done():
		c.vm.Abort()
		<-ch
		err = ctx.Err()
	case err = <-ch:
	}
	return
}

// Size of compiled script in bytes (as much as we can calculate it without reflection and black magic)
func (c *Compiled) Size() int64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.bytecode.Size() + int64(len(c.globalIndexes)+len(c.globals))
}

// Clone creates a new copy of Compiled. Cloned copies are safe for concurrent use by multiple goroutines.
func (c *Compiled) Clone(a *core.Arena) (*Compiled, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	globalIndexes := make(map[string]int, len(c.globalIndexes))
	maps.Copy(globalIndexes, c.globalIndexes)

	globals := make([]core.Value, len(c.globals))
	for i, v := range c.globals {
		t, err := v.Copy(a)
		if err != nil {
			return nil, err
		}
		globals[i] = t
	}

	clone := &Compiled{
		alloc:         a,
		globalIndexes: globalIndexes,
		bytecode:      c.bytecode,
		globals:       globals,
		maxFrames:     c.maxFrames,
		maxStack:      c.maxStack,
	}

	return clone, nil
}

// IsDefined returns true if the variable name is defined (has value) before or after the execution.
func (c *Compiled) IsDefined(name string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	idx, ok := c.globalIndexes[name]
	if !ok {
		return false
	}
	if c.globalsRuntime != nil {
		v := c.globalsRuntime[idx]
		return v.Type != core.VT_UNDEFINED
	}
	v := c.globals[idx]
	return v.Type != core.VT_UNDEFINED
}

// GetValue returns a value identified by the name.
func (c *Compiled) GetValue(name string) core.Value {
	c.lock.RLock()
	defer c.lock.RUnlock()

	v := core.Undefined
	if c.globalsRuntime != nil {
		// if the script has been executed, get the variable value from the runtime globals
		if idx, ok := c.globalIndexes[name]; ok {
			v = c.globalsRuntime[idx]
		}
	} else {
		// if the script has not been executed, get the variable value from the compile-time globals
		if idx, ok := c.globalIndexes[name]; ok {
			v = c.globals[idx]
		}
	}

	return v
}

// Get returns a variable identified by the name.
func (c *Compiled) Get(name string) *Variable {
	v := c.GetValue(name)
	return NewVariable(name, v)
}

// GetAll returns all the variables that are defined by the compiled script.
func (c *Compiled) GetAll() []*Variable {
	c.lock.RLock()
	defer c.lock.RUnlock()

	vars := make([]*Variable, 0, len(c.globalIndexes))
	if c.globalsRuntime != nil {
		for name, idx := range c.globalIndexes {
			v := c.globalsRuntime[idx]
			vars = append(vars, NewVariable(name, v))
		}
	} else {
		for name, idx := range c.globalIndexes {
			v := c.globals[idx]
			vars = append(vars, NewVariable(name, v))
		}
	}

	return vars
}

// Set replaces the value of a global variable identified by the name.
// An error will be returned if the name was not defined during compilation.
func (c *Compiled) Set(name string, val core.Value) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	idx, ok := c.globalIndexes[name]
	if !ok {
		return fmt.Errorf("'%s' is not defined", name)
	}
	c.globals[idx] = val
	return nil
}

func (c *Compiled) prepareRun() error {
	// first run
	if c.vm == nil {
		c.globalsRuntime = make([]core.Value, len(c.globals))
		for i, v := range c.globals {
			t, err := v.Copy(c.alloc)
			if err != nil {
				return err
			}
			c.globalsRuntime[i] = t
		}
		c.vm = vm.NewVM(c.alloc, c.bytecode, c.globalsRuntime, c.maxFrames, c.maxStack)
		return nil
	}

	// subsequent runs
	c.alloc.Reset()
	for i, v := range c.globals {
		t, err := v.Copy(c.alloc)
		if err != nil {
			return err
		}
		c.globalsRuntime[i] = t
	}
	c.vm.Reset(c.bytecode, c.globalsRuntime)

	return nil
}
