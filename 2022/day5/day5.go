package day5

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Stack []string
type State []Stack
type Move struct {
	from  int
	to    int
	count int
}

func (s State) Print() {
	// Print horizontally for convenience
	for i := 0; i < len(s); i++ {
		stack := strings.Join(s[i], " ")
		fmt.Printf("%d: %s\n", i+1, stack)
	}
}

func parseStackLine(text *string, state *State) {
	// allocate items in this text into the correct stack in the state
	// 4 chars occupy each stack, and could be all whitespace
	s := *state
	sidx := 0
	t := *text
	for i := 0; i < len(t); i += 4 {
		stack := t[i : i+3]
		if strings.Contains(stack, "[") {
			s[sidx] = append(s[sidx], string(stack[1]))
		}
		sidx++
	}
}

func readState(input []string) State {
	/*
		Example
			[D]
		[N] [C]
		[Z] [M] [P]
		 1   2   3
	*/

	// max stack height
	maxstack := len(input) - 1

	// read last line to get stack count
	lastline := strings.Fields(input[len(input)-1])

	state := make(State, len(lastline))
	// go in reverse order so stacks are FILO
	for i := maxstack - 1; i >= 0; i-- {
		// search line for stack items
		parseStackLine(&input[i], &state)
	}
	return state
}

func readMoves(input [][]string) []Move {
	moves := make([]Move, len(input))
	for i, text := range input {
		from, _ := strconv.Atoi(text[3])
		to, _ := strconv.Atoi(text[5])
		count, _ := strconv.Atoi(text[1])
		moves[i] = Move{
			from:  from,
			to:    to,
			count: count,
		}
	}
	return moves
}

func PerformMove(move *Move, state *State) {
	m := *move
	s := *state
	for i := 0; i < m.count; i++ {
		to := m.to - 1 // adjust for zero indexing
		from := m.from - 1
		items := &(s[from])                            // source item stack
		s[to] = append(s[to], (*items)[len(*items)-1]) // add item
		s[from] = s[from][:len(s[from])-1]             // remove item
	}
}

func Main(testmode bool) {
	var inputstate []string
	var inputmoves [][]string

	if testmode {
		inputstate = []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
		}
		inputmoves = [][]string{
			{"move", "1", "from", "2", "to", "1"},
			{"move", "3", "from", "1", "to", "3"},
			{"move", "2", "from", "2", "to", "1"},
			{"move", "1", "from", "1", "to", "2"},
		}

	} else {
		inputstate = util.ReadInput("day5/day5_init.txt", "").Lines
		inputmoves = util.ReadInput("day5/day5_moves.txt", " ").Tokens
	}

	state := readState(inputstate)
	moves := readMoves(inputmoves)

	state.Print()
	for _, m := range moves {
		fmt.Printf("------------- %+v\n", m)
		PerformMove(&m, &state)
		state.Print()
	}

	topcrates := make([]byte, len(state))
	for i, s := range state {
		topcrates[i] = byte(s[len(s)-1][0])
	}
	fmt.Printf("Top crates on stacks: %s\n", topcrates)
}
