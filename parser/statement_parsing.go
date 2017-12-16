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
	case token.RETRUN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
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
	parser.nextToken()
	stmt.Value = parser.parseExpression(LOWEST)
	for parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt, nil
}

// parseReturnStatement Parses a ReturnStatement
func (parser *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := ast.NewReturnStatement(parser.currentToken)
	parser.nextToken()
	stmt.ReturnValue = parser.parseExpression(LOWEST)
	for parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt, nil
}

// parseExpressionStatement Parses an ExpressionStatement
func (parser *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := ast.NewExpressionStatement(parser.currentToken)
	stmt.Expression = parser.parseExpression(LOWEST)
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt, nil
}

// parseBlockStatement Parses a BlockStatement
func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	block := ast.NewBlockStatement(parser.currentToken)
	parser.nextToken()
	if !parser.currentTokenIs(token.RBRACE) && !parser.currentTokenIs(token.EOF) {
		stmt, err := parser.parseStatement()
		if err != nil {
			parser.AddError(err)
		} else if stmt != nil {
			block.AddStatements(stmt)
		}
		parser.nextToken()
	}
	return block
}
