package scanner

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"

	"github.com/NickTaporuk/2021ai_test/operators"
	"github.com/NickTaporuk/2021ai_test/stack"
	"github.com/NickTaporuk/2021ai_test/token"
	"github.com/NickTaporuk/2021ai_test/utils"
	"github.com/fatih/set"
)

// Scanner is the interface that wraps the basic Scan method.
//
// Implementations must not retain p.
type Scanner interface {
	Scan() (data set.Interface, err error)
}

// Scanner represents a lexical scanner.
type LexicalScanner struct {
	r *bufio.Reader
}

// NewLexicalScanner returns a new instance of LexicalScanner.
func NewLexicalScanner(r io.Reader) *LexicalScanner {
	return &LexicalScanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *LexicalScanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

// unread places the previously read rune back on the reader.
func (s *LexicalScanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *LexicalScanner) Scan() (res set.Interface, err error) {
	// validateBrackets
	var validateBrackets int
	//sets := stack.NewSetStack()
	sets := make(map[int][]set.Interface)
	ops := stack.NewStringStack()
	//
	var prev token.Token
loop:
	for {

		// Read the next rune.
		ch := s.read()

		if isWhitespace(ch) {
			s.unread()
			s.scanWhitespace()
			continue
			//return tok, lit, nil
		} else if isLetter(ch) {
			s.unread()
			tok, lit := s.scanIdent()

			if isFile(lit) {
				data, err := utils.ReadDataFromFile(lit)
				if err != nil {
					return nil, err
				}

				if data != nil {
					sets[validateBrackets] = append(sets[validateBrackets], data)
				}
				prev = tok
				continue
			}

			if operators.IsOperator(lit) {
				switch lit {
				case operators.DifferentOperatorLex:
					ops.Push(operators.DifferentOperatorLex)
					prev = tok
					continue
				case operators.UnionOperatorLex:
					ops.Push(operators.UnionOperatorLex)
					prev = tok
					continue
				case operators.IntersectionOperatorLex:
					ops.Push(operators.IntersectionOperatorLex)
					prev = tok
					continue
				default:
					inspect := string(ch)
					if strings.TrimSpace(lit) != "" {
						inspect += " (`" + lit + "`)"
					}

					return nil, errors.New("Unrecognized token " + inspect + " in expression")
				}
			}

		}

		// Otherwise read the individual character.
		switch ch {
		case eof:
			break loop
		case token.LBRACKETRUNE:
			{
				validateBrackets += 1
				prev = token.LBRACKETRUNE
				continue
			}
		case token.RBRACKETRUNE:
			{

				if prev == 1 {
					print("prev = 1")
				}
				op, err := ops.Pop()
				if err != nil {
					return nil, err
				}

				d := sets[validateBrackets]

				res, err = runOp(op, d)
				//delete(sets, validateBrackets)
				validateBrackets -= 1

				if len(sets) > 0 {
					sets[validateBrackets] = append(sets[validateBrackets], res)
				}

				prev = token.RBRACKETRUNE

				continue
			}
		default:
			inspect := string(ch)
			//if strings.TrimSpace(lit) != "" {
			//	inspect += " (`" + lit + "`)"
			//}

			return nil, errors.New("Unrecognized token " + inspect + " in expression")
		}
	}

	return res, nil
}

func runOp(opName string, data []set.Interface) (set.Interface, error) {
	op := operators.FindOperatorFromString(opName)
	if op == nil {
		return nil, errors.New("Either unmatched paren or unrecognized operator")
	}

	res := op.Operation(data...)

	return res, nil
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *LexicalScanner) scanWhitespace() (tok token.Token, lit string) {
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
func (s *LexicalScanner) scanIdent() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return token.IDENT, buf.String()
}

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isLetter
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '.') || (ch == '/')
}

// isFile
func isFile(lit string) bool {
	ext := filepath.Ext(lit)

	return ext != ""
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }
