package parser

import (
	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/token"
)

// parseStatement Parsers a statement from input
func (parser *Parser) parseStatement() (ast.Statement, error) {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil, nil
	}
}

// parseLetStatement Parses a LetStatement
func (parser *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := ast.NewLetStatement(parser.currentToken)
	if err := parser.expectPeek(token.IDENT); err != nil {
		return nil, err
	}
	stmt.Name = ast.NewIdentifier(parser.currentToken, parser.currentToken.Literal)
	if err := parser.expectPeek(token.ASSIGN); err != nil {
		return nil, err
	}
	// TODO: Write expression parsing
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt, nil
}
