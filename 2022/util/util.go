package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x ...int) int {
	min := x[0]
	for i := range x {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

func Max(x ...int) int {
	max := x[0]
	for i := range x {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func StrToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func StringsToInts(items []string) []int {
	ints := make([]int, len(items))
	for i, s := range items {
		ints[i] = StrToInt(s)
	}
	return ints
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
