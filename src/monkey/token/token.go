package token

// AbstractType is either of ILLEGAL, EOF, IDENT, etc
type AbstractType string

// Token is a tupe of AbstractType and sometimes a value for the Type
type Token struct {
	Type    AbstractType
	Literal string
}

const (
	// ILLEGAL any char that is not allowd in Monkey
	ILLEGAL = "ILLEGAL"

	// EOF as in "E"nd "O"f "F"un
	EOF = "EOF"

	// IDENT are Identifiers like add, foobar, x, y
	IDENT = "IDENT"

	// INT is currently the only literal
	INT = "INT"

	// ASSIGN assigns values to identifiers
	ASSIGN = "="

	// PLUS is the special character for adding two numbers
	PLUS = "+"

	// COMMA separates identifiers
	COMMA = ","

	// SEMICOLON ends the evalution
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION declares a function
	FUNCTION = "FUNCTION"

	// LET it be
	LET = "LET"
)
