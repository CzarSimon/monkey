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
	let x = 5;
  let y = 10;
  let foobar = 838383;
  `
	noStatements := 3
	expectedErrors := []string{}
	program := testParseProgram(t, input, expectedErrors)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	testNumberOfStatemets(t, program, noStatements)
	testStatements := []testStatement{{"x"}, {"y"}, {"foobar"}}
	for i, testStmt := range testStatements {
		stmt := program.Statements[i]
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

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;`
	noStatements := 3
	program := testParseProgram(t, input, []string{})

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	testNumberOfStatemets(t, program, noStatements)
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not an *ast.ReturnStatement. Got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return' Got=‰q", returnStmt.TokenLiteral())
		}
	}
}

func testParseProgram(t *testing.T, input string, expectedErrors []string) *ast.Program {
	lex := lexer.New(input)
	parser := New(lex)
	program := parser.ParseProgram()
	checkParserErrors(t, parser, expectedErrors)
	return program
}

func TestIndentifierExpression(t *testing.T) {
	input := "foobar;"
	expectedErrors := []string{}
	program := testParseProgram(t, input, expectedErrors)
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expresion is not an ast.Identifier. Got=%T",
			stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. Got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. Got=%s", "foobar", ident.TokenLiteral())
	}
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

func testNumberOfStatemets(t *testing.T, program *ast.Program, expectedNoStmts int) {
	if len(program.Statements) != expectedNoStmts {
		t.Fatalf("Wrong number of program statemets. Expected=%d got=%d",
			expectedNoStmts, len(program.Statements))
	}
}
