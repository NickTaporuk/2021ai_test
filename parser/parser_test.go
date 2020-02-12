package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	s := `[ SUM [ INT  ../testdata/a.txt ../testdata/b.txt ../testdata/c.txt ]
[ DIF ../testdata/a.txt [SUM ../testdata/a.txt ../testdata/b.txt] ]
[DIF ../testdata/a.txt ../testdata/c.txt] ]`

	r := strings.NewReader(s)
	p := NewParser(r)

	d, err := p.Parse()

	assert.NoError(t, err)
	assert.NotNil(t, d)
}
