package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers + literals
	IDENT = "IDENT" // basically variable name
	INT   = "INT"   // Integer type

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// TokenType Represents a possible type of tokens
type TokenType string

// Token Lexer representation of source code componentes
type Token struct {
	Type    TokenType
	Literal string
}
