package common

import (
	"bufio"
	"io"
	"os"
)

func Read(input string) ([]string, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	r4 := bufio.NewReader(f)

	var output []string

	for {
		line, _, err := r4.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		output = append(output, string(line))
	}

	return output, nil
}
