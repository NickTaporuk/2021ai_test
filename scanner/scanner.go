package scanner

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/NickTaporuk/2021ai_test/operators"
	"github.com/NickTaporuk/2021ai_test/set"
	"github.com/NickTaporuk/2021ai_test/stack"
	"github.com/NickTaporuk/2021ai_test/token"
	"github.com/NickTaporuk/2021ai_test/utils"
)

var (
	// ErrorSyntaxShouldBeFileOrOperator is error if some token before doest't have a token of a file
	// or a token of a operator
	ErrorSyntaxShouldBeFileOrOperator = errors.New("syntax error with parameters. Before file should be some operator or some file")

	// ErrorSyntaxShouldBeLeftBracketBeforeOperator is error if some token before doest't have a token of
	// left bracket
	ErrorSyntaxShouldBeLeftBracketBeforeOperator = errors.New("syntax error. You should write [ before operator SUM, INT, DIF")

	// ErrorSyntaxNoEnoughBrackets use if the app hasn't enough brackets
	ErrorSyntaxNoEnoughBrackets = errors.New("syntax error. The string doesn't have enought brackets")

	// ErrorTextUnrecognizedToken use if the app has a unrecognized token
	ErrorTextUnrecognizedToken = "Unrecognized token %s in expression"

	// ErrorTextUnrecognizedOperator use if the app has a unrecognized operator
	ErrorTextUnrecognizedOperator = errors.New("either unmatched parent or unrecognized operator")
)

// Scanner is the interface that wraps the basic Scan method.
// Implementations must not retain p.
type Scanner interface {
	Scan() (set set.Interface, err error)
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

// nolint
// Scan returns the next token and literal value.
func (s *LexicalScanner) Scan() (res set.Interface, err error) {
	// validateBrackets is a checker of brackets exists
	var validateBrackets int
	// sets is map of sets
	sets := make(map[int][]set.Interface)
	ops := stack.NewStringStack()
	// prev is  a last token which used
	var prev token.Token
loop:
	for {

		// Read the next rune.
		ch := s.read()

		if isWhitespace(ch) {
			s.unread()
			s.scanWhitespace()
			continue
		} else if isLetter(ch) {
			s.unread()
			tok, lit := s.scanIdent()

			if isFile(lit) {

				if prev != token.SUM && prev != token.INT && prev != token.DIF && prev != token.FILE {
					return nil, ErrorSyntaxShouldBeFileOrOperator
				}

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
				if prev != token.LBRACKETRUNE {
					return nil, ErrorSyntaxShouldBeLeftBracketBeforeOperator
				}
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

					errStr := fmt.Sprintf(ErrorTextUnrecognizedToken, inspect)

					return nil, errors.New(errStr)
				}
			}

		}

		// Otherwise read the individual character.
		switch ch {
		case eof:
			break loop
		case token.LBRACKETRUNE:
			validateBrackets++
			prev = token.LBRACKETRUNE
			continue
		case token.RBRACKETRUNE:
			{
				op, err := ops.Pop()
				if err != nil {
					return nil, err
				}

				d := sets[validateBrackets]

				res, err = runOp(op, d)
				if err != nil {
					return nil, err
				}

				sets[validateBrackets] = nil
				validateBrackets--

				if len(sets) > 0 {
					sets[validateBrackets] = append(sets[validateBrackets], res)
				}

				prev = token.RBRACKETRUNE

				continue
			}
		default:
			inspect := string(ch)

			errStr := fmt.Sprintf(ErrorTextUnrecognizedToken, inspect)

			return nil, errors.New(errStr)
		}
	}

	if validateBrackets != 0 {
		return nil, ErrorSyntaxNoEnoughBrackets
	}

	return res, nil
}

func runOp(opName string, data []set.Interface) (set.Interface, error) {
	op := operators.FindOperatorFromString(opName)
	if op == nil {
		return nil, ErrorTextUnrecognizedOperator
	}

	res, err := op.Operation(data...)

	return res, err
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *LexicalScanner) scanWhitespace() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.

	for {
		ch := s.read()

		if !isWhitespace(ch) {
			s.unread()
			break
		}

		switch ch {
		case eof:
			break
		default:
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
	tok = token.IDENT
	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		ch := s.read()

		if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		}

		switch ch {
		case eof:
			break
		default:
			_, _ = buf.WriteRune(ch)
		}
	}

	lit = buf.String()

	if isFile(lit) {
		tok = token.FILE

		return tok, lit
	}

	switch lit {
	case operators.DifferentOperatorLex:
		tok = token.DIF
	case operators.IntersectionOperatorLex:
		tok = token.INT
	case operators.UnionOperatorLex:
		tok = token.SUM
	}
	// Otherwise return as a regular identifier.
	return tok, lit
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
