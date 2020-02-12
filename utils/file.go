package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/set"
)

// ReadDataFromFile is helper which chech file path and read all data from the file
func ReadDataFromFile(path string) (st set.Interface, err error) {

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(abs)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	s := set.New(set.NonThreadSafe)

	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		s.Add(i)
	}

	return s, nil
}
