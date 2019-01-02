package token

// Typ is token type
type Typ string

// Token of the Monkey PL
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
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FALSE    = "FALSE"
	TRUE     = "TRUE"

	// ops
	NOT     = "!"
	MINUS   = "-"
	DIVIDE  = "/"
	MULT    = "*"
	LESS    = "<"
	GREATER = ">"
	EQ      = "=="
	NEQ     = "!="

	IDENT = "IDENT"

	// literals
	INT = "INT"
)
