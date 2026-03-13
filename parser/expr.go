package parser

import (
	"strings"

	"github.com/jokruger/gs/token"
	"github.com/jokruger/gs/types"
)

// Expr represents an expression node in the AST.
type Expr interface {
	Node
	exprNode()
}

// ArrayLit represents an array literal.
type ArrayLit struct {
	Elements []Expr
	LBrack   types.Pos
	RBrack   types.Pos
}

func (e *ArrayLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ArrayLit) Pos() types.Pos {
	return e.LBrack
}

// End returns the position of first character immediately after the node.
func (e *ArrayLit) End() types.Pos {
	return e.RBrack + 1
}

func (e *ArrayLit) String() string {
	var elements []string
	for _, m := range e.Elements {
		elements = append(elements, m.String())
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// BadExpr represents a bad expression.
type BadExpr struct {
	From types.Pos
	To   types.Pos
}

func (e *BadExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BadExpr) Pos() types.Pos {
	return e.From
}

// End returns the position of first character immediately after the node.
func (e *BadExpr) End() types.Pos {
	return e.To
}

func (e *BadExpr) String() string {
	return "<bad expression>"
}

// BinaryExpr represents a binary operator expression.
type BinaryExpr struct {
	LHS      Expr
	RHS      Expr
	Token    token.Token
	TokenPos types.Pos
}

func (e *BinaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BinaryExpr) Pos() types.Pos {
	return e.LHS.Pos()
}

// End returns the position of first character immediately after the node.
func (e *BinaryExpr) End() types.Pos {
	return e.RHS.End()
}

func (e *BinaryExpr) String() string {
	return "(" + e.LHS.String() + " " + e.Token.String() +
		" " + e.RHS.String() + ")"
}

// BoolLit represents a boolean literal.
type BoolLit struct {
	Value    bool
	ValuePos types.Pos
	Literal  string
}

func (e *BoolLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BoolLit) Pos() types.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BoolLit) End() types.Pos {
	return types.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BoolLit) String() string {
	return e.Literal
}

// CallExpr represents a function call expression.
type CallExpr struct {
	Func     Expr
	LParen   types.Pos
	Args     []Expr
	Ellipsis types.Pos
	RParen   types.Pos
}

func (e *CallExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CallExpr) Pos() types.Pos {
	return e.Func.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CallExpr) End() types.Pos {
	return e.RParen + 1
}

func (e *CallExpr) String() string {
	var args []string
	for _, e := range e.Args {
		args = append(args, e.String())
	}
	if len(args) > 0 && e.Ellipsis.IsValid() {
		args[len(args)-1] = args[len(args)-1] + "..."
	}
	return e.Func.String() + "(" + strings.Join(args, ", ") + ")"
}

// CharLit represents a character literal.
type CharLit struct {
	Value    rune
	ValuePos types.Pos
	Literal  string
}

func (e *CharLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CharLit) Pos() types.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *CharLit) End() types.Pos {
	return types.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *CharLit) String() string {
	return e.Literal
}

// CondExpr represents a ternary conditional expression.
type CondExpr struct {
	Cond        Expr
	True        Expr
	False       Expr
	QuestionPos types.Pos
	ColonPos    types.Pos
}

func (e *CondExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CondExpr) Pos() types.Pos {
	return e.Cond.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CondExpr) End() types.Pos {
	return e.False.End()
}

func (e *CondExpr) String() string {
	return "(" + e.Cond.String() + " ? " + e.True.String() +
		" : " + e.False.String() + ")"
}

// ErrorExpr represents an error expression
type ErrorExpr struct {
	Expr     Expr
	ErrorPos types.Pos
	LParen   types.Pos
	RParen   types.Pos
}

func (e *ErrorExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ErrorExpr) Pos() types.Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ErrorExpr) End() types.Pos {
	return e.RParen
}

func (e *ErrorExpr) String() string {
	return "error(" + e.Expr.String() + ")"
}

// FloatLit represents a floating point literal.
type FloatLit struct {
	Value    float64
	ValuePos types.Pos
	Literal  string
}

func (e *FloatLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FloatLit) Pos() types.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *FloatLit) End() types.Pos {
	return types.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *FloatLit) String() string {
	return e.Literal
}

// FuncLit represents a function literal.
type FuncLit struct {
	Type *FuncType
	Body *BlockStmt
}

func (e *FuncLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FuncLit) Pos() types.Pos {
	return e.Type.Pos()
}

// End returns the position of first character immediately after the node.
func (e *FuncLit) End() types.Pos {
	return e.Body.End()
}

func (e *FuncLit) String() string {
	return "func" + e.Type.Params.String() + " " + e.Body.String()
}

// FuncType represents a function type definition.
type FuncType struct {
	FuncPos types.Pos
	Params  *IdentList
}

func (e *FuncType) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FuncType) Pos() types.Pos {
	return e.FuncPos
}

// End returns the position of first character immediately after the node.
func (e *FuncType) End() types.Pos {
	return e.Params.End()
}

func (e *FuncType) String() string {
	return "func" + e.Params.String()
}

// Ident represents an identifier.
type Ident struct {
	Name    string
	NamePos types.Pos
}

func (e *Ident) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *Ident) Pos() types.Pos {
	return e.NamePos
}

