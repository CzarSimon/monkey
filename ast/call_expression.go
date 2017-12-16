package ast

import (
	"bytes"
	"strings"

	"github.com/CzarSimon/monkey/token"
)

// CallExpression AST node for a function call
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (call *CallExpression) expressionNode() {}

func (call *CallExpression) TokenLiteral() string {
	return call.Token.Literal
}

// String Retruns a string representation of a CallExpression
func (call *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, arg := range call.Arguments {
		args = append(args, arg.String())
	}
	out.WriteString(call.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// NewCallExpression Creates a new CallExpression and retruns a reference to it
func NewCallExpression(tok token.Token, function Expression) *CallExpression {
	return &CallExpression{
		Token:    tok,
		Function: function,
	}
}
