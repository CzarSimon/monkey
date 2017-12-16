package ast

import (
	"testing"

	"github.com/CzarSimon/monkey/token"
)

func TestString(t *testing.T) {
	program := NewProgram()
	letStmt := NewLetStatement(token.New(token.LET, "let"))
	letStmt.Name = NewIdentifier(token.New(token.IDENT, "myVar"), "myVar")
	letStmt.Value = NewIdentifier(token.New(token.IDENT, "anotherVar"), "anotherVar")
	program.AddStatements(letStmt)

	exprectedProgramString := "let myVar = anotherVar;"
	if program.String() != exprectedProgramString {
		t.Errorf("program.String() wrong. Exprexted=[ %s ] Got=[ %s ]",
			exprectedProgramString, program.String())
	}
}

func TestIdentifier(t *testing.T) {
	tok := token.New(token.IDENT, "x")
	id := NewIdentifier(tok, "x")
	id.expressionNode()
	if id.TokenLiteral() != "x" {
		t.Errorf("identifier.TokenLiteral wrong. Exprected=%s Got=%s", "x", id.TokenLiteral())
	}
}

func TestProgram(t *testing.T) {
	program := NewProgram()
	if program.TokenLiteral() != "" {
		t.Errorf("program.TokenLiteral wrong. Exprected blank Got=%s",
			program.TokenLiteral())
	}
	stmt := NewLetStatement(token.New(token.LET, "x"))
	program.AddStatements(stmt)
	if program.TokenLiteral() != "x" {
		t.Errorf("program.TokenLiteral wrong. Exprected=x Got=%s",
			program.TokenLiteral())
	}
}

func TestRetrunStatement(t *testing.T) {
	stmt := NewReturnStatement(token.New(token.RETRUN, "return"))
	stmt.ReturnValue = NewIdentifier(token.New(token.IDENT, "result"), "result")
	stmt.statementNode()
	if stmt.TokenLiteral() != "return" {
		t.Errorf("returnStatement.TokenLiteral wrong. Exprected='return' Got=%s",
			stmt.TokenLiteral())
	}
	exprectedReturnString := "return result;"
	if stmt.String() != exprectedReturnString {
		t.Errorf("returnStatement.String() wrong. Exprexted=[ %s ] Got=[ %s ]",
			exprectedReturnString, stmt.String())
	}
}

func TestExpressionStatement(t *testing.T) {
	tok := token.New(token.IDENT, "x")
	stmt := NewExpressionStatement(tok)
	stmt.statementNode()
	if stmt.TokenLiteral() != "x" {
		t.Errorf("expressionStatement.TokenLiteral wrong. Exprected=x Got=%s",
			stmt.TokenLiteral())
	}
	if stmt.String() != "" {
		t.Errorf("expressionStatement.String wrong. Exprected blank Got=%s",
			stmt.String())
	}
	stmt.Expression = NewIdentifier(tok, "x")
	if stmt.String() != "x" {
		t.Errorf("expressionStatement.String wrong. Exprected=x Got=%s",
			stmt.String())
	}
}

func TestIntegerLiteral(t *testing.T) {
	intLit, err := NewIntegerLiteral(token.New(token.FUNCTION, "5"))
	if intLit != nil {
		t.Errorf("Expected intLit to be got=[ %+q ]", intLit)
	}
	if err == nil {
		t.Errorf(
			"Expected error to be=[ Unexpected TokenType. Expected=INT Got=%s ] Got=nil",
			token.FUNCTION)
	}
	intLit, err = NewIntegerLiteral(token.New(token.INT, "badValue"))
	if intLit != nil {
		t.Errorf("Expected intLit to be nil got=[ %+q ]", intLit)
	}
	if err == nil {
		t.Errorf("Expected parse error got nil")
	}
	intLit, err = NewIntegerLiteral(token.New(token.INT, "10"))
	if intLit == nil {
		t.Errorf("Expected non-nil IntegerLiteral Got nil")
	}
	if err != nil {
		t.Errorf("Expected no error, Got=[ %s ]", err.Error())
	}
	intLit.expressionNode()
	if intLit.TokenLiteral() != "10" {
		t.Errorf("intLit.TokenLiteral() wrong. Expected=10, Got=%s",
			intLit.TokenLiteral())
	}
	if intLit.String() != "10" {
		t.Errorf("intLit.String() wrong. Expected=10, Got=%s", intLit.String())
	}
	if intLit.Value != 10 {
		t.Errorf("intLit.Valeu wrong. Expected=10, Got=%d", intLit.Value)
	}
}

func TestPrefixExpression(t *testing.T) {
	prefixExpr, err := NewPrefixExpression(token.New(token.COMMA, ","))
	if prefixExpr != nil {
		t.Errorf("Expected prefixExpr to be nil got=[ %+q ]", prefixExpr)
	}
	if err == nil {
		t.Errorf(
			"Expected error to be=[ Invalid token type %s supplied ] Got=nil",
			token.COMMA)
	}
	validTokens := []struct {
		token          token.Token
		literal        string
		expectedString string
	}{
		{token.New(token.NOT, "!"), "!", "(!5)"},
		{token.New(token.MINUS, "-"), "-", "(-5)"},
	}
	for i, validToken := range validTokens {
		prefixExpr, err = NewPrefixExpression(validToken.token)
		if err != nil {
			t.Errorf("%d - Expected no error, Got=[ %s ]", i, err.Error())
		}
		if prefixExpr == nil {
			t.Errorf("Expected non-nil PrefixExpression Got nil")
		}
		prefixExpr.Right, _ = NewIntegerLiteral(token.New(token.INT, "5"))
		prefixExpr.expressionNode()
		if prefixExpr.TokenLiteral() != validToken.literal {
			t.Errorf("%d - Wrong prefixExpr.TokenLiteral(), Expected=%s Got=%s",
				i, validToken.literal, prefixExpr.TokenLiteral())
		}
		if prefixExpr.String() != validToken.expectedString {
			t.Errorf("%d - Wrong prefixExpr.String() Exprected=%s Got=%s",
				i, validToken.expectedString, prefixExpr.String())
		}
	}
}

