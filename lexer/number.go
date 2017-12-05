package lexer

// isDigit Checks if a character is a digit
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
