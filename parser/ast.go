package parser

import (
	"strings"

	"github.com/jokruger/kavun/core"
)

const (
	nullRep = "<null>"
)

// Node represents a node in the AST.
type Node interface {
	// Pos returns the position of first character belonging to the node.
	Pos() core.Pos
	// End returns the position of first character immediately after the node.
	End() core.Pos
	// String returns a string representation of the node.
	String() string
}

// IdentList represents a list of identifiers.
type IdentList struct {
	LParen  core.Pos
	VarArgs bool
	List    []*Ident
	RParen  core.Pos
}

// Pos returns the position of first character belonging to the node.
func (n *IdentList) Pos() core.Pos {
	if n.LParen.IsValid() {
		return n.LParen
	}
	if len(n.List) > 0 {
		return n.List[0].Pos()
	}
	return core.NoPos
}

// End returns the position of first character immediately after the node.
func (n *IdentList) End() core.Pos {
	if n.RParen.IsValid() {
		return n.RParen + 1
	}
	if l := len(n.List); l > 0 {
		return n.List[l-1].End()
	}
	return core.NoPos
}

// NumFields returns the number of fields.
func (n *IdentList) NumFields() int {
	if n == nil {
		return 0
	}
	return len(n.List)
}

func (n *IdentList) String() string {
	list := make([]string, 0, len(n.List))
	for i, e := range n.List {
		if n.VarArgs && i == len(n.List)-1 {
			list = append(list, "..."+e.String())
		} else {
			list = append(list, e.String())
		}
	}
	return "(" + strings.Join(list, ", ") + ")"
}
