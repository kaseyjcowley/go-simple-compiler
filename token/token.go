package token

import (
	"fmt"
	"strings"
)

// Define some token enums for making life easier

type TokenType int

const (
	// Reserved words
	VOID TokenType = iota
	MAIN
	INT
	COUT
	IF
	WHILE
	FOR
	DO
	OR
	AND
	BEGIN
	END
	ENDL

	// Relational operators
	LESS
	LESSEQUAL
	GREATER
	GREATEREQUAL
	EQUAL
	NOTEQUAL

	// Other operators
	INSERTION
	ASSIGNMENT
	PLUS
	MINUS
	TIMES
	DIVIDE
	PLUSEQUAL
	MINUSEQUAL

	// Other characters
	SEMICOLON
	LEFTPAREN
	RIGHTPAREN
	LEFTCURLY
	RIGHTCURLY

	// Other token types
	IDENTIFIER
	INTEGER
	BAD
	ENDFILE

	LAST
)

// Some global variables
var (
	tokenTypeNames, reservedWords []string
)

// Init some of our variables
func init() {
	tokenTypeNames = []string{
		"VOID", "MAIN", "INT", "COUT", "IF", "WHILE", "FOR", "DO", "OR", "AND", "BEGIN", "END", "ENDL",
		"LESS", "LESSEQUAL", "GREATER", "GREATEREQUAL", "EQUAL", "NOTEQUAL",
		"INSERTION", "ASSIGNMENT", "PLUS", "MINUS", "TIMES", "DIVIDE", "PLUSEQUAL", "MINUSEQUAL",
		"SEMICOLON", "LEFTPAREN", "RIGHTPAREN", "LEFTCURLY", "RIGHTCURLY",
		"IDENTIFIER", "INTEGER", "BAD", "ENDFILE",
		"LAST",
	}

	reservedWords = []string{
		"VOID", "MAIN", "INT", "COUT", "IF", "WHILE", "FOR", "DO", "OR", "AND", "BEGIN", "END", "ENDL",
	}
}

// New Types
type Token struct {
	Type   TokenType
	Lexeme string
}

// Constructor
func New(t TokenType, l string) *Token {
	token := &Token{
		Type:   t,
		Lexeme: l,
	}
	token.CheckReserved()

	return token
}

func (t *Token) GetTokenType() TokenType {
	return t.Type
}

func (t *Token) GetTokenTypeName() string {
	return tokenTypeNames[t.Type]
}

func (t *Token) GetLexeme() string {
	return t.Lexeme
}

func (t *Token) CheckReserved() {
	count := len(reservedWords)

	for i := 0; i < count; i++ {
		if strings.ToLower(t.Lexeme) == strings.ToLower(reservedWords[i]) {
			t.Type = TokenType(i)
			break
		}
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("Type: %v, Lexeme: %s\n", t.Type, t.Lexeme)
}
