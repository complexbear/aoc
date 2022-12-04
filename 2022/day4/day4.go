package day4

// https://adventofcode.com/2022/day/4

import (
	// "aoc/util"
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Assignment struct {
	start int
	end   int
}

func (a Assignment) Size() int {
	return a.end - a.start + 1
}

type SharedAssignments []Assignment

func parseInput(tokens [][]string) []SharedAssignments {
	elfpairs := make([]SharedAssignments, len(tokens))
	for i, tokens := range tokens {
		// tokens in the form {"3-4", "8-10"}
		for _, token := range tokens {
			// token in the form "3-7"
			sections := []int{}
			for _, t := range strings.Split(token, "-") {
				idx, _ := strconv.Atoi(t)
				sections = append(sections, idx)
			}
			assignment := Assignment{start: sections[0], end: sections[1]}
			elfpairs[i] = append(elfpairs[i], assignment)
		}
	}
	return elfpairs
}

type EvalFunc func(*Assignment, *Assignment) bool

func hasFullyCoveredSection(elf1, elf2 *Assignment) bool {
	var firstelf, secondelf *Assignment

	// decide which elf has min start, or if they start at the same point
	// then select the one with the largest range
	if elf1.start == elf2.start {
		if elf1.Size() > elf2.Size() {
			firstelf = elf1
			secondelf = elf2
		} else {
			firstelf = elf2
			secondelf = elf1
		}
	} else if elf1.start < elf2.start {
		firstelf = elf1
		secondelf = elf2
	} else {
		firstelf = elf2
		secondelf = elf1
	}

	// the other elf's start and end must fit within the end of the first elf
	if secondelf.start <= firstelf.end && secondelf.end <= firstelf.end {
		return true
	}
	return false
}

func hasPartialCoveredSection(elf1, elf2 *Assignment) bool {
	if elf1.start >= elf2.start && elf1.start <= elf2.end {
		return true
	}
	if elf2.start >= elf1.start && elf2.start <= elf1.end {
		return true
	}
	return false
}

func countCoveredSections(elfpairs []SharedAssignments, evalfunc EvalFunc) int {
	count := 0
	for _, elves := range elfpairs {
		fmt.Printf("Elves %v", elves)
		if evalfunc(&elves[0], &elves[1]) {
			count += 1
			fmt.Printf("  overlap")
		}
		fmt.Printf("\n")
	}
	return count
}

func Main(testmode bool) {
	var tokens [][]string
	if testmode {
		tokens = [][]string{
			{"2-4", "6-8"},
			{"2-3", "4-5"},
			{"5-7", "7-9"},
			{"2-8", "3-7"},
			{"6-6", "4-6"},
			{"2-6", "4-8"},
		}
	} else {
		tokens = util.ReadInput("day4/day4.txt", ",").Tokens
	}

	elfpairs := parseInput(tokens)

	value := countCoveredSections(elfpairs, hasFullyCoveredSection)
	fmt.Printf("Elf pairs with fully covered sections: %d\n", value)

	value = countCoveredSections(elfpairs, hasPartialCoveredSection)
	fmt.Printf("Elf pairs with partial covered sections: %d\n", value)
}
