package lexer

// isLetter Checks if a given charachter is a letter
func isLetter(char byte) bool {
	return isLowerCaseLetter(char) || isUpperCaseLetter(char) || isSpecialCharacter(char)
}

// isLowerCaseLetter Checks if a given charachter is a lower case letter
func isLowerCaseLetter(char byte) bool {
	return 'a' <= char && char <= 'z'
}

// isUpperCaseLetter Checks if a given charachter is an upper case letter
func isUpperCaseLetter(char byte) bool {
	return 'A' <= char && char <= 'Z'
}

// isSpecialCharacter Checks if a charachter is considered non-alphabetic letter
func isSpecialCharacter(char byte) bool {
	return char == '_' || char == '?'
}
