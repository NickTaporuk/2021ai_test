package utils

import (
	"bufio"
	"github.com/fatih/set"
	"os"
	"strconv"
)

func ReadDataFromFile(path string) (st set.Interface, err error) {
	f, err := os.Open(path)
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
