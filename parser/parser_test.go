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

func TestLetStatementsWExpressions(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		exprectedValue     interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let isTrue = true;", "isTrue", true},
		{"let foobar = y;", "foobar", "y"},
	}
	expectedErrors := []string{}
	for _, test := range tests {
		program := testParseProgram(t, test.input, expectedErrors)
		testNumberOfStatemets(t, program, 1)
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, test.expectedIdentifier) {
			return
		}
		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, test.exprectedValue) {
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
		input    string
		operator string
		value    interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
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
		if !testLiteralExpression(t, exp.Right, test.value) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	prefixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"15 - 15", 15, "-", 15},
		{"10 * 10", 10, "*", 10},
		{"8 / 2", 8, "/", 2},
		{"0 == 1", 0, "==", 1},
		{"7 != 6", 7, "!=", 6},
		{"3 > 11", 3, ">", 11},
		{"12 < 9", 12, "<", 9},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	for _, test := range prefixTests {
		program := testParseProgram(t, test.input, []string{})
		testNumberOfStatemets(t, program, 1)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
		}
		infixExpr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt not an *ast.InfixExpression Got=%t", stmt.Expression)
		}
		if infixExpr.Operator != test.operator {
			t.Fatalf("Wrong exp.Operator Expected=%s Got=%s",
				test.operator, infixExpr.Operator)
		}
		if !testLiteralExpression(t, infixExpr.Left, test.leftValue) {
			return
		}
		if !testLiteralExpression(t, infixExpr.Right, test.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a*b*c", "((a * b) * c)"},
		{"a* b /c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 > 4 != 3 < 4", "((5 > 4) != (3 < 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == false", "((3 < 5) == false)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
		{"a + add(b * c) + d", "((a + add((b * c))) + d)"},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}
	for _, test := range tests {
		program := testParseProgram(t, test.input, []string{})
		actual := program.String()
		if actual != test.expected {
			t.Errorf("Exprected=%s Got=%s", test.expected, actual)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x }"
	program := testParseProgram(t, input, []string{})
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
	}
	ifExpr, ok := stmt.Expression.(*ast.IFExpression)
	if !ok {
		t.Fatalf("stmt.Expression not an *ast.IFExpression, Got=%T", stmt.Expression)
	}
	if !testInfixExpression(t, ifExpr.Condition, "x", "<", "y") {
		return
	}
	if len(ifExpr.Consequence.Statements) != 1 {
		t.Fatalf(
			"Unexpected number of statements in ifExpr.Consequence Expected=1 Got=%d",
			len(ifExpr.Consequence.Statements))
	}
	cons, ok := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("cons not an *ast.ExpressionStatement, Got=%T",
			ifExpr.Consequence.Statements[0])
	}
	if !testIdentifier(t, cons.Expression, "x") {
		return
	}
	if ifExpr.Alternative != nil {
		t.Errorf("ifExpr.Alternaive was not nil Got=%+v", ifExpr.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := "if (x < y) { x } else { y }"
	program := testParseProgram(t, input, []string{})
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
	}
	ifExpr, ok := stmt.Expression.(*ast.IFExpression)
	if !ok {
		t.Fatalf("stmt.Expression not an *ast.IFExpression, Got=%T", stmt.Expression)
	}
	if !testInfixExpression(t, ifExpr.Condition, "x", "<", "y") {
		return
	}
	if len(ifExpr.Consequence.Statements) != 1 {
		t.Fatalf(
			"Unexpected number of statements in ifExpr.Consequence Expected=1 Got=%d",
			len(ifExpr.Consequence.Statements))
	}
	cons, ok := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("cons not an *ast.ExpressionStatement, Got=%T",
			ifExpr.Consequence.Statements[0])
	}
	if !testIdentifier(t, cons.Expression, "x") {
		return
	}
	if len(ifExpr.Alternative.Statements) != 1 {
		t.Fatalf(
			"Unexpected number of statements in ifExpr.Alternative Expected=1 Got=%d",
			len(ifExpr.Alternative.Statements))
	}
	alt, ok := ifExpr.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("cons not an *ast.ExpressionStatement, Got=%T",
			ifExpr.Alternative.Statements[0])
	}
	if !testIdentifier(t, alt.Expression, "y") {
		return
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := "fn(x, y) { x + y; }"
	program := testParseProgram(t, input, []string{})
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt not an *ast.ExpressionStatement, Got=%T", program.Statements[0])
	}
	fn, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression not an *ast.FunctionLiteral, Got=%T", stmt.Expression)
	}
	if len(fn.Parameters) != 2 {
		t.Fatalf("Unexpected number of function parameters Expected=2 Got=%d",
			len(fn.Parameters))
	}
	testLiteralExpression(t, fn.Parameters[0], "x")
	testLiteralExpression(t, fn.Parameters[1], "y")
	if len(fn.Body.Statements) != 1 {
		t.Fatalf(
			"Unexpected number of statements in fn.Body Expected=1 Got=%d",
			len(fn.Body.Statements))
	}
	bodyStmt, ok := fn.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("bodyStmt not an *ast.ExpressionStatement, Got=%T", fn.Body.Statements[0])
	}
	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}
	for _, test := range tests {
		program := testParseProgram(t, test.input, []string{})
		testNumberOfStatemets(t, program, 1)
		stmt := program.Statements[0].(*ast.ExpressionStatement)
		fn := stmt.Expression.(*ast.FunctionLiteral)
		if len(fn.Parameters) != len(test.expectedParams) {
			t.Error("Wrong number of fn.Prameters Expected=%d Got=%d",
				len(test.expectedParams), len(fn.Parameters))
		}
		for i, ident := range test.expectedParams {
			testLiteralExpression(t, fn.Parameters[i], ident)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5)"
	program := testParseProgram(t, input, []string{})
	testNumberOfStatemets(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement Got=%T",
			program.Statements[0])
	}
	call, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression Got=%T",
			stmt.Expression)
	}
	if !testIdentifier(t, call.Function, "add") {
		return
	}
	if len(call.Arguments) != 3 {
		t.Fatalf("Wrong number of arguments Exprected=3 Got=%d", len(call.Arguments))
	}
	testLiteralExpression(t, call.Arguments[0], 1)
	testInfixExpression(t, call.Arguments[1], 2, "*", 3)
	testInfixExpression(t, call.Arguments[2], 4, "+", 5)
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	intLiteral, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not and ast.IntegerLiteral Got=%T", exp)
		return false
	}
	if intLiteral.Value != value {
		t.Fatalf("Wrong intLiteral.Value Exptected=%d Got=%d", value, intLiteral.Value)
		return false
	}
	if intLiteral.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Fatalf("Wrong intLiteral.Value Exptected=%d Got=%s",
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

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not and *ast.Identifier Got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("Wrong ident.Value Expected=%s Got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("Wrong ident.TokenLiteral() Expected=%s Got=%s",
			value, ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not hadled Got=%T", exp)
	return false
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression Got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator ins not '%s' Got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	b, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp is not ast.Boolean Got=%T", exp)
		return false
	}
	if b.Value != value {
		t.Errorf("Wrong b.Value Expected=%t Got=%t", value, b.Value)
		return false
	}
	if b.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("Wrong b.TokenLiteral() Expected=%t Got=%s", value, b.TokenLiteral())
		return false
	}
	return true
}
