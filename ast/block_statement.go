package ast

import (
	"bytes"

	"github.com/CzarSimon/monkey/token"
)

// BlockStatement AST node representing a grouped series of statements
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (block *BlockStatement) statementNode() {}

func (block *BlockStatement) TokenLiteral() string {
	return block.Token.Literal
}

// String Returns string representation of a BlockStatement
func (block *BlockStatement) String() string {
	var out bytes.Buffer
	for _, stmt := range block.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

// NewBlockStatement Cretes a new, empty BlockStatement and returns a reference to it
func NewBlockStatement(tok token.Token) *BlockStatement {
	return &BlockStatement{
		Token:      tok,
		Statements: make([]Statement, 0),
	}
}

// AddStatements Adds statements to the block
func (block *BlockStatement) AddStatements(stmts ...Statement) {
	block.Statements = append(block.Statements, stmts...)
}
