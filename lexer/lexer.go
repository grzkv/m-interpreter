package lexer

import "github.com/grzkv/m-interpreter/token"
// import "log"

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

// NextToken gets next token from the code
func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()

	var t token.Token
	switch l.current {
	case '+':
		t = token.Token{Typ: token.PLUS, Literal: "+"}
	case '=':
		if l.peek() == '=' {
			l.readCh()
			t = token.Token{Typ: token.EQ, Literal: "=="}
		} else {
			t = token.Token{Typ: token.ASSIGN, Literal: "="}
		}
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
	case '!':
		if l.peek() == '=' {
			l.readCh()
			t = token.Token{Typ: token.NEQ, Literal: "!="}
		} else {
			t = token.Token{Typ: token.NOT, Literal: "!"}
		}
	case '-':
		t = token.Token{Typ: token.MINUS, Literal: "-"}
	case '/':
		t = token.Token{Typ: token.DIVIDE, Literal: "/"}
	case '*':
		t = token.Token{Typ: token.MULT, Literal: "*"}
	case '<':
		t = token.Token{Typ: token.LESS, Literal: "<"}
	case '>':
		t = token.Token{Typ: token.GREATER, Literal: ">"}
	case 0:
		t = token.Token{Typ: token.EOF, Literal: ""}
	default:
		if isLetter(l.current) {
			word := l.readWord()
			// log.Printf("read word %s", word)

			kw, prs := keywords[word]

			if prs {
				// log.Println("classified word as a keyword")
				t = token.Token{Typ: kw, Literal: word}
			} else {
				t = token.Token{Typ: token.IDENT, Literal: word}
			}
		} else if isDigit(l.current) {
			number := l.readNumber()
			// log.Printf("read number %s", number)

			t = token.Token{Typ: token.INT, Literal: number}
		} else {
			t = token.Token{Typ: token.ILLEGAL, Literal: ""}
		}

		return t
	}

	l.readCh()
	return t
}

func (l *Lexer) readCh() {
	if l.rPos >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.rPos]
	}

	l.pos = l.rPos
	l.rPos++
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.current) {
		l.readCh()
	}
}

func (l Lexer) peek() byte {
	if l.rPos >= len(l.input) {
		return 0
	}
	return l.input[l.rPos]
}

func (l *Lexer) readWord() string {
	firstLetterPos := l.pos
	for isLetter(l.current) {
		l.readCh()
	}

	return l.input[firstLetterPos:l.pos]
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.current) {
		l.readCh()
	}

	return l.input[start:l.pos]
}

var keywords = map[string]token.Typ{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
	"false":  token.FALSE,
	"true":   token.TRUE,
}

// char utils

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
