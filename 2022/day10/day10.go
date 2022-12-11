package day10

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

var clock int = 1
var x int = 1
var register []int

func run(input *[]string) {
	register = make([]int, len(*input)*2)

	for _, text := range *input {
		tokens := strings.Split(text, " ")
		op := tokens[0]
		switch op {
		case "noop":
			{
				clock++
				register[clock] = x
			}
		case "addx":
			{
				val, _ := strconv.Atoi(tokens[1])
				clock++
				register[clock] = x
				clock++
				x += val
				register[clock] = x
			}
		}
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
	// for i, s := range signal {
	// 	fmt.Printf("%d:\t%d\n", i, s)
	// }

	signalPoints := []int{20, 60, 100, 140, 180, 220}
	signalSum := 0
	for _, s := range signalPoints {
		fmt.Printf("Cyl:%d\t%d\t%d\n", s, register[s], register[s]*s)
		signalSum += register[s] * s
	}
	fmt.Printf("Total: %d\n", signalSum)
}
