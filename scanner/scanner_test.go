package scanner

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"github.com/NickTaporuk/2021ai_test/set"
)

func TestLexicalScanner_Scan(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}

	// positive test
	positiveStr := `[ SUM [ DIF ../testdata/a.txt ../testdata/b.txt ../testdata/c.txt ] [ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	positiveRedear := bufio.NewReader(strings.NewReader(positiveStr))
	positiveSet := set.New()
	positiveSet.Add(1, 3, 4)

	// positive test
	positiveStr1 := `[ SUM [ DIF ../testdata/a.txt [ SUM ../testdata/b.txt [ INT ../testdata/c.txt ../testdata/b.txt ] ] ]
[ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	positiveRedear1 := bufio.NewReader(strings.NewReader(positiveStr1))
	positiveSet1 := set.New()
	positiveSet1.Add(1, 3, 4)

	// negative tests
	negativeStr := `[ SUM [ DIF DIF ../testdata/a.txt ../testdata/b.txt ../testdata/c.txt ] [ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	negativeRedear := bufio.NewReader(strings.NewReader(negativeStr))

	negativeStr1 := `[ SUM [ DIF ../testdata/a.txt ] [ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	negativeRedear1 := bufio.NewReader(strings.NewReader(negativeStr1))

	negativeStr2 := `[[ SUM [ DIF ../testdata/a.txt ] [ INT ../testdata/b.txt ../testdata/c.txt ] ]`
	negativeRedear2 := bufio.NewReader(strings.NewReader(negativeStr2))

	tests := []struct {
		name    string
		fields  fields
		wantRes set.Interface
		wantErr bool
	}{
		{
			name: "positive test",
			fields: fields{
				r: positiveRedear,
			},
			wantRes: positiveSet,
			wantErr: false,
		},
		{
			name: "positive test 2;",
			fields: fields{
				r: positiveRedear1,
			},
			wantRes: positiveSet1,
			wantErr: false,
		},
		{
			name: "negative test 1; double operator",
			fields: fields{
				r: negativeRedear,
			},
			wantErr: true,
		},
		{
			name: "negative test 2; one file inside some set",
			fields: fields{
				r: negativeRedear1,
			},
			wantErr: true,
		},
		{
			name: "negative test 3; left brackets is more than needed ",
			fields: fields{
				r: negativeRedear2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LexicalScanner{
				// nolint
				r: tt.fields.r,
			}
			gotRes, err := s.Scan()
			// nolint
			if (err != nil) != tt.wantErr {
				// nolint
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// nolint
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Scan() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
