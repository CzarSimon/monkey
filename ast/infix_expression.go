package ast

import (
	"bytes"

	"github.com/CzarSimon/monkey/token"
)

// InfixExpression AST node represneting an two expressions separated
// by an infix operator
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (infixExpr *InfixExpression) expressionNode() {}

func (infixExpr *InfixExpression) TokenLiteral() string {
	return infixExpr.Token.Literal
}

// String Returns string representation of an InfixExpression
func (infixExpr *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(infixExpr.Left.String())
	out.WriteString(" " + infixExpr.Operator + " ")
	out.WriteString(infixExpr.Right.String())
	out.WriteString(")")
	return out.String()
}

// NewInfixExpression Creates a new InfixExpression based on a supplied token
func NewInfixExpression(tok token.Token) *InfixExpression {
	return &InfixExpression{
		Token:    tok,
		Operator: tok.Literal,
	}
}
