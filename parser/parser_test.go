package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	s := `[ SUM [ INT  a.txt b.txt c.txt ][ DIF [SUM a.txt b.txt] ]
[DIF a.txt c.txt] ]`
	//s := `a.txt`
	r := strings.NewReader(s)
	p := NewParser(r)

	d, err := p.Parse()

	assert.NoError(t, err)
	assert.NotNil(t, d)
}
