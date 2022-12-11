package day10

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

var clock int = 0
var x int = 1
var history []int

func run(input *[]string) {
	history = make([]int, len(*input)*2)
	history[0] = 1

	for _, text := range *input {
		tokens := strings.Split(text, " ")
		op := tokens[0]
		switch op {
		case "noop":
			{
				clock++
			}
		case "addx":
			{
				val, _ := strconv.Atoi(tokens[1])
				history[clock+1] = history[clock]
				clock += 2
				x += val
			}
		}
		history[clock] = x
	}
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day10/test.txt", "").Lines
	} else {
		input = util.ReadInput("day10/day10.txt", "").Lines
	}

	run(&input)
}
