package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	s := `[ SUM [ INT  /home/users/aaa/sadd.txt bbb.txt cdd.txt ][ DIF [SUM a.txt b.txt] ]
[DIF ss.txt dd.txt] ]`
	r := strings.NewReader(s)
	p := NewParser(r)

	stmt, err := p.Parse()

	if stmt != nil && len(stmt.Fields) > 0 {
		for _, d := range stmt.Fields {
			fmt.Println(d)
		}
	}

	assert.NoError(t, err)
	assert.NotNil(t, stmt)
}
