package ast

import (
	"github.com/grzkv/m-interpreter/token"
	"testing"
)

func TestASTLet(t *testing.T) {
	prg := &Program{
		StNodes: []Node{
			&LetSt{
				Token: token.Token{Typ: token.LET, Literal: "let"},
				Ident: &IdentifierEx{
					Token: token.Token{Typ: token.IDENT, Literal: "aleph"},
					Value: "aleph",
				},
				Expr: &IdentifierEx{
					Token: token.Token{Typ: token.IDENT, Literal: "alpha"},
					Value: "alpha",
				},
			},
		},
	}

	const expStr = "let aleph = alpha;\n"

	if prg.String() != expStr {
		t.Fatalf("Expected string %q but got %q", expStr, prg.String())
	}
}
