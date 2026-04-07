package core

type CompiledFunction struct {
	Instructions  []byte
	NumLocals     int // number of local variables (including function parameters)
	NumParameters int
	VarArgs       bool
	SourceMap     map[int]Pos
	Free          []*Value
}

func (o *CompiledFunction) Size() int64 {
	return int64(len(o.Instructions) + len(o.SourceMap) + len(o.Free))
}

func (o *CompiledFunction) SourcePos(ip int) Pos {
	for ip >= 0 {
		if p, ok := o.SourceMap[ip]; ok {
			return p
		}
		ip--
	}
	return NoPos
}
