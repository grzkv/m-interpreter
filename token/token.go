package token

type Typ string

type Token struct {
	Typ     Typ
	Literal string
}

// token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// ops
	PLUS   = "+"
	ASSIGN = "="

	//delims
	COMMA     = ","
	SEMICOLON = ";"

	// parens
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
