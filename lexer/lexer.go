package lexer

import (
	"monkey-interpreter/token"
)

type Lexer struct { // only supports ASCII characters
	input string
	// reason for having two "pointers" (positing and readPosition) pointing
	// into our input string is the fact that we need to be able to "peek"
	// further into the input and look after the current character to see
	// what comes up next

	// current position in input, which is the position where we last read
	position int
	// current reading position in input, which is the position we're going to read next
	readPosition int
	ch           byte // current char under examination (as pointed by position)
}

// Initialize the lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Give us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code for the "NUL" character and signifies either
		// "we haven't read anything yet" or "end of file" for us.
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Look at the current char under examination (l.ch) and return a token depending
// on which character it is.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// Before returning the token we advance our pointers into the input so when
	// we call NextToken() again the l.ch field is already updated.
	l.readChar()
	return tok
}

// Helps to initialize new tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
