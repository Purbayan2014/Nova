package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// types of tokens availables
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "ident"
	INT   = "int"
	FRAC  = "frac"

	// Separators
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LSQUARE   = "["
	RSQUARE   = "]"

	// Operators
	ASSIGN        = "="
	PLUS          = "+"
	SUB           = "-"
	MULTI         = "*"
	DIV           = "/"
	LESS          = "<"
	BANG          = "!"
	GREATER       = ">"
	EQ            = "=="
	NOTEQ         = "!="
	LESSTHANEQ    = "<="
	GREATERTHANEQ = ">="

	// Kewords
	FUNCTION = "fn"
	LETVAR   = "letvar"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"letvar": LETVAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
