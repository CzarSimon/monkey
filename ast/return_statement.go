package ast

import (
	"bytes"

	"github.com/CzarSimon/monkey/token"
)

// ReturnStatement AST node for retrning the value of an expression
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (returnStmt ReturnStatement) statementNode() {}

// TokenLiteral Retruns the node token literal
func (returnStmt ReturnStatement) TokenLiteral() string {
	return returnStmt.Token.Literal
}

// NewReturnStatement Creates a new ReturnStatement and retruns a reference to it
func NewReturnStatement(tok token.Token) *ReturnStatement {
	return &ReturnStatement{
		Token: tok,
	}
}

// String Returns a string representation of the ReturnStatement node
func (returnStmt *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(returnStmt.TokenLiteral() + " ")
	if returnStmt.ReturnValue != nil {
		out.WriteString(returnStmt.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
