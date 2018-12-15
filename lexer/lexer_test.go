package lexer

import (
	"github.com/grzkv/m-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(){}+=,;`

	tests := []struct {
		expTyp     token.Typ
		expLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

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
