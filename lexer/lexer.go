package lexer

import (
	"github.com/CzarSimon/monkey/token"
)

// Lexer Type for converting source code intokens
type Lexer struct {
	input         string
	inputLength   int
	position      int  // current position in the input (points to current char)
	readPosition  int  // current readin gpositon in input (after current char)
	currentChar   byte // current char under examination
	byteToTypeMap ByteToTypeMap
}

// CurrentChar Gets the current character as a string
func (lexer Lexer) CurrentChar() string {
	return string(lexer.currentChar)
}

// New Creates a new lexer based on an input string
func New(input string) *Lexer {
	lexer := &Lexer{
		input:         input,
		inputLength:   len(input),
		byteToTypeMap: NewByteToTypeMap(),
	}
	lexer.readChar()
	return lexer
}

// NextToken Gets the next token from the input
func (lexer *Lexer) NextToken() token.Token {
	lexer.skipWhitespace()
	nextToken, shouldReadNextChar := lexer.buildNextToken()
	if shouldReadNextChar {
		lexer.readChar()
	}
	return nextToken
}

// buildNextToken Constructs the next token and instructs if a further character shold be read
func (lexer *Lexer) buildNextToken() (token.Token, bool) {
	tokenType, isPresent := lexer.byteToTypeMap[lexer.currentChar]
	if isPresent {
		return token.New(tokenType, lexer.CurrentChar()), true
	}
	switch lexer.currentChar {
	case 0:
		return token.New(token.EOF, ""), true
	case '=':
		if lexer.peekChar() != '=' {
			return token.New(token.ASSIGN, lexer.CurrentChar()), true
		}
		previousChar := lexer.CurrentChar()
		lexer.readChar()
		return token.New(token.EQ, previousChar+lexer.CurrentChar()), true
	case '!':
		if lexer.peekChar() != '=' {
			return token.New(token.NOT, lexer.CurrentChar()), true
		}
		previousChar := lexer.CurrentChar()
		lexer.readChar()
		return token.New(token.NOT_EQ, previousChar+lexer.CurrentChar()), true
	default:
		return lexer.handleDefault()
	}
}

// handleDefault Handles the default case for building a token
func (lexer *Lexer) handleDefault() (token.Token, bool) {
	if isLetter(lexer.currentChar) {
		literal := lexer.readIdentifier()
		return token.New(token.LookupIdent(literal), literal), false
	}
	if isDigit(lexer.currentChar) {
		return token.New(token.INT, lexer.readNumber()), false
	}
	return token.New(token.ILLEGAL, lexer.CurrentChar()), true
}

// readChar Reads the current char fo the input string
func (lexer *Lexer) readChar() {
	lexer.currentChar = lexer.peekChar()
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// peekChar Looks up and retruns the next character in input
func (lexer Lexer) peekChar() byte {
	if lexer.readPosition >= lexer.inputLength {
		return 0
	}
	return lexer.input[lexer.readPosition]
}

// readIdentifier Reads identifier name from input
func (lexer *Lexer) readIdentifier() string {
	return lexer.readType(isLetter)
}

// readNumber Reads a number from input
func (lexer *Lexer) readNumber() string {
	return lexer.readType(isDigit)
}

// readType Reads a string of a particular type defined in the method typeCheck
func (lexer *Lexer) readType(typeCheck func(char byte) bool) string {
	startPosition := lexer.position
	for typeCheck(lexer.currentChar) {
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
