package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Lines  []string
	Tokens [][]string
}

func ReadInput(filename string, asTokens bool) Input {
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
		if asTokens {
			input.Tokens = append(input.Tokens, strings.Split(text, " "))
		} else {
			input.Lines = append(input.Lines, text)
		}
	}

	return input
}
