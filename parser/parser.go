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

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := NewLetStatement(parser.currentToken)
	if !(parser.peekToken.Type == token.IDENT) {
		return nil
	}
	return stmt
}
