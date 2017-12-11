package ast

import (
	"github.com/CzarSimon/monkey/token"
)

// Identifier AST node for variable identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (id Identifier) expressionNode() {}

// TokenLiteral Retruns the node token literal
func (id Identifier) TokenLiteral() string {
	return id.Token.Literal
}

// NewIdentifier Creates a new identifier and returns its reference
func NewIdentifier(tok token.Token, value string) *Identifier {
	return &Identifier{
		Token: tok,
		Value: value,
	}
}
