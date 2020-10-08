package workWithFiles

import (
	"bufio"
	"os"
)

func Read(fileName string) ([]string, error) {
	var input []string
	if fileName == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return input, nil
}
