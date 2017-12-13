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

func (intLiteral *IntegerLiteral) TokenLiteral() string {
	return intLiteral.Token.Literal
}

func (intLiteral *IntegerLiteral) String() string {
	return intLiteral.TokenLiteral()
}

// NewIntegerLiteral Creates a new lite
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