// End returns the position of first character immediately after the node.
func (e *Ident) End() types.Pos {
	return types.Pos(int(e.NamePos) + len(e.Name))
}

func (e *Ident) String() string {
	if e != nil {
		return e.Name
	}
	return nullRep
}

// ImmutableExpr represents an immutable expression
type ImmutableExpr struct {
	Expr     Expr
	ErrorPos types.Pos
	LParen   types.Pos
	RParen   types.Pos
}

func (e *ImmutableExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImmutableExpr) Pos() types.Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ImmutableExpr) End() types.Pos {
	return e.RParen
}

func (e *ImmutableExpr) String() string {
	return "immutable(" + e.Expr.String() + ")"
}

// ImportExpr represents an import expression
type ImportExpr struct {
	ModuleName string
	Token      token.Token
	TokenPos   types.Pos
}

func (e *ImportExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImportExpr) Pos() types.Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *ImportExpr) End() types.Pos {
	// import("moduleName")
	return types.Pos(int(e.TokenPos) + 10 + len(e.ModuleName))
}

func (e *ImportExpr) String() string {
	return `import("` + e.ModuleName + `")`
}

// IndexExpr represents an index expression.
type IndexExpr struct {
	Expr   Expr
	LBrack types.Pos
	Index  Expr
	RBrack types.Pos
}

func (e *IndexExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IndexExpr) Pos() types.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *IndexExpr) End() types.Pos {
	return e.RBrack + 1
}

func (e *IndexExpr) String() string {
	var index string
	if e.Index != nil {
		index = e.Index.String()
	}
	return e.Expr.String() + "[" + index + "]"
}

// IntLit represents an integer literal.
type IntLit struct {
	Value    int64
	ValuePos types.Pos
	Literal  string
}

func (e *IntLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IntLit) Pos() types.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *IntLit) End() types.Pos {
	return types.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *IntLit) String() string {
	return e.Literal
}

// MapElementLit represents a map element.
type MapElementLit struct {
	Key      string
	KeyPos   types.Pos
	ColonPos types.Pos
	Value    Expr
}

func (e *MapElementLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapElementLit) Pos() types.Pos {
	return e.KeyPos
}

// End returns the position of first character immediately after the node.
func (e *MapElementLit) End() types.Pos {
	return e.Value.End()
}

func (e *MapElementLit) String() string {
	return e.Key + ": " + e.Value.String()
}

// MapLit represents a map literal.
type MapLit struct {
	LBrace   types.Pos
	Elements []*MapElementLit
	RBrace   types.Pos
}

func (e *MapLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapLit) Pos() types.Pos {
	return e.LBrace
}

// End returns the position of first character immediately after the node.
func (e *MapLit) End() types.Pos {
	return e.RBrace + 1
}

func (e *MapLit) String() string {
	var elements []string
	for _, m := range e.Elements {
		elements = append(elements, m.String())
	}
	return "{" + strings.Join(elements, ", ") + "}"
}

// ParenExpr represents a parenthesis wrapped expression.
type ParenExpr struct {
	Expr   Expr
	LParen types.Pos
	RParen types.Pos
}

func (e *ParenExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ParenExpr) Pos() types.Pos {
	return e.LParen
}

// End returns the position of first character immediately after the node.
func (e *ParenExpr) End() types.Pos {
	return e.RParen + 1
}

func (e *ParenExpr) String() string {
	return "(" + e.Expr.String() + ")"
}

// SelectorExpr represents a selector expression.
type SelectorExpr struct {
	Expr Expr
	Sel  Expr
}

func (e *SelectorExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SelectorExpr) Pos() types.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SelectorExpr) End() types.Pos {
	return e.Sel.End()
}

func (e *SelectorExpr) String() string {
	return e.Expr.String() + "." + e.Sel.String()
}

// SliceExpr represents a slice expression.
type SliceExpr struct {
	Expr   Expr
	LBrack types.Pos
	Low    Expr
	High   Expr
	RBrack types.Pos
}

func (e *SliceExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SliceExpr) Pos() types.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SliceExpr) End() types.Pos {
	return e.RBrack + 1
}

func (e *SliceExpr) String() string {
	var low, high string
	if e.Low != nil {
		low = e.Low.String()
	}
	if e.High != nil {
		high = e.High.String()
	}
	return e.Expr.String() + "[" + low + ":" + high + "]"
}

// StringLit represents a string literal.
type StringLit struct {
	Value    string
	ValuePos types.Pos
	Literal  string
}

func (e *StringLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *StringLit) Pos() types.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *StringLit) End() types.Pos {
	return types.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *StringLit) String() string {
	return e.Literal
}

// UnaryExpr represents an unary operator expression.
type UnaryExpr struct {
	Expr     Expr
	Token    token.Token
	TokenPos types.Pos
}

func (e *UnaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *UnaryExpr) Pos() types.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *UnaryExpr) End() types.Pos {
	return e.Expr.End()
}

func (e *UnaryExpr) String() string {
	return "(" + e.Token.String() + e.Expr.String() + ")"
}

// UndefinedLit represents an undefined literal.
type UndefinedLit struct {
	TokenPos types.Pos
}

func (e *UndefinedLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *UndefinedLit) Pos() types.Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *UndefinedLit) End() types.Pos {
	return e.TokenPos + 9 // len(undefined) == 9
}

func (e *UndefinedLit) String() string {
	return "undefined"
}
