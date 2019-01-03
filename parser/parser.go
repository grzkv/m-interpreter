package parser

import (
	"fmt"
	"log"
	"github.com/grzkv/m-interpreter/ast"
	"github.com/grzkv/m-interpreter/lexer"
	"github.com/grzkv/m-interpreter/token"
)

// Parser parses the code tokenized by lexer
type Parser struct {
	l *lexer.Lexer

	current token.Token
	peek    token.Token
	errors []string
}

// New makes a new parser. Usable out-of-the-box
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	log.Println("Making new parser")

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
		}else{
			log.Println("error: got nil statment during parsing")
		}
	}

	p.printErrorsIfAny()

	return &prg
}

func (p *Parser) printErrorsIfAny() {
	if len(p.errors) == 0 {
		return
	}

	fmt.Printf("PARSER GOT %d ERRORS:\n", len(p.errors));
	for _, e := range p.errors {
		fmt.Println(e)
	}
}

func (p *Parser) parseStatement() ast.StNode {
	log.Println("Parsing a stament")

	switch p.current.Typ {
	case token.LET:
		return p.parseLetSt()
	default:
		p.nextToken()
		return nil
	}
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

	if (p.current.Typ == token.SEMICOLON) {
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
