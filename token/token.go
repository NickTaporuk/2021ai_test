package token

const (
	LBRACKETRUNE = '['
	RBRACKETRUNE = ']'
)

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // fields, table_name

	// Keywords
	SUM
	DIF
	INT

	//
	LBRACKET
	RBRACKET
)
