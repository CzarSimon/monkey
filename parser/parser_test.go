package parser

import (
	"testing"

	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/lexer"
)

type testStatement struct {
	expectedIdentifier string
}

func TestLetStatements(t *testing.T) {
	input := `
  let x 5;
  let = 10;
  let 838383;
	let x = 5;
  let y = 10;
  let foobar = 838383;
  `
	noStatements := 3
	expectedErrors := []string{
		"peekToken: Expected type== Got=INT",
		"peekToken: Expected type=IDENT Got==",
		"peekToken: Expected type=IDENT Got=INT",
	}
	lex := lexer.New(input)
	parser := New(lex)
	program := parser.ParseProgram()
	checkParserErrors(t, parser, expectedErrors)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != noStatements {
		t.Fatalf("Wrong number of program statemets. Expected=%d got=%d",
			noStatements, len(program.Statements))
	}
	testStatements := []testStatement{{"x"}, {"y"}, {"foobar"}}
	for i, testStmt := range testStatements {
		stmt := program.Statements[i]
		if stmt == nil {
			continue
		}
		if !testLetStatement(t, stmt, testStmt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt == nil {
		t.Errorf("Stmt is nil")
		return false
	}
	if stmt.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral not let, got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not an *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not=%s. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name not=%s. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, parser *Parser, expectedErrors []string) {
	errors := parser.Errors()
	if len(errors) == len(expectedErrors) {
		for i, err := range errors {
			if err.Error() != expectedErrors[i] {
				t.Fatalf("%d. - Wrong error. Expected=[ %s ] but Got=[ %s ]",
					i, expectedErrors[i], err.Error())
			}
		}
		return
	}
	t.Errorf("parser has %d errors, expected: %d", len(errors), len(expectedErrors))
	for _, err := range errors {
		t.Errorf("parser error: %s", err.Error())
	}
	t.FailNow()
}
