package lexer

import (
	"github.com/CzarSimon/monkey/token"
)

// ByteToTypeMap Map of characters represented as bytes to its corresponding token
type ByteToTypeMap map[byte]token.TokenType

// NewByteToTypeMap Creates a new ByteToTypeMap
func NewByteToTypeMap() ByteToTypeMap {
	return ByteToTypeMap{
		';': token.SEMICOLON,
		'(': token.LPAREN,
		')': token.RPAREN,
		',': token.COMMA,
		'+': token.PLUS,
		'{': token.LBRACE,
		'}': token.RBRACE,
		'-': token.MINUS,
		'*': token.MULTIPLY,
		'/': token.DIVIDE,
		'<': token.LT,
		'>': token.GT,
	}
}
