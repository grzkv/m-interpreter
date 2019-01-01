package lexer

import (
	"testing"

	. "github.com/grzkv/m-interpreter/token"
)

type ExpToken struct {
	expTyp     Typ
	expLiteral string
}

func TestNextTokenSimple(t *testing.T) {
	input := `(){}+=,;`

	tests := []ExpToken{
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{PLUS, "+"},
		{ASSIGN, "="},
		{COMMA, ","},
		{SEMICOLON, ";"},
	}

	runLexerTest(t, input, tests)
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
	  x + y;
	};
	
	let result = add(five, ten);
	`
	tests := []ExpToken{
		{LET, "let"},
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},

		{LET, "let"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},

		{LET, "let"},
		{IDENT, "add"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},

		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{SEMICOLON, ";"},

		{RBRACE, "}"},
		{SEMICOLON, ";"},

		{LET, "let"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "ten"},
		{SEMICOLON, ";"},
	}

	runLexerTest(t, input, tests)
}

func runLexerTest(t *testing.T, input string, tests []ExpToken) {
	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Typ != tt.expTyp {
			t.Fatalf("Test %d failed. Expected token type %q - got %q", i,
				tt.expTyp, tk.Typ)
		}

		if tk.Literal != tt.expLiteral {
			t.Fatalf("Test %d failed. Expected token literal %q - got %q", i,
				tt.expTyp, tk.Typ)
		}

	}
}
