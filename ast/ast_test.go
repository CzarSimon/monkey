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