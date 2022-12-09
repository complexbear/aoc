package day9

import (
	"aoc/util"
	"fmt"

	"strconv"
)

type Move struct {
	direction string
	steps     int
}
type Position struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readMoves(input *[][]string) []Move {
	moves := make([]Move, len(*input))
	for i, tokens := range *input {
		steps, _ := strconv.Atoi(tokens[1])
		moves[i] = Move{direction: tokens[0], steps: steps}
	}
	return moves
}

func Main(testmode bool) {
	var input [][]string
	if testmode {
		input = [][]string{
			{"R", "4"},
			{"U", "4"},
			{"L", "3"},
			{"D", "1"},
			{"R", "4"},
			{"D", "1"},
			{"L", "5"},
			{"R", "2"},
		}
	} else {
		input = util.ReadInput("day9/day9.txt", " ").Tokens
	}

	moves := readMoves(&input)
	head := Position{0, 0}
	tail := Position{0, 0}

	tailVisited := map[Position]int{tail: 1}

	for _, m := range moves {
		fmt.Print(m)
		for s := 0; s < m.steps; s++ {
			initialHeadPosition := head
			switch m.direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			// Does the tail have to move?
			if Abs(head.x-tail.x) > 1 || Abs(head.y-tail.y) > 1 {
				tail = initialHeadPosition
				tailVisited[tail] = 1
			}
			fmt.Print(head, tail)
		}
		fmt.Println("--------")
	}

	fmt.Printf("Tail visited %d positions\n", len(tailVisited))
}
