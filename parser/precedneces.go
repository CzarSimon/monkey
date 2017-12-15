package parser

import "github.com/CzarSimon/monkey/token"

// precedences Maps a TokenType to a given precedence
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGRATER,
	token.GT:       LESSGRATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.DIVIDE:   PRODUCT,
	token.MULTIPLY: PRODUCT,
}
