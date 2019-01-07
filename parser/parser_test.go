package parser

import (
	"fmt"
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

func testLetStatement(t *testing.T, st ast.Node, expIdent string) {
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

func TestParsingIdentExpression(t *testing.T) {
	input := "beth;"
	const expNumSt = 1

	l := lexer.New(input)
	p := New(l)

	prg := p.Parse()
	if len(p.errors) != 0 {
		t.Fatalf("Parser got errors")
	}

	if prg == nil {
		t.Fatal("Got nil program")
	}

	if len(prg.StNodes) != expNumSt {
		t.Fatalf("Got %d nodes instead of %d", len(prg.StNodes), expNumSt)
	}

	exprSt, ok := (prg.StNodes[0]).(*ast.ExpressionSt)
	if !ok {
		t.Fatalf("Wanted expr st node, got %T", prg.StNodes[0])
	}

	identExpr, ok := exprSt.Expr.(*ast.IdentifierEx)
	if !ok {
		t.Fatalf("Wanted expr type identifier, got %T", prg.StNodes[0])
	}

	if identExpr.Token.Literal != "beth" {
		t.Fatalf("Wanted token literal *beth*, got %s", identExpr.Token.Literal)
	}

	if identExpr.Value != "beth" {
		t.Fatalf("wanted idnet expression value *beth*, got %s", identExpr.Value)
	}
}

func TestParsingIntExpression(t *testing.T) {
	input := "5;"
	expNumSt := 1

	l := lexer.New(input)
	p := New(l)

	prg := p.Parse()

	if len(p.errors) != 0 {
		t.Fatal("Parser has errors")
	}

	if prg == nil {
		t.Fatal("Parsed program is nil")
	}

	if len(prg.StNodes) != expNumSt {
		t.Fatalf("Want %d statements, got %d", expNumSt, len(prg.StNodes))
	}

	exprSt, ok := (prg.StNodes[0]).(*ast.ExpressionSt)

	if !ok {
		t.Fatalf("Expected expression statement, got %T", prg.StNodes[0])
	}

	intLiteralEx, ok := (exprSt.Expr).(*ast.IntegerLiteralEx)

	if !ok {
		t.Fatalf("Expected integer literal node, got %T", exprSt.Expr)
	}

	if intLiteralEx.TokenLiteral() != "5" {
		t.Fatalf("Expected token literal 5 from int literal, got %s", intLiteralEx.TokenLiteral())
	}

	if intLiteralEx.Value != 5 {
		t.Fatalf("Expected int literal value to be %d, got %d", 5, intLiteralEx.Value)
	}
}

func TestPrefixExpressionParsing(t *testing.T) {
	tests := []struct {
		in  string
		op  string
		val int64
	}{
		{in: "!1;", op: "!", val: 1},
		{in: "-11;", op: "-", val: 11},
	}

	for _, tst := range tests {
		l := lexer.New(tst.in)
		p := New(l)
		prg := p.Parse()

		if len(p.errors) != 0 {
			t.Fatal("Parser has errors")
		}

		if prg == nil {
			t.Fatal("Parser returned nil program")
		}

		if len(prg.StNodes) != 1 {
			t.Fatal("Parser got incorrect numbers of statements")
		}

		exprSt, ok := (prg.StNodes[0]).(*ast.ExpressionSt)
		if !ok {
			t.Fatalf("Want expr st, got %T", prg.StNodes[0])
		}

		prefExpr, ok := (exprSt.Expr).(*ast.PrefixExpr)
		if !ok {
			t.Fatalf("Want prefix expr, got %T", exprSt.Expr)
		}

		if prefExpr.Op != tst.op {
			t.Fatalf("Want op %s in prefix expr, got %s", tst.op, prefExpr.Op)
		}

		testIntLiteral(t, prefExpr.Right, tst.val)
	}
}

func testIntLiteral(t *testing.T, expr ast.ExprNode, expectedVal int64) {
	intLiteralExpr, ok := expr.(*ast.IntegerLiteralEx)
	if !ok {
		t.Fatalf("Wanted integer expression node, got %T", expr)
	}

	if intLiteralExpr.Value != expectedVal {
		t.Fatalf("Expected val %d of integer literal expr, got %d", expectedVal, intLiteralExpr.Value)
	}

	if intLiteralExpr.TokenLiteral() != fmt.Sprintf("%d", expectedVal) {
		t.Fatalf("Expected token literal %q, got %q", fmt.Sprintf("%d", expectedVal), intLiteralExpr.TokenLiteral())
	}
}

func TestParsingInfixExpr(t *testing.T) {
	tests := []struct {
		in   string
		lval int64
		op   string
		rval int64
	}{
		{"1 + 1", 1, "+", 1},
		{"1 - 1", 1, "-", 1},
		{"1 * 1", 1, "*", 1},
		{"1 / 1", 1, "/", 1},
		{"1 > 1", 1, ">", 1},
		{"1 < 1", 1, "<", 1},
		{"1 == 1", 1, "==", 1},
		{"1 != 1", 1, "!=", 1},
	}

	for _, tst := range tests {
		l := lexer.New(tst.in)
		p := New(l)

		prg := p.Parse()
		if len(p.errors) != 0 {
			t.Fatalf("Parser got errors")
		}

		if prg == nil {
			t.Fatalf("Program is nil")
		}

		if len(prg.StNodes) != 1 {
			t.Fatalf("Expected one statement, got %d", len(prg.StNodes))
		}

		exprSt, ok := (prg.StNodes[0]).(*ast.ExpressionSt)
		if !ok {
			t.Fatalf("Expression statement expected")
		}

		infixExpr, ok := (exprSt.Expr).(*ast.InfixExpr)
		if !ok {
			t.Fatalf("Expected infix expression, got %T", exprSt.Expr)
		}

		testIntLiteral(t, infixExpr.Left, tst.lval)
		testIntLiteral(t, infixExpr.Right, tst.rval)

		if infixExpr.Op != tst.op {
			t.Fatalf("Expected oprator %s, got %s", tst.op, infixExpr.Op)
		}
	}
}

func TestExpressionParsingWithOpPrecedence(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{
			"a + b",
			"(a + b)\n",
		},
		{
			"a + b * c",
			"(a + (b * c))\n",
		},
		{
			"-a * b",
			"((-a) * b)\n",
		},
		{
			"a + b + c / d",
			"((a + b) + (c / d))\n",
		},
		{
			"a == b * c",
			"(a == (b * c))\n",
		},
	}

	for _, tst := range tests {
		l := lexer.New(tst.in)
		p := New(l)
		prg := p.Parse()

		if prg == nil {
			t.Fatal("Got nil program")
		}

		if prg.String() != tst.expected {
			t.Fatalf("Got %q, expected %q", prg.String(), tst.expected)
		}
	}
}
