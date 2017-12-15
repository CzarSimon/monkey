package ast

import (
	"bytes"

	"github.com/CzarSimon/monkey/token"
)

// IFExpression AST node representing an if-else expression
type IFExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifExpr *IFExpression) expressionNode() {}

func (ifExpr *IFExpression) TokenLiteral() string {
	return ifExpr.Token.Literal
}

// String Returns the string represtation of an IFExpression
func (ifExpr *IFExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if ")
	out.WriteString(ifExpr.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExpr.Consequence.String())
	if ifExpr.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ifExpr.Alternative.String())
	}
	return out.String()
}

// NewIFExpression Creates a new IFExpression and returns a reference to it
func NewIFExpression(tok token.Token) *IFExpression {
	return &IFExpression{
		Token: tok,
	}
}
