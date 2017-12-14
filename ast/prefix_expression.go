package ast

import (
	"bytes"
	"fmt"

	"github.com/CzarSimon/monkey/token"
)

// PrefixExpression AST node representing an expression preceeded
// by a prefix operator
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (prefixExpr *PrefixExpression) expressionNode() {}

func (prefixExpr *PrefixExpression) TokenLiteral() string {
	return prefixExpr.Token.Literal
}

// String Retruns a string representation of a PrefixExpression
func (prefixExpr *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(prefixExpr.Operator)
	out.WriteString(prefixExpr.Right.String())
	out.WriteString(")")
	return out.String()
}

func NewPrefixExpression(tok token.Token) (*PrefixExpression, error) {
	if !isValidTokenTypeForPrefixExpression(tok.Type) {
		return nil, fmt.Errorf("Invalid token type %s supplied", tok.Type)
	}
	return &PrefixExpression{
		Token:    tok,
		Operator: tok.Literal,
	}, nil
}

// isValidTokenTypeForPrefixExpression Checks if a given TokenType is valid
// for creating a PrefixExpression
func isValidTokenTypeForPrefixExpression(tokenType token.TokenType) bool {
	return tokenType == token.MINUS || tokenType == token.NOT
}