func TestInfixExpression(t *testing.T) {
	validTokens := []struct {
		token          token.Token
		literal        string
		expectedString string
	}{
		{token.New(token.PLUS, "+"), "+", "(5 + 5)"},
		{token.New(token.MULTIPLY, "*"), "*", "(5 * 5)"},
	}
	intLiteral, _ := NewIntegerLiteral(token.New(token.INT, "5"))
	for i, validToken := range validTokens {
		infixExpr := NewInfixExpression(validToken.token)
		if infixExpr == nil {
			t.Errorf("Expected non-nil InfixExpression Got nil")
		}
		infixExpr.Left = intLiteral
		infixExpr.Right = intLiteral
		infixExpr.expressionNode()
		if infixExpr.TokenLiteral() != validToken.literal {
			t.Errorf("%d - Wrong infixExpr.TokenLiteral(), Expected=%s Got=%s",
				i, validToken.literal, infixExpr.TokenLiteral())
		}
		if infixExpr.String() != validToken.expectedString {
			t.Errorf("%d - Wrong prefixExpr.String() Exprected=%s Got=%s",
				i, validToken.expectedString, infixExpr.String())
		}
	}
}

func TestBoolean(t *testing.T) {
	falseToken := token.New(token.FALSE, "false")
	trueToken := token.New(token.TRUE, "true")
	b := NewBoolean(falseToken)
	if b.Value {
		t.Fatalf("Wrong b.Value Expected=false Got=true")
	}
	b = NewBoolean(trueToken)
	if !b.Value {
		t.Fatalf("Wrong b.Value Expected=true Got=false")
	}
	b.expressionNode()
	if b.String() != "true" {
		t.Fatalf("Wrong b.String Expected=true Got=%s", b.String())
	}
}

func TestBlockStatement(t *testing.T) {
	block := getTestBlockStatment(t)
	block.statementNode()
	if len(block.Statements) != 1 {
		t.Fatalf("Unexpected number of statements. Expected=1, Got=%s", len(block.Statements))
	}
	exprectedString := "x"
	if block.String() != exprectedString {
		t.Fatalf("block.String() wrong Expected=%s Got=%s", exprectedString, block.String())
	}
}

func TestIFExprssion(t *testing.T) {
	ifExpr := NewIFExpression(token.New(token.IF, "if"))
	ifExpr.Condition = NewBoolean(token.New(token.TRUE, "true"))
	ifExpr.Consequence = getTestBlockStatment(t)
	ifExpr.expressionNode()
	if ifExpr.TokenLiteral() != "if" {
		t.Fatalf("ifExpr.TokenLiteral wrong. Expected=if Got=%s", ifExpr.TokenLiteral())
	}
	expectedStr1 := "if true x"
	if ifExpr.String() != expectedStr1 {
		t.Fatalf("ifExpr.String() wrong Expected=[ %s ] Got= [ %s ]",
			expectedStr1, ifExpr.String())
	}
	ifExpr.Alternative = getTestBlockStatment(t)
	expectedStr2 := "if true x else x"
	if ifExpr.String() != expectedStr2 {
		t.Fatalf("ifExpr.String() wrong Expected=[ %s ] Got= [ %s ]",
			expectedStr2, ifExpr.String())
	}
}

func TestFunctionLiteral(t *testing.T) {
	fn := NewFunctionLiteral(token.New(token.FUNCTION, "fn"))
	fn.AddParam(NewIdentifier(token.New(token.IDENT, "x"), "x"))
	fn.AddParam(NewIdentifier(token.New(token.IDENT, "y"), "y"))
	fn.Body = getTestBlockStatment(t)
	fn.expressionNode()
	if fn.TokenLiteral() != "fn" {
		t.Fatalf("fn.TokenLiteral wrong Exprected=fn Got=%s", fn.TokenLiteral())
	}
	expectedStr := "fn(x, y) x"
	if fn.String() != expectedStr {
		t.Fatalf("fn.String() wrong Expected=[ %s ] Got= [ %s ]",
			expectedStr, fn.String())
	}
}

func getTestBlockStatment(t *testing.T) *BlockStatement {
	block := NewBlockStatement(token.New(token.LBRACE, "{"))
	if block.TokenLiteral() != "{" {
		t.Fatalf("block.TokenLiteral wrong. Expected={ Got=%s", block.TokenLiteral())
	}
	stmt := NewExpressionStatement(token.New(token.IDENT, "x"))
	stmt.Expression = NewIdentifier(token.New(token.IDENT, "x"), "x")
	block.AddStatements(stmt)
	return block
}

func TestCallExpression(t *testing.T) {
	call := NewCallExpression(
		token.New(token.LPAREN, "("),
		NewIdentifier(token.New(token.IDENT, "add"), "add"))
	call.Arguments = []Expression{
		NewIdentifier(token.New(token.IDENT, "x"), "x"),
		NewIdentifier(token.New(token.IDENT, "y"), "y"),
	}
	call.expressionNode()
	if call.TokenLiteral() != "(" {
		t.Fatalf("Wrong call.TokenLiteral() Exprected=( Got=%s", call.TokenLiteral())
	}
	expectedStr := "add(x, y)"
	if call.String() != expectedStr {
		t.Fatalf("call.String() wrong Expected=[ %s ] Got= [ %s ]",
			expectedStr, call.String())
	}
}
