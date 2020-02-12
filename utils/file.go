package utils

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/NickTaporuk/2021ai_test/set"
)

var (
	// ErrorFileHasNoData is error if some file doesn't have any data inside
	ErrorFileHasNoData = errors.New("file doesn't has any data")
)

// ReadDataFromFile is helper which check file exist by path
// After that push to a set all data from the file
func ReadDataFromFile(path string) (st set.Interface, err error) {

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(abs)
	// nolint
	defer f.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	s := set.New()

	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		s.Add(i)
	}

	if s.IsEmpty() {
		return nil, ErrorFileHasNoData
	}
	return s, nil
}
