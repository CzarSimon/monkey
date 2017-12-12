package ast

import "github.com/CzarSimon/monkey/token"

// ExpressionStatement AST node for retrning the value of an expression
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (stmt ExpressionStatement) statementNode() {}

// TokenLiteral Retruns the node token literal
func (stmt ExpressionStatement) TokenLiteral() string {
	return stmt.Token.Literal
}

// NewExpressionStatement Creates a new ExpressionStatement and retruns a reference to it
func NewExpressionStatement(tok token.Token) *ExpressionStatement {
	return &ExpressionStatement{
		Token: tok,
	}
}

// String Returns a string representation of the ExpressionStatement node
func (stmt *ExpressionStatement) String() string {
	if stmt.Expression != nil {
		return stmt.Expression.String()
	}
	return ""
}
