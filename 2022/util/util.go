package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Input struct {
	Lines  []string
	Tokens [][]string
}

func ReadInput(filename string, delim string) Input {
	input := Input{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return input
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if delim != "" {
			input.Tokens = append(input.Tokens, strings.Split(text, delim))
		} else {
			input.Lines = append(input.Lines, text)
		}
	}

	return input
}
