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
	IDENT

	// Keywords
	SUM
	DIF
	INT

	//
	LBRACKET
	RBRACKET

	// File
	FILE
)
