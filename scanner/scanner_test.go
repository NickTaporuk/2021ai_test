package scanner

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	//s := `[ SUM ../testdata/a.txt [ DIF ../testdata/a.txt ../testdata/b.txt ../testdata/c.txt ] [ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	s := `[ SUM [ DIF ../testdata/a.txt ../testdata/b.txt ../testdata/c.txt ] [ INT ../testdata/b.txt ../testdata/	c.txt ] ]`
	r := strings.NewReader(s)


	sc := LexicalScanner{r:bufio.NewReader(r)}

	data, err := sc.Scan()
	fmt.Println(data, err)
}