package parser

import (
	"strings"

	"github.com/jokruger/gs/token"
	gst "github.com/jokruger/gs/types"
)

// Expr represents an expression node in the AST.
type Expr interface {
	Node
	exprNode()
}

// ArrayLit represents an array literal.
type ArrayLit struct {
	Elements []Expr
	LBrack   gst.Pos
	RBrack   gst.Pos
}

func (e *ArrayLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ArrayLit) Pos() gst.Pos {
	return e.LBrack
}

// End returns the position of first character immediately after the node.
func (e *ArrayLit) End() gst.Pos {
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
	From gst.Pos
	To   gst.Pos
}

func (e *BadExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BadExpr) Pos() gst.Pos {
	return e.From
}

// End returns the position of first character immediately after the node.
func (e *BadExpr) End() gst.Pos {
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
	TokenPos gst.Pos
}

func (e *BinaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BinaryExpr) Pos() gst.Pos {
	return e.LHS.Pos()
}

// End returns the position of first character immediately after the node.
func (e *BinaryExpr) End() gst.Pos {
	return e.RHS.End()
}

func (e *BinaryExpr) String() string {
	return "(" + e.LHS.String() + " " + e.Token.String() +
		" " + e.RHS.String() + ")"
}

// BoolLit represents a boolean literal.
type BoolLit struct {
	Value    bool
	ValuePos gst.Pos
	Literal  string
}

func (e *BoolLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BoolLit) Pos() gst.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BoolLit) End() gst.Pos {
	return gst.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BoolLit) String() string {
	return e.Literal
}

// CallExpr represents a function call expression.
type CallExpr struct {
	Func     Expr
	LParen   gst.Pos
	Args     []Expr
	Ellipsis gst.Pos
	RParen   gst.Pos
}

func (e *CallExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CallExpr) Pos() gst.Pos {
	return e.Func.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CallExpr) End() gst.Pos {
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
	ValuePos gst.Pos
	Literal  string
}

func (e *CharLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CharLit) Pos() gst.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *CharLit) End() gst.Pos {
	return gst.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *CharLit) String() string {
	return e.Literal
}

// CondExpr represents a ternary conditional expression.
type CondExpr struct {
	Cond        Expr
	True        Expr
	False       Expr
	QuestionPos gst.Pos
	ColonPos    gst.Pos
}

func (e *CondExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CondExpr) Pos() gst.Pos {
	return e.Cond.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CondExpr) End() gst.Pos {
	return e.False.End()
}

func (e *CondExpr) String() string {
	return "(" + e.Cond.String() + " ? " + e.True.String() +
		" : " + e.False.String() + ")"
}

// ErrorExpr represents an error expression
type ErrorExpr struct {
	Expr     Expr
	ErrorPos gst.Pos
	LParen   gst.Pos
	RParen   gst.Pos
}

func (e *ErrorExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ErrorExpr) Pos() gst.Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ErrorExpr) End() gst.Pos {
	return e.RParen
}

func (e *ErrorExpr) String() string {
	return "error(" + e.Expr.String() + ")"
}

// FloatLit represents a floating point literal.
type FloatLit struct {
	Value    float64
	ValuePos gst.Pos
	Literal  string
}

func (e *FloatLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FloatLit) Pos() gst.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *FloatLit) End() gst.Pos {
	return gst.Pos(int(e.ValuePos) + len(e.Literal))
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
func (e *FuncLit) Pos() gst.Pos {
	return e.Type.Pos()
}

// End returns the position of first character immediately after the node.
func (e *FuncLit) End() gst.Pos {
	return e.Body.End()
}

func (e *FuncLit) String() string {
	return "func" + e.Type.Params.String() + " " + e.Body.String()
}

// FuncType represents a function type definition.
type FuncType struct {
	FuncPos gst.Pos
	Params  *IdentList
}

func (e *FuncType) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FuncType) Pos() gst.Pos {
	return e.FuncPos
}

// End returns the position of first character immediately after the node.
func (e *FuncType) End() gst.Pos {
	return e.Params.End()
}

func (e *FuncType) String() string {
	return "func" + e.Params.String()
}

// Ident represents an identifier.
type Ident struct {
	Name    string
	NamePos gst.Pos
}

