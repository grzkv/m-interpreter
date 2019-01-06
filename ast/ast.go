package ast

import (
	"github.com/grzkv/m-interpreter/token"
	"strings"
)

// Node represents a node in AST. Can be either an expression or a statement
type Node interface {
	TokenLiteral() string
	String() string
}

// ExprNode is an expression node
type ExprNode interface {
	Node
	expr()
}

// StNode is a statement node
type StNode interface {
	Node
	statement()
}

// Program represents the whole interpreted program. Root of AST
type Program struct {
	StNodes []Node
}

// TokenLiteral makes Program a Node
func (p *Program) TokenLiteral() string {
	if len(p.StNodes) == 0 {
		return ""
	}
	return p.StNodes[0].TokenLiteral()
}

func (p *Program) String() string {
	var b strings.Builder

	for _, s := range p.StNodes {
		b.WriteString(s.String() + "\n")
	}

	return b.String()
}

// LetSt is the *let* statement
type LetSt struct {
	Token token.Token // always LET
	Ident *IdentifierEx
	Expr  ExprNode
}

// TokenLiteral makes LetSt a Node
func (s *LetSt) TokenLiteral() string {
	return s.Token.Literal
}

// statement makes LetSt a statement
func (s *LetSt) statement() {}

func (s *LetSt) String() string {
	var b strings.Builder

	b.WriteString("let ")
	b.WriteString(s.Ident.String() + " = ")
	b.WriteString(s.Expr.String() + ";")

	return b.String()
}

// ReturnSt represents the *return* statement
type ReturnSt struct {
	RootToken token.Token
	Expr      ExprNode
}

// TokenLiteral makes ReturnSt a Node
func (s *ReturnSt) TokenLiteral() string {
	return s.RootToken.Literal
}

func (s *ReturnSt) statement() {}

func (s *ReturnSt) String() string {
	return s.RootToken.Literal + " " + s.Expr.String()
}

// ExpressionSt represents statemnets like > 1 + 1;
type ExpressionSt struct {
	RootToken token.Token
	Expr      ExprNode
}

// TokenLiteral makes expression statement a node
func (s *ExpressionSt) TokenLiteral() string {
	return s.RootToken.Literal
}

func (s *ExpressionSt) statement() {}

func (s *ExpressionSt) String() string {
	if s != nil {
		return s.Expr.String()
	}
	return ""
}

// IdentifierEx is an identifier expression
type IdentifierEx struct {
	Token token.Token // always IDENT
	Value string
}

// TokenLiteral makes IdentifierEx a Node
func (e *IdentifierEx) TokenLiteral() string {
	return e.Token.Literal
}

// ex makes IdentifierEx an expression
func (e *IdentifierEx) expr() {}

func (e *IdentifierEx) String() string {
	return e.Value
}

// IntegerLiteralEx is an integer value
type IntegerLiteralEx struct {
	Token token.Token
	Value int64
}

// TokenLiteral makes integer literal a Node
func (expr *IntegerLiteralEx) TokenLiteral() string {
	return expr.Token.Literal
}

// String needed to make a Node
func (expr *IntegerLiteralEx) String() string {
	return expr.Token.Literal
}

func (expr *IntegerLiteralEx) expr() {}

// PrefixExpr represents e.g. !true, -(a+b), -5, etc.
type PrefixExpr struct {
	Token token.Token
	Op    string
	Right ExprNode
}

// TokenLiteral makes PrefixExpr a Node
func (expr *PrefixExpr) TokenLiteral() string {
	return expr.Token.Literal
}

// String makes PrefixExpr a Node
func (expr *PrefixExpr) String() string {
	return "(" + expr.Op + expr.Right.String() + ")"
}

// This makes PrefixExpr and expression
func (expr *PrefixExpr) expr() {}

// InfixExpr represents *1+1*, *(3-2) + (a * b)*, etc.
type InfixExpr struct {
	OpToken token.Token
	Left ExprNode
	Op string
	Right ExprNode
}

// TokenLiteral makes InfixExpr a Node
func (expr *InfixExpr) TokenLiteral() string {
	return expr.OpToken.Literal
}

func (expr *InfixExpr) String() string {
	var b strings.Builder

	b.WriteString("(")
	b.WriteString(expr.Left.String())
	b.WriteString(" " + expr.Op + " ")
	b.WriteString(expr.Right.String())
	b.WriteString(")")

	return b.String()
}

func (expr *InfixExpr) expr() {}