package scanner

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/NickTaporuk/2021ai_test/token"
)

// Compute
func Compute(in string) (float64, error) {

	//	sc := initScanner(in)
	//	//var prev token.Token = token.ILLEGAL
	//	//var back int = -1
	//
	//Scan:
	//	for {
	//		pos, tok, lit := sc.Scan()
	//
	//		fmt.Println(pos, tok, lit)
	//		//break
	//
	//		switch {
	//		case tok == token.EOF:
	//			break Scan
	//
	//			//case tok == "SUM".(int): {}
	//		}
	//
	//	}

	return 0, nil

}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.

func (s *Scanner) Scan() (tok token.Token, lit string, err error) {
	// Read the next rune.
	ch := s.read()
	return s.checkToken(ch, lit)
}

func (s *Scanner) checkToken(ch rune, lit string) (token.Token, string, error) {
	if isWhitespace(ch) {
		s.unread()
		tok, lit := s.scanWhitespace()
		return tok, lit, nil
	} else if isLetter(ch) {
		s.unread()
		tok, lit := s.scanIdent()
		println(tok)
		switch lit {
		case string(token.DIF):
			return token.DIF, string(token.DIF), nil
		case string(token.SUM):
			return token.SUM, string(token.SUM), nil
		case string(token.INT):
			return token.INT, string(token.INT), nil
		default:
			inspect := string(ch)
			if strings.TrimSpace(lit) != "" {
				inspect += " (`" + lit + "`)"
			}

			return token.Token(ch), string(ch), errors.New("Unrecognized token " + inspect + " in expression")
		}
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return token.EOF, string(eof), nil
	case token.LBRACKETRUNE:
		{
			return token.LBRACKET, string(token.LBRACKETRUNE), nil
		}
	case token.RBRACKETRUNE:
		{
			return token.RBRACKET, string(token.RBRACKETRUNE), nil
		}
	case '.':
		{
			return '.', ".", nil
		}
	default:
		inspect := string(ch)
		if strings.TrimSpace(lit) != "" {
			inspect += " (`" + lit + "`)"
		}

		return token.Token(ch), string(ch), errors.New("Unrecognized token " + inspect + " in expression")
	}
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return token.WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	//
	//// If the string matches a keyword then return that keyword.
	//switch strings.ToUpper(buf.String()) {
	//case "SELECT":
	//	return SELECT, buf.String()
	//case "FROM":
	//	return FROM, buf.String()
	//}

	// Otherwise return as a regular identifier.
	return token.IDENT, buf.String()
}

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '.') || (ch == '/')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }
