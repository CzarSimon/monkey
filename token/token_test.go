package token

import "testing"

func TestLookupIdent(t *testing.T) {
	testStrings := []string{
		"fn", "add", "let", "present?",
	}
	expectedTokenTypes := []TokenType{
		FUNCTION, IDENT, LET, IDENT,
	}
	for i, testString := range testStrings {
		tokenType := LookupIdent(testString)
		expectedTokenType := expectedTokenTypes[i]
		if tokenType != expectedTokenType {
			t.Fatalf("Test[%d] - Wrong TokenType, expected=%q got=%q",
				i, expectedTokenType, tokenType)
		}
	}
}

func TestNew(t *testing.T) {
	token := New(LPAREN, "{")
	expectedLiteral := "{"
	if token.Type != LPAREN {
		t.Fatalf("Test- Wrong TokenType, expected=%q got=%q", LPAREN, token.Type)
	}
	if token.Literal != expectedLiteral {
		t.Fatalf("Test- Wrong Literal, expected=%q got=%q", expectedLiteral, token.Literal)
	}
}
