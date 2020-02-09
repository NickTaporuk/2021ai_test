package parser

import (
	"io"

	"github.com/NickTaporuk/2021ai_test/scanner"
	"github.com/NickTaporuk/2021ai_test/token"
)

// Parser represents a parser.
type Parser struct {
	s   *scanner.Scanner
	buf struct {
		tok token.Token // last read token
		lit string      // last read literal
		n   int         // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: scanner.NewScanner(r)}
}

type ParsedStatement struct {
	Fields    []string
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (*ParsedStatement, error) {
	stmt := &ParsedStatement{}

	// Next we should loop over all our comma-delimited fields.
	for {
		// Read a field.
		tok, lit, err := p.scanIgnoreWhitespace()

		if tok == token.EOF {
			break
		}

		if err != nil {
			return nil, err
		}


		stmt.Fields = append(stmt.Fields,lit)
	}

	// Return the successfully parsed statement.
	return stmt, nil
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok token.Token, lit string, err error) {
	tok, lit, err = p.scan()
	if tok == token.WS {
		tok, lit, err = p.scan()
	}

	return
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok token.Token, lit string, err error) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit, err
	}

	// Otherwise read the next token from the scanner.
	tok, lit, err = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
