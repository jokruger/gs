package value

import "github.com/jokruger/gs/types"

type CompiledFunction struct {
	ObjectImpl
	Instructions  []byte
	NumLocals     int // number of local variables (including function parameters)
	NumParameters int
	VarArgs       bool
	SourceMap     map[int]types.Pos
	Free          []*ObjectPtr
}

func (o *CompiledFunction) TypeName() string {
	return "compiled-function"
}

func (o *CompiledFunction) String() string {
	return "<compiled-function>"
}

func (o *CompiledFunction) Size() int64 {
	return int64(len(o.Instructions) + len(o.SourceMap) + len(o.Free))
}

func (o *CompiledFunction) Copy() Object {
	return &CompiledFunction{
		Instructions:  append([]byte{}, o.Instructions...),
		NumLocals:     o.NumLocals,
		NumParameters: o.NumParameters,
		VarArgs:       o.VarArgs,
		Free:          append([]*ObjectPtr{}, o.Free...), // DO NOT Copy() of elements; these are variable pointers
	}
}

func (o *CompiledFunction) Equals(_ Object) bool {
	return false
}

func (o *CompiledFunction) SourcePos(ip int) types.Pos {
	for ip >= 0 {
		if p, ok := o.SourceMap[ip]; ok {
			return p
		}
		ip--
	}
	return types.NoPos
}

func (o *CompiledFunction) CanCall() bool {
	return true
}
