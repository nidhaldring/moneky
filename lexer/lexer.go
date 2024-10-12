package lexer

import (
	"monkey/token"
)

const EOFCHAR byte = 0

type Lexer struct {
	code       string
	chPosition int  // position for ch
	ch         byte // byte testing against
}

func NewLexer(code string) Lexer {
	lex := Lexer{
		code:       code,
		chPosition: 0,
	}

	lex.readChar()
	return lex
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpaces()

	var t token.Token
	switch l.ch {
	case '=':
		if l.Peek() == '=' {
			l.readChar()
			t = token.Token{Type: token.EQUAL, Literal: "=="}
		} else {
			t = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		}
	case ';':
		t = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case ',':
		t = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case '{':
		t = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		t = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case '(':
		t = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		t = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '+':
		t = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '-':
		t = token.Token{Type: token.MINUS, Literal: string(l.ch)}
	case '/':
		t = token.Token{Type: token.SLASH, Literal: string(l.ch)}
	case '*':
		t = token.Token{Type: token.ASTERISK, Literal: string(l.ch)}
	case '<':
		t = token.Token{Type: token.LT, Literal: string(l.ch)}
	case '>':
		t = token.Token{Type: token.GT, Literal: string(l.ch)}
	case EOFCHAR:
		t = token.Token{Type: token.EOF}
		return t
	case '!':
		if l.Peek() == '=' {
			l.readChar()
			t = token.Token{Type: token.NEQUAL, Literal: "!="}
		} else {
			t = token.Token{Type: token.BANG, Literal: string(l.ch)}
		}
	default:
		if isCharacter(l.ch) {
			literal := l.readWord()
			typ := token.LookupWordTokenType(literal)
			t = token.Token{Type: typ, Literal: literal}
			return t
		} else if isNumber(l.ch) {
			literal := l.readNumber()
			t = token.Token{Type: token.INT, Literal: literal}
			return t
		} else {
			t = token.Token{Type: token.ILLEGAL}
		}

	}

	l.readChar()
	return t
}

func (l *Lexer) PeekToken() token.Token {
	currentPos := l.chPosition
	currentCh := l.ch

	tok := l.NextToken()

	l.chPosition = currentPos
	l.ch = currentCh

	return tok
}

func (l *Lexer) skipWhiteSpaces() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readWord() string {
	literal := ""
	for isCharacter(l.ch) {
		literal += string(l.ch)
		l.readChar()
	}

	return literal
}

func (l *Lexer) readNumber() string {
	literal := ""
	for isNumber(l.ch) {
		literal += string(l.ch)
		l.readChar()
	}

	return literal
}

func (l *Lexer) readChar() bool {
	if l.chPosition == len(l.code) {
		l.ch = EOFCHAR
		return false
	}

	l.ch = l.code[l.chPosition]
	l.chPosition++
	return true

}

func (l *Lexer) Peek() byte {
	if l.chPosition == len(l.code) {
		return EOFCHAR
	}
	return l.code[l.chPosition]
}

func isCharacter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
