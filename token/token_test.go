package token

import (
	"fmt"
	"log"
	"testing"
)

import (
	lex "github.com/timtadh/lexmachine"
)

func Test(t *testing.T) {
	s, err := Lexer.Scanner([]byte(`[ SUM [ DIF A.TXT B.TXT C.TXT][ INT B.TXT C.TXT ]]`))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Type    | Lexeme     | Position")
	fmt.Println("--------+------------+------------")
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n", Tokens[token.Type], string(token.Lexeme), token.StartLine, token.StartColumn, token.EndLine, token.EndColumn)
	}

}