package parser

import (
	"github.com/grzkv/m-interpreter/ast"
	"github.com/grzkv/m-interpreter/token"
	"github.com/grzkv/m-interpreter/lexer"
)

// Parser parses the code tokenized by lexer
type Parser struct {
	l *lexer.Lexer

	current token.Token
	peek token.Token
}

// New makes a new parser. Usable out-of-the-box
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.l.NextToken()
}

// Parse the loaded code
func (p *Parser) Parse() *ast.Program {
	return nil
}