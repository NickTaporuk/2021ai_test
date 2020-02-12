package parser

import (
	"io"

	"github.com/NickTaporuk/2021ai_test/scanner"
	"github.com/NickTaporuk/2021ai_test/set"
)

// Parser represents a parser.
type Parser struct {
	s scanner.Scanner
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: scanner.NewLexicalScanner(r)}
}

// Parse
func (p *Parser) Parse() (set.Interface, error) {

	data, err := p.s.Scan()
	return data, err
}
