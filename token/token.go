package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "INT"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LBREACE = "{"
	RBRASE  = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
