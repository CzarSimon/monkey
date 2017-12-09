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
  @
  !-/*5;
  5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	10 == 10;
	10 != 9;
  `
	tests := []expectedTokenType{
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="},
		{token.INT, "5"}, {token.SEMICOLON, ";"}, {token.LET, "let"},
		{token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"},
		{token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "add"},
		{token.ASSIGN, "="}, {token.FUNCTION, "fn"}, {token.LPAREN, "("},
		{token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"},
		{token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.IDENT, "x"},
		{token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.SEMICOLON, ";"}, {token.LET, "let"},
		{token.IDENT, "result"}, {token.ASSIGN, "="}, {token.IDENT, "add"},
		{token.LPAREN, "("}, {token.IDENT, "five"}, {token.COMMA, ","},
		{token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
		{token.ILLEGAL, "@"}, {token.NOT, "!"}, {token.MINUS, "-"},
		{token.DIVIDE, "/"}, {token.MULTIPLY, "*"}, {token.INT, "5"},
		{token.SEMICOLON, ";"}, {token.INT, "5"}, {token.LT, "<"},
		{token.INT, "10"}, {token.GT, ">"}, {token.INT, "5"},
		{token.SEMICOLON, ";"}, {token.IF, "if"}, {token.LPAREN, "("},
		{token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"},
		{token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.RETRUN, "return"},
		{token.TRUE, "true"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"},
		{token.ELSE, "else"}, {token.LBRACE, "{"}, {token.RETRUN, "return"},
		{token.FALSE, "false"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"},
		{token.INT, "10"}, {token.EQ, "=="}, {token.INT, "10"},
		{token.SEMICOLON, ";"}, {token.INT, "10"}, {token.NOT_EQ, "!="},
		{token.INT, "9"}, {token.SEMICOLON, ";"}, {token.EOF, ""},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q",
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

func TestCurrentChar(t *testing.T) {
	lexer := New("+-")
	if lexer.CurrentChar() != "+" {
		t.Fatalf("Wrong CurrentChar (as string): Expected=+ got=%s", lexer.CurrentChar())
	}
	lexer.readChar()
	if lexer.CurrentChar() != "-" {
		t.Fatalf("Wrong CurrentChar (as string): Expected=- got=%s", lexer.CurrentChar())
	}
}
