package ast

import (
	"github.com/CzarSimon/monkey/token"
)

// LetStatement AST node for variable assignement
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (letStmt LetStatement) statementNode() {}

// TokenLiteral Retruns the node token literal
func (letStmt LetStatement) TokenLiteral() string {
	return letStmt.Token.Literal
}
