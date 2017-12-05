package lexer

import (
	"testing"

	"github.com/CzarSimon/monkey/token"
)

type expectedTokenType struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `
  let five = 5;
  let ten = 10;
  let add = fn(x, y) {
    x + y;
  };
  let result = add(five, ten);
  `
	tests := []expectedTokenType{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenOnEmptyInput(t *testing.T) {
	lexer := New("")
	expectedToken := token.Token{token.EOF, ""}
	tok := lexer.NextToken()
	if tok.Type != expectedToken.Type {
		t.Fatalf("tests - tokentype wrong. expected=%q, got=%q",
			expectedToken.Type, tok.Type)
	}
	if tok.Literal != expectedToken.Literal {
		t.Fatalf("tests- literal wrong. expected=%q, got=%q",
			expectedToken.Literal, tok.Literal)
	}
}
