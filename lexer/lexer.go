package lexer

import (
	"learn-interpreter/token"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         rune
	line         int
	row          int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) newToken(tokenType token.TokenType, ch rune) token.Token {
	c := string(ch)
	if ch == 0 {
		c = ""
	}
	return token.Token{Type: tokenType, Literal: c, Line: l.line, Row: l.row}
}

func (l *Lexer) skipWhitespace() {
	for strings.ContainsRune(" \t\n\r", l.char) {
		if l.char == '\n' {
			l.line += 1
			l.row = 0
		}
		l.readChar()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	var width int
	if l.readPosition >= len(l.input) {
		l.char = 0
		width = 1
	} else {
		runeValue, w := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.char = runeValue
		width = w
	}
	l.position = l.readPosition
	l.readPosition += width
	l.row += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		t = l.newToken(token.ASSIGN, l.char)
	case ';':
		t = l.newToken(token.SEMICOLON, l.char)
	case '(':
		t = l.newToken(token.LPAREN, l.char)
	case ')':
		t = l.newToken(token.RPAREN, l.char)
	case ',':
		t = l.newToken(token.COMMA, l.char)
	case '+':
		t = l.newToken(token.PLUS, l.char)
	case '{':
		t = l.newToken(token.LBRACE, l.char)
	case '}':
		t = l.newToken(token.RBRACE, l.char)
	case 0:
		t = l.newToken(token.EOF, 0)
	default:
		if unicode.IsLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if unicode.IsDigit(l.char) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = l.newToken(token.ILLEGAL, l.char)
		}
	}
	l.readChar()
	return t
}
