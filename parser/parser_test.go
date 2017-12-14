package parser

import (
	"fmt"
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

func TestIntegerLiteral(t *testing.T) {
	input := "5;"
	program := testParseProgram(t, input, []string{})
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression not an ast.IntegerLiteral Got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Fatalf("Wrong literal.Value Expected=5 Got=%d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("Wrong literal.TokenLiteral() Expected=5 Got=%s", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}
	for _, test := range prefixTests {
		program := testParseProgram(t, test.input, []string{})
		testNumberOfStatemets(t, program, 1)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression not an *ast.PrefixExpression, Got=%T", stmt.Expression)
		}
		if exp.Operator != test.operator {
			t.Fatalf("Wrong exp.Operator Expected=%s Got=%s", test.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, test.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	intLiteral, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("exp not and ast.IntegerLiteral Got=%T", exp)
		return false
	}
	if intLiteral.Value != value {
		t.Errorf("Wrong intLiteral.Value Exptected=%d Got=%d", value, intLiteral.Value)
		return false
	}
	if intLiteral.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("Wrong intLiteral.Value Exptected=%d Got=%s",
			intLiteral.Value, intLiteral.TokenLiteral())
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

func testNumberOfStatemets(t *testing.T, program *ast.Program, expectedNoStmts int) {
	if len(program.Statements) != expectedNoStmts {
		t.Fatalf("Wrong number of program statemets. Expected=%d got=%d",
			expectedNoStmts, len(program.Statements))
	}
}
