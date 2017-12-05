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

// CurrentChar Gets the current character as a string
func (lexer Lexer) CurrentChar() string {
	return string(lexer.currentChar)
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
	lexer.skipWhitespace()
	switch lexer.currentChar {
	case '=':
		nextToken = token.New(token.ASSIGN, lexer.CurrentChar())
	case ';':
		nextToken = token.New(token.SEMICOLON, lexer.CurrentChar())
	case '(':
		nextToken = token.New(token.LPAREN, lexer.CurrentChar())
	case ')':
		nextToken = token.New(token.RPAREN, lexer.CurrentChar())
	case ',':
		nextToken = token.New(token.COMMA, lexer.CurrentChar())
	case '+':
		nextToken = token.New(token.PLUS, lexer.CurrentChar())
	case '{':
		nextToken = token.New(token.LBRACE, lexer.CurrentChar())
	case '}':
		nextToken = token.New(token.RBRACE, lexer.CurrentChar())
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			nextToken.Literal = lexer.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)
			return nextToken
		} else if isDigit(lexer.currentChar) {
			return token.New(token.INT, lexer.readNumber())
		} else {
			nextToken = token.New(token.ILLEGAL, lexer.CurrentChar())
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

// readIdentifier Reads identifier name from input
func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.position
	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[startPosition:lexer.position]
}

// readNumber Reads a number from input
func (lexer *Lexer) readNumber() string {
	startPosition := lexer.position
	for isDigit(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[startPosition:lexer.position]
}

// skipWhitespace Skips over whitespace characters in input
func (lexer *Lexer) skipWhitespace() {
	for isWhitespace(lexer.currentChar) {
		lexer.readChar()
	}
}

// isWhitespace Checks if a character is considered a whitespace character
func isWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}
