package ast

import (
	"fmt"
	"strconv"

	"github.com/CzarSimon/monkey/token"
)

// IntegerLiteral AST node for integer values
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (intLiteral *IntegerLiteral) expressionNode() {}

// TokenLiteral Returns the integer literal as a string
func (intLiteral *IntegerLiteral) TokenLiteral() string {
	return intLiteral.Token.Literal
}

// String Returns string representation of IntegerLiteral
func (intLiteral *IntegerLiteral) String() string {
	return intLiteral.TokenLiteral()
}

// NewIntegerLiteral Creates a new IntegerLiteral
func NewIntegerLiteral(tok token.Token) (*IntegerLiteral, error) {
	if tok.Type != token.INT {
		return nil, fmt.Errorf("Unexpected TokenType. Expected=INT Got=%s", tok.Type)
	}
	value, err := strconv.ParseInt(tok.Literal, 0, 64)
	if err != nil {
		return nil, err
	}
	return &IntegerLiteral{
		Token: tok,
		Value: value,
	}, nil
}
