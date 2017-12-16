package ast

import (
	"bytes"
	"strings"

	"github.com/CzarSimon/monkey/token"
)

// FunctionLiteral AST node for representing function declaration
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fn *FunctionLiteral) expressionNode() {}

func (fn *FunctionLiteral) TokenLiteral() string {
	return fn.Token.Literal
}

// Stirng Retrurns a string representation of a FunctionLiteral
func (fn *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, 0)
	for _, param := range fn.Parameters {
		params = append(params, param.String())
	}
	out.WriteString(fn.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fn.Body.String())
	return out.String()
}

// NewFunctionLiteral Creates a new function liteal and retruns a reference to it
func NewFunctionLiteral(tok token.Token) *FunctionLiteral {
	return &FunctionLiteral{
		Token:      tok,
		Parameters: make([]*Identifier, 0),
	}
}

// AddParam Adds a parameter to the the lisf of function parameters
func (fn *FunctionLiteral) AddParam(param *Identifier) {
	fn.Parameters = append(fn.Parameters, param)
}
