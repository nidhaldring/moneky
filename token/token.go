package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	COMMA     = ","
	SEMICOLON = ";"

	LET    = "LET"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IDENT  = "IDENT"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQUAL  = "EQUAl"
	NEQUAL = "NEQUAL"

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
	"func":     FUNC,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupWordTokenType(w string) TokenType {
	t, ok := words[w]
	if !ok {
		return IDENT
	}
	return t
}
