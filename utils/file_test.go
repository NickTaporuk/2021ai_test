package utils

import (
	"reflect"
	"testing"

	"github.com/NickTaporuk/2021ai_test/set"
)

func TestReadDataFromFile(t *testing.T) {

	s := set.New()
	s.Add(1, 2, 3)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantSt  set.Interface
		wantErr bool
	}{
		{
			name:   "positive test",
			args:   args{path: "../testdata/a.txt"},
			wantSt: s,
		},
		{
			name:    "negative test 1; file isn't exist",
			args:    args{path: "../testdata/aa.txt"},
			wantErr: true,
		},
		{
			name:    "negative test 2; absolute path isn't reached",
			args:    args{path: "kqdnoqnc"},
			wantErr: true,
		},
		{
			name:    "negative test 3; file is exist but is empty",
			args:    args{path: "../testdata/d.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			gotSt, err := ReadDataFromFile(tt.args.path)
			// nolint
			if (err != nil) != tt.wantErr {
				// nolint
				t.Errorf("ReadDataFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// nolint
			if !reflect.DeepEqual(gotSt, tt.wantSt) {
				t.Errorf("ReadDataFromFile() gotSt = %v, want %v", gotSt, tt.wantSt)
			}
		})
	}
}
