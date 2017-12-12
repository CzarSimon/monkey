package parser

import (
	"errors"
	"fmt"

	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/token"
)

// Parser A series of tokens into an abstract source tree (AST)
type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []error
}

// New Creates a new parser based on a supplied lexer
func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:    lex,
		errors: make([]error, 0),
	}
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
