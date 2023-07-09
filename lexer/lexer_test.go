package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
    let five = 5;
    let ten = 10;
    let add = fn(x, y) {
      x+y;
    }
    let result = add(five, ten);
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for key, value := range tests {
		token := lexer.NextToken()
		if token.Type != value.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				key, value.expectedType, token.Type)
		}

		if token.Literal != value.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				key, value.expectedLiteral, token.Literal)
		}
	}
}
