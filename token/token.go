package token

// TokenType Represents a possible type of tokens
type TokenType string

// Token Lexer representation of source code componentes
type Token struct {
	Type    TokenType
	Literal string
}

// LookupIdent Checks if a provided string is a keywords, if so returns its Type
// otherwise identifies it as an identifier
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}

// New Creats a new token based on a given TokenType and string literal
func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// keywords Map of keywords to token type
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETRUN,
	"true":   TRUE,
	"false":  FALSE,
}
