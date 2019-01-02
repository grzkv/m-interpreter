package ast

import (
	"github.com/grzkv/m-interpreter/token"
)

// Node represents a node in AST. Can be either an expression or a statement
type Node interface {
	TokenLiteral() string
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
	StNodes []StNode
}

// TokenLiteral makes Program a Node
func (p *Program) TokenLiteral() string {
	if len(p.StNodes) == 0 {
		return ""
	}
	return p.StNodes[0].TokenLiteral()
}

// LetSt is the *let* statement
type LetSt struct {
	Token token.Token // always LET
	Ident *IdentifierEx
	Expr ExprNode
}

// TokenLiteral makes LetSt a Node
func (s *LetSt) TokenLiteral() string {
	return s.Token.Literal
}

// statement makes LetSt a statement
func (s *LetSt) statement() {}

// IdentifierEx is an identifier expression
type IdentifierEx struct {
	Token token.Token // always IDENT
}

// TokenLiteral makes IdentifierEx a Node
func (e *IdentifierEx) TokenLiteral() string {
	return e.Token.Literal
}

// ex makes IdentifierEx an expression
func (e *IdentifierEx) expr() {}