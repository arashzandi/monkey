package lexer

import "monkey/token"

// Lexer is a struct
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New method reutns a Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// readChar reads next character
func (l *Lexer) readChar() {
	if l.readPosition == len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken compares the character at hand to with the list of tokens
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// newToken brings back a token
func newToken(tokenType token.AbstractType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
