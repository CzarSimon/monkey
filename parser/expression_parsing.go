package parser

import (
	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// parseExpression Parses an expresion based on the state of the parser an the supplied precedence
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.getCurrentTokensPrefixFn()
	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	for !parser.peekTokenIs(token.SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.getPeekTokensInfixFn()
		if infix == nil {
			return leftExp
		}
		parser.nextToken()
		leftExp = infix(leftExp)
	}
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

// parsePrefixExpression Parses an PrefixExpression
func (parser *Parser) parsePrefixExpression() ast.Expression {
	prefixExpr, err := ast.NewPrefixExpression(parser.currentToken)
	if err != nil {
		parser.AddError(err)
		return nil
	}
	parser.nextToken()
	prefixExpr.Right = parser.parseExpression(PREFIX)
	return prefixExpr
}

// parseInfixExpression Parses an InfixExpression
func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := ast.NewInfixExpression(parser.currentToken)
	expression.Left = left
	precedence := parser.currentPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedence)
	return expression
}
