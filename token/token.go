package token

type Token struct {
	Type    string
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

var keywords = map[string]string{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) string {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
