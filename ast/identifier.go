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
