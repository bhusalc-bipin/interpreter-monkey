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

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// recognize whether the current character is a letter and if so, read the
	// rest of the identifier/keyword until a non-letter-character is encountered.
	// Having read that identifier/keyword, find out if it is a identifier or a
	// keyword, so we can use the correct 'token.TokenType'.
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			// early exit here is necessay because when calling readIdentifier(),
			// we call readChar() repeatedly and advance 'readPosition' and 'position'
			// fields past the last character of the current identifier. So, we
			// don't need the call to readChar() after the switch statement again.
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT // supports integer only
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
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

// reads in an identifier and advances the lexer's position until it encounters
// a non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// specify if a characrer is allowed/valid in identifiers and keywords
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' // monkey language supports only integers
}

// We only want to 'peek' ahead in the input and not move around in it. This method
// is similar to readChar(), except that it doesn't increment l.position and l.readPosition.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
