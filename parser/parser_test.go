package parser

import (
	"testing"

	"github.com/grzkv/m-interpreter/ast"
	"github.com/grzkv/m-interpreter/lexer"
	"github.com/grzkv/m-interpreter/token"
)

func TestLet(t *testing.T) {
	input := `
	let x = 1;
	let y = 2;
	let zzz = 838383;
	`
	expectedNSt := 3

	l := lexer.New(input)
	p := New(l)

	prg := p.Parse()

	if prg == nil {
		t.Fatal("Parser returned nil")
	}

	if len(prg.StNodes) != expectedNSt {
		t.Fatalf("Expected %d statements, got %d", expectedNSt, len(prg.StNodes))
	}

	expSts := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
		{"zzz"},
	}

	for i, expSt := range expSts {
		testLetStatement(t, prg.StNodes[i], expSt.expIdent)
	}
}

func testLetStatement(t *testing.T, st ast.StNode, expIdent string) {
	if st.TokenLiteral() != "let" {
		t.Fatalf("Expected token literal let, got %s", st.TokenLiteral())
	}

	letSt, ok := st.(*ast.LetSt)

	if !ok {
		t.Fatalf("Statement type is %T, expected *ast.LetSt", st)
	}

	if letSt.Token.Typ != token.LET {
		t.Fatalf("Expected token type LET, got %s", letSt.Token.Typ)
	}

	if letSt.Ident.Token.Literal != expIdent {
		t.Fatalf("Expected identifier name %s, got %s", expIdent, letSt.Ident.Token.Literal)
	}

	if letSt.Ident.Value != expIdent {
		t.Fatalf("Expected identifier value %s, got %s", expIdent, letSt.Ident.Value)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 1;
	return 2;
	return (a + b);
	return 88888;
	`
	numSt := 4

	l := lexer.New(input)
	p := New(l)

	prg := p.Parse()

	if prg == nil {
		t.Fatal("Got nil program")
	}

	if len(prg.StNodes) != numSt {
		t.Fatalf("Got %d statements, expected %d", len(prg.StNodes), numSt)
	}

	for i, st := range prg.StNodes {
		if st.TokenLiteral() != "return" {
			t.Fatalf("Error in statement %d. Expected token literal return, got %s", i, st.TokenLiteral())
		}
		_, ok := st.(*ast.ReturnSt) 
		if !ok {
			t.Fatalf("Got wrong node type, expected return statement")
		}
	}
}
