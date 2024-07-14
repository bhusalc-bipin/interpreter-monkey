package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// Make life simpler here by using "string" as the type for our source code.
	// In a production environment it makes sense to attach filenames and line
	// numbers to tokens, to better track down lexing and parsing errors. So it
	// would be better to initialize the lexer with an "io.Reader" and the filename.
	// But that would add more complexity, and our goal for now is to learn
	// how interpreter works in general, we start small and just use a "string"
	// and ignore filenames and line numbers.
	input := `=+(){},;`

	// table test
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

	// Initialize the lexer with our source code
	l := New(input)

	for i, tt := range tests {
		// Go through the source code, token by token, character by character
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
