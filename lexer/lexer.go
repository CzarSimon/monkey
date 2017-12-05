package lexer

import (
	"github.com/CzarSimon/monkey/token"
)

// Lexer Type for converting source code intokens
type Lexer struct {
	input        string
	inputLength  int
	position     int  // current position in the input (points to current char)
	readPosition int  // current readin gpositon in input (after current char)
	currentChar  byte // current char under examination
}

// New Creates a new lexer based on an input string
func New(input string) *Lexer {
	lexer := &Lexer{
		input:       input,
		inputLength: len(input),
	}
	lexer.readChar()
	return lexer
}

// NextToken Gets the next token from the input
func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token
	switch lexer.currentChar {
	case '=':
		nextToken = newToken(token.ASSIGN, lexer.currentChar)
	case ';':
		nextToken = newToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		nextToken = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		nextToken = newToken(token.RPAREN, lexer.currentChar)
	case ',':
		nextToken = newToken(token.COMMA, lexer.currentChar)
	case '+':
		nextToken = newToken(token.PLUS, lexer.currentChar)
	case '{':
		nextToken = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		nextToken = newToken(token.RBRACE, lexer.currentChar)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			nextToken.Literal = lexer.readIdentifier()
			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, lexer.currentChar)
		}
	}
	lexer.readChar()
	return nextToken
}

// readChar Reads the current char fo the input string
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= lexer.inputLength {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// newToken Creates new token based on a supplied type and a charachter
func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

// readIdentifier Reads identifier name from input
func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.position
	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[startPosition:lexer.position]
}
