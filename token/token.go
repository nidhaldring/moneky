package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	COMMA     = ","
	SEMICOLON = ";"

	LET   = "LET"
	IDENT = "IDENT"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	INT  = "INT"
	FUNC = "FUNC"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"
)

type Token struct {
	Type    TokenType
	Literal string
}

var words = map[string]TokenType{
	"fn":  FUNC,
	"let": LET,
}

func LookupWordTokenType(w string) TokenType {
	t, ok := words[w]
	if !ok {
		return IDENT
	}
	return t
}
