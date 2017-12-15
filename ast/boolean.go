package ast

import "github.com/CzarSimon/monkey/token"

// Boolean AST node represting a boolen expression
type Boolean struct {
	Token token.Token
	Value bool
}

func (boolean *Boolean) expressionNode() {}

func (boolean *Boolean) TokenLiteral() string {
	return boolean.Token.Literal
}

func (boolean *Boolean) String() string {
	return boolean.TokenLiteral()
}

func NewBoolean(tok token.Token) *Boolean {
	return &Boolean{
		Token: tok,
		Value: (tok.Type == token.TRUE),
	}
}
