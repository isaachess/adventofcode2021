package common

import (
	"bufio"
	"os"
)

func FileData(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var values []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	return values, scanner.Err()
}
