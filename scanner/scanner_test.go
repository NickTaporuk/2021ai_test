package scanner

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	s := `[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]`

	r := strings.NewReader(s)


	sc := LexicalScanner{r:bufio.NewReader(r)}

	data, err := sc.Scan()
	fmt.Println(data.List(), err)
}