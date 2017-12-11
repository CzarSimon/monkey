package parser

import (
	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/token"
)

// Parser A series of tokens into an abstract source tree (AST)
type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

// New Creates a new parser based on a supplied lexer
func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex: lex,
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
	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}
	return program
}

// parseStatement Parsers a statement from input
func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement Parses a LetStatement
func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.NewLetStatement(parser.currentToken)
	if !parser.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = ast.NewIdentifier(parser.currentToken, parser.currentToken.Literal)
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: Write expression parsing
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

// expectPeek Checks the type of peekToken and andvances the token pointers if the type was expected
func (parser *Parser) expectPeek(tokenType token.TokenType) bool {
	if parser.peekTokenIs(tokenType) {
		parser.nextToken()
		return true
	}
	return false
}

// currentTokenIs Checks if currentToken is of a supplied type
func (parser *Parser) currentTokenIs(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

// peekTokenIs Checks if peekToken is of a supplied type
func (parser *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return parser.peekToken.Type == tokenType
}
