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

	// 数学运算符
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULTI  = "*"
	DIV    = "/"
	MOD    = "%"

	// 逻辑运算符
	BANG   = "!"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tt, ok := keywords[ident]; ok {
		return tt
	}
	return IDENT
}
