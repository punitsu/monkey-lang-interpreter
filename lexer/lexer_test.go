package lexer

import (
	"monkey-lang-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x,y) {
		x + y;
	}
	let result = add(five, ten);
	`
	tests := []struct {
		expectedType     token.TokenType
		exepectedLiteral string
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

	l := New(input)
	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected =%q, got = %q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.exepectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected =%q, got = %q", i, tt.exepectedLiteral, token.Literal)
		}
	}

}
