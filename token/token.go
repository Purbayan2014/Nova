package token

type tokenType string

type token struct {
	Type    tokenType
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
	ASSIGN  = "="
	PLUS    = "+"
	SUB     = "-"
	MULTI   = "*"
	DIV     = "/"
	LESS    = "<"
	GREATER = ">"

	// Kewords
	FUNCTION = "fn"
	LETVAR   = "letvar"
)