func (e *Ident) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *Ident) Pos() gst.Pos {
	return e.NamePos
}

// End returns the position of first character immediately after the node.
func (e *Ident) End() gst.Pos {
	return gst.Pos(int(e.NamePos) + len(e.Name))
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
	ErrorPos gst.Pos
	LParen   gst.Pos
	RParen   gst.Pos
}

func (e *ImmutableExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImmutableExpr) Pos() gst.Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ImmutableExpr) End() gst.Pos {
	return e.RParen
}

func (e *ImmutableExpr) String() string {
	return "immutable(" + e.Expr.String() + ")"
}

// ImportExpr represents an import expression
type ImportExpr struct {
	ModuleName string
	Token      token.Token
	TokenPos   gst.Pos
}

func (e *ImportExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImportExpr) Pos() gst.Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *ImportExpr) End() gst.Pos {
	// import("moduleName")
	return gst.Pos(int(e.TokenPos) + 10 + len(e.ModuleName))
}

func (e *ImportExpr) String() string {
	return `import("` + e.ModuleName + `")`
}

// IndexExpr represents an index expression.
type IndexExpr struct {
	Expr   Expr
	LBrack gst.Pos
	Index  Expr
	RBrack gst.Pos
}

func (e *IndexExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IndexExpr) Pos() gst.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *IndexExpr) End() gst.Pos {
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
	ValuePos gst.Pos
	Literal  string
}

func (e *IntLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IntLit) Pos() gst.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *IntLit) End() gst.Pos {
	return gst.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *IntLit) String() string {
	return e.Literal
}

// MapElementLit represents a map element.
type MapElementLit struct {
	Key      string
	KeyPos   gst.Pos
	ColonPos gst.Pos
	Value    Expr
}

func (e *MapElementLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapElementLit) Pos() gst.Pos {
	return e.KeyPos
}

// End returns the position of first character immediately after the node.
func (e *MapElementLit) End() gst.Pos {
	return e.Value.End()
}

func (e *MapElementLit) String() string {
	return e.Key + ": " + e.Value.String()
}

// MapLit represents a map literal.
type MapLit struct {
	LBrace   gst.Pos
	Elements []*MapElementLit
	RBrace   gst.Pos
}

func (e *MapLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapLit) Pos() gst.Pos {
	return e.LBrace
}

// End returns the position of first character immediately after the node.
func (e *MapLit) End() gst.Pos {
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
	LParen gst.Pos
	RParen gst.Pos
}

func (e *ParenExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ParenExpr) Pos() gst.Pos {
	return e.LParen
}

// End returns the position of first character immediately after the node.
func (e *ParenExpr) End() gst.Pos {
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
func (e *SelectorExpr) Pos() gst.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SelectorExpr) End() gst.Pos {
	return e.Sel.End()
}

func (e *SelectorExpr) String() string {
	return e.Expr.String() + "." + e.Sel.String()
}

// SliceExpr represents a slice expression.
type SliceExpr struct {
	Expr   Expr
	LBrack gst.Pos
	Low    Expr
	High   Expr
	RBrack gst.Pos
}

func (e *SliceExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SliceExpr) Pos() gst.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SliceExpr) End() gst.Pos {
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
	ValuePos gst.Pos
	Literal  string
}

func (e *StringLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *StringLit) Pos() gst.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *StringLit) End() gst.Pos {
	return gst.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *StringLit) String() string {
	return e.Literal
}

// UnaryExpr represents an unary operator expression.
type UnaryExpr struct {
	Expr     Expr
	Token    token.Token
	TokenPos gst.Pos
}

func (e *UnaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *UnaryExpr) Pos() gst.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *UnaryExpr) End() gst.Pos {
	return e.Expr.End()
}

func (e *UnaryExpr) String() string {
	return "(" + e.Token.String() + e.Expr.String() + ")"
}

// UndefinedLit represents an undefined literal.
type UndefinedLit struct {
	TokenPos gst.Pos
}

func (e *UndefinedLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *UndefinedLit) Pos() gst.Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *UndefinedLit) End() gst.Pos {
	return e.TokenPos + 9 // len(undefined) == 9
}

func (e *UndefinedLit) String() string {
	return "undefined"
}
