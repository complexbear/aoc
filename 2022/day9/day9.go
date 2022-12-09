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

type VisitedPositions map[Position]int

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Print(positions []Position) {
	// Find max/min coords
	var xMax int
	var yMax int
	var xMin int
	var yMin int
	for _, p := range positions {
		if p.x > xMax {
			xMax = p.x
		}
		if p.x < xMin {
			xMin = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
		if p.y < yMin {
			yMin = p.y
		}
	}
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			z := Position{x, y}
			t := "."
			for i, p := range positions {
				if z == p {
					t = strconv.Itoa(i)
					break
				}
			}
			fmt.Printf(" %s ", t)
		}
		fmt.Println("")
	}
}

func readMoves(input *[][]string) []Move {
	moves := make([]Move, len(*input))
	for i, tokens := range *input {
		steps, _ := strconv.Atoi(tokens[1])
		moves[i] = Move{direction: tokens[0], steps: steps}
	}
	return moves
}

func moveHead(m Move, head *Position) {
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
}

func moveRope(head, tail *Position) *Position {
	// Does the tail have to move?
	if Abs(head.x-tail.x) > 1 || Abs(head.y-tail.y) > 1 {
		// If gap is across row AND column, then the tail moves diagonally
		if (head.x-tail.x) != 0 && (head.y-tail.y) != 0 {
			if head.x-tail.x > 0 {
				tail.x++
			} else {
				tail.x--
			}
			if head.y-tail.y > 0 {
				tail.y++
			} else {
				tail.y--
			}
		} else {
			// normal horiz or vertical move
			if head.x-tail.x == 0 {
				if head.y-tail.y > 0 {
					tail.y++
				} else {
					tail.y--
				}
			} else {
				if head.x-tail.x > 0 {
					tail.x++
				} else {
					tail.x--
				}
			}
		}
		return tail
	}
	return nil
}

func whipRope(moves []Move, knots []Position) int {
	tailVisited := VisitedPositions{knots[0]: 1} // We know the tail starts at same pos as head
	for _, m := range moves {
		// fmt.Println(m)
		for s := 0; s < m.steps; s++ {
			moveHead(m, &knots[0])
			for i := 0; i < len(knots)-1; i++ {
				tailMoved := moveRope(&knots[i], &knots[i+1])
				if tailMoved == nil {
					break
				}
				// Track true end of the tail
				if i+1 == len(knots)-1 {
					tailVisited[*tailMoved] = 1
				}
			}
		}
	}
	// fmt.Println(tailVisited)
	return len(tailVisited)
}

func Main(testmode bool) {
	var input1 [][]string
	var input2 [][]string
	if testmode {
		input1 = [][]string{
			{"R", "4"},
			{"U", "4"},
			{"L", "3"},
			{"D", "1"},
			{"R", "4"},
			{"D", "1"},
			{"L", "5"},
			{"R", "2"},
		}
		input2 = [][]string{
			{"R", "5"},
			{"U", "8"},
			{"L", "8"},
			{"D", "3"},
			{"R", "17"},
			{"D", "10"},
			{"L", "25"},
			{"U", "20"},
		}
	} else {
		input1 = util.ReadInput("day9/day9.txt", " ").Tokens
		input2 = input1
	}

	// Part 1
	moves := readMoves(&input1)
	knots := make([]Position, 2)
	tailVisited := whipRope(moves, knots)
	// Print(knots)
	fmt.Printf("Tail visited %d positions\n", tailVisited)

	// Part 2
	moves = readMoves(&input2)
	knots = make([]Position, 10)
	tailVisited = whipRope(moves, knots)
	// Print(knots)
	fmt.Printf("Tail visited %d positions\n", tailVisited)
}
