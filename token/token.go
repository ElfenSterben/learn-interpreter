package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Row     int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符 + 字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN = "="
	PLUS   = "+"
	DIV    = "-"
	MUL    = "*"
	MOD    = "/"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tt, ok := keywords[ident]; ok {
		return tt
	}
	return IDENT
}
