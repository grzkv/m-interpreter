package lexer

import "github.com/grzkv/m-interpreter/token"
import "log"

// Lexer breaks code into tokens
type Lexer struct {
	input   string
	pos     int
	rPos    int
	current byte
}

// New makes a lexer
func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readCh()

	return &l
}

// Read a char. Puts it in the lexer buffer
func (l *Lexer) readCh() {
	if l.rPos >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.rPos]
	}

	l.pos = l.rPos
	l.rPos++
}

// NextToken gets next token from the code
func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()

	var t token.Token
	switch l.current {
	case '+':
		t = token.Token{Typ: token.PLUS, Literal: "+"}
	case '=':
		t = token.Token{Typ: token.ASSIGN, Literal: "="}
	case ',':
		t = token.Token{Typ: token.COMMA, Literal: ","}
	case ';':
		t = token.Token{Typ: token.SEMICOLON, Literal: ";"}
	case '(':
		t = token.Token{Typ: token.LPAREN, Literal: "("}
	case ')':
		t = token.Token{Typ: token.RPAREN, Literal: ")"}
	case '{':
		t = token.Token{Typ: token.LBRACE, Literal: "{"}
	case '}':
		t = token.Token{Typ: token.RBRACE, Literal: "}"}
	case 0:
		t = token.Token{Typ: token.EOF, Literal: ""}
	default:
		if isLetter(l.current) {
			word := l.readWord()
			log.Println(word)
			
			kw, prs := keywords[word]

			if prs {
				t = token.Token{Typ: kw, Literal: word}
			}else{
				t = token.Token{Typ: token.IDENT, Literal: word}
			}
		}
	}

	l.readCh()
	return t
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.current) {
		l.readCh()
	}
}

var keywords = map[string]token.Typ {
	"fn": token.FUNCTION,
	"let": token.LET,
}

func (l *Lexer) readWord() string {
	firstLetterPos := l.pos
	for isLetter(l.current) {
		l.readCh();
	}

	return l.input[firstLetterPos:l.pos]
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 0 && c <= 9) ||
		c == '_';
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}