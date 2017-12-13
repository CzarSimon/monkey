package parser

import (
	"errors"
	"fmt"

	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/token"
)

// Integer constants denoting operator precedence
const (
	_ int = iota
	LOWEST
	EQUALS     // ==
	LESSGRATER // > OR <
	SUM        // +
	PRODUCT    // *
	PREFIX     // -X or !X
	CALL       // myFunc(X)
)

// Parser A series of tokens into an abstract source tree (AST)
type Parser struct {
	lex            *lexer.Lexer
	errors         []error
	currentToken   token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New Creates a new parser based on a supplied lexer
func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:            lex,
		errors:         make([]error, 0),
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}
	parser.registerPrefix(token.IDENT, parser.parseIdentifier)
	parser.nextToken()
	parser.nextToken()
	return parser
}

// nextToken Moves the tokens in the parser forward
func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

// ParseProgram Parses a programed described in the supplied lexer
func (parser *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()
	for !parser.currentTokenIs(token.EOF) {
		stmt, err := parser.parseStatement()
		if err != nil {
			parser.AddError(err)
		} else if stmt != nil {
			program.AddStatements(stmt)
		}
		parser.nextToken()
	}
	return program
}

// expectPeek Checks the type of peekToken and andvances the token pointers if the type was expected
func (parser *Parser) expectPeek(tokenType token.TokenType) error {
	if parser.peekTokenIs(tokenType) {
		parser.nextToken()
		return nil
	}
	return parser.peekError(tokenType)
}

// currentTokenIs Checks if currentToken is of a supplied type
func (parser *Parser) currentTokenIs(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

// peekTokenIs Checks if peekToken is of a supplied type
func (parser *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return parser.peekToken.Type == tokenType
}

// Errors Retruns the list of ParseErrors
func (parser *Parser) Errors() []error {
	return parser.errors
}

// AddError Adds a parse error to the parsers list of errors
func (parser *Parser) AddError(err error) {
	parser.errors = append(parser.errors, err)
}

// peekError Adds an error caused by unexpected token type
func (parser *Parser) peekError(tokenType token.TokenType) error {
	return errors.New(fmt.Sprintf(
		"peekToken: Expected type=%s Got=%s", tokenType, parser.peekToken.Type))
}

// registerPrefix Adds a prefix function to a particular token type
func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

// getCurrentTokensPrefixFn Gets the prefix function of the currentToken
func (parser *Parser) getCurrentTokensPrefixFn() prefixParseFn {
	return parser.prefixParseFns[parser.currentToken.Type]
}

// registerInfix Adds a infix function to a particular token type
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

// getCurrentTokensinfixFn Gets the infix function of the currentToken
func (parser *Parser) getCurrentTokensInfixFn() infixParseFn {
	return parser.infixParseFns[parser.currentToken.Type]
}
