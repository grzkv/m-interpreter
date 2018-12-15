package lexer

import "github.com/grzkv/m-interpreter/token"

type Lexer struct {
	input   string
	pos     int
	rPos    int
	current byte
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.Read()

	return &l
}

func (l *Lexer) Read() {
	if l.rPos >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.rPos]
	}

	l.pos = l.rPos
	l.rPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	switch l.current {
	case '+':
		t = token.Token{token.PLUS, "+"}
	case '=':
		t = token.Token{token.ASSIGN, "="}
	case ',':
		t = token.Token{token.COMMA, ","}
	case ';':
		t = token.Token{token.SEMICOLON, ";"}
	case '(':
		t = token.Token{token.LPAREN, "("}
	case ')':
		t = token.Token{token.RPAREN, ")"}
	case '{':
		t = token.Token{token.LBRACE, "{"}
	case '}':
		t = token.Token{token.RBRACE, "}"}
	case 0:
		t = token.Token{token.EOF, ""}
	default:
		t = token.Token{token.ILLEGAL, ""}
	}

	l.Read()
	return t
}
