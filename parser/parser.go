package parser

import (
	"fmt"
	"github.com/grzkv/m-interpreter/ast"
	"github.com/grzkv/m-interpreter/lexer"
	"github.com/grzkv/m-interpreter/token"
	"log"
	"strconv"
)

// Parser parses the code tokenized by lexer
type Parser struct {
	l *lexer.Lexer

	current token.Token
	peek    token.Token
	errors  []string

	prefixParseFns map[token.Typ]prefixParseFn
	infixParseFns  map[token.Typ]infixParseFn
}

const (
	// operation precedence
	_ int = iota
	// LOWEST is the default
	LOWEST
	// EQ is ==
	EQ
	// LESSGR is for > and <
	LESSGR
	// SUM is for +
	SUM
	// PRODUCT is for *
	PRODUCT
	// PREFIX is for prefix oprators
	PREFIX
	// CALL is for function calls
	CALL
)

type (
	prefixParseFn func() ast.ExprNode
	infixParseFn  func(ast.ExprNode) ast.ExprNode
)

// New makes a new parser. Usable out-of-the-box
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	log.Println("Making new parser")

	p.prefixParseFns = make(map[token.Typ]prefixParseFn)
	p.prefixParseFns[token.IDENT] = p.parseIdent
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
	p.prefixParseFns[token.NOT] = p.parsePrefixExpr
	p.prefixParseFns[token.MINUS] = p.parsePrefixExpr

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	log.Println("Moving to next token")

	p.current = p.peek
	p.peek = p.l.NextToken()
}

// Parse the loaded code
func (p *Parser) Parse() *ast.Program {
	prg := ast.Program{}

	for p.current.Typ != token.EOF {
		st := p.parseStatement()

		if st != nil {
			prg.StNodes = append(prg.StNodes, st)
		} else {
			log.Println("error: got nil statement during parsing")
		}
	}

	p.printErrorsIfAny()

	return &prg
}

func (p *Parser) printErrorsIfAny() {
	if len(p.errors) == 0 {
		return
	}

	fmt.Printf("PARSER GOT %d ERRORS:\n", len(p.errors))
	for _, e := range p.errors {
		fmt.Println(e)
	}
}

func (p *Parser) parseStatement() ast.StNode {
	log.Println("Parsing a stament")

	switch p.current.Typ {
	case token.LET:
		return p.parseLetSt()
	case token.RETURN:
		return p.parseReturnSt()
	default:
		return p.parseExpressionSt()
	}
}

func (p *Parser) parseExpressionSt() *ast.ExpressionSt {
	st := ast.ExpressionSt{RootToken: p.current}

	st.Expr = p.parseExpr(LOWEST)

	if p.peek.Typ == token.SEMICOLON {
		p.nextToken()
	}

	p.nextToken()

	return &st
}

func (p *Parser) parseExpr(prio int) ast.ExprNode {
	prefixFn := p.prefixParseFns[p.current.Typ]

	if prefixFn == nil {
		return nil
	}

	return prefixFn()
}

func (p *Parser) parseIdent() ast.ExprNode {
	return &ast.IdentifierEx{Token: p.current, Value: p.current.Literal}
}

func (p *Parser) parseLetSt() *ast.LetSt {
	if p.current.Typ != token.LET {
		log.Println("error: wrong token type")
		p.errors = append(p.errors, "wrong token type")
		return nil
	}

	st := ast.LetSt{Token: p.current}

	p.nextToken()

	if p.current.Typ != token.IDENT {
		log.Println("error: wrong token type")
		p.errors = append(p.errors, "wrong token type")
		return nil
	}

	st.Ident = &ast.IdentifierEx{Token: p.current, Value: p.current.Literal}

	p.nextToken()
	if p.current.Typ != token.ASSIGN {
		log.Println("error: wrong token type")
		p.errors = append(p.errors, "wrong token type")
		return nil
	}

	// go until semicolon
	p.nextToken()

	if p.current.Typ == token.SEMICOLON {
		log.Println("error: empty expression in let statement")
		p.errors = append(p.errors, "wrong token type")
		return nil
	}

	for p.current.Typ != token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken() // pass the semicolon

	st.Expr = nil

	return &st
}

func (p *Parser) parseReturnSt() *ast.ReturnSt {
	if p.current.Typ != token.RETURN {
		p.errors = append(p.errors, fmt.Sprintf("Got wrong token type %s for return statement", p.current.Typ))
		log.Printf("error: got wrong token type for return statement")
	}

	returnSt := ast.ReturnSt{RootToken: p.current, Expr: nil}

	for p.current.Typ != token.SEMICOLON {
		p.nextToken()
	}

	p.nextToken() // pass semicolon

	return &returnSt
}

func (p *Parser) parseIntegerLiteral() ast.ExprNode {
	intLitExpr := ast.IntegerLiteralEx{Token: p.current}

	val, err := strconv.ParseInt(p.current.Literal, 0, 64)

	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Error while parsing integer literal: %v", err))
		log.Println("error parsing integer literal")

		return nil
	}

	intLitExpr.Value = val

	return &intLitExpr
}

func (p *Parser) parsePrefixExpr() ast.ExprNode {
	prefixExpr := ast.PrefixExpr{
		Token: p.current,
		Op:    p.current.Literal,
	}

	p.nextToken()

	prefixExpr.Right = p.parseExpr(PREFIX)

	return &prefixExpr
}
