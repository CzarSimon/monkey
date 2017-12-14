package parser

import (
	"github.com/CzarSimon/monkey/ast"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// parseExpression Parses an expresion based on the state of the parser an the supplied precedence
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.getCurrentTokensPrefixFn()
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

// parseIdentifier Parses an Identifier expression
func (parser *Parser) parseIdentifier() ast.Expression {
	return ast.NewIdentifier(parser.currentToken, parser.currentToken.Literal)
}

// parseIntegerLiteral Parses an IntegerLiteral expression
func (parser *Parser) praseIntegerLiteral() ast.Expression {
	literal, err := ast.NewIntegerLiteral(parser.currentToken)
	if err != nil {
		parser.AddError(err)
	}
	return literal
}
