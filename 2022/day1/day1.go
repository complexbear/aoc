package day1
// https://adventofcode.com/2022/day/1

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
)

type Elf struct {
	calories []int
	total    int
}

func NewElf(calories []int) Elf {
	total := 0
	for _, t := range calories {
		total += t
	}
	return Elf{calories: calories, total: total}
}

func SumElves(elves []Elf) int {
	total := 0
	for _, e := range elves {
		total += e.total
	}
	return total
}

func readInput(filename string) []Elf {

	input := util.ReadInput(filename, "")

	elves := make([]Elf, 0)
	calories := make([]int, 0)
	for _, text := range input.Lines {
		if text == "" {
			// Create new Elf
			elf := NewElf(calories)
			elves = append(elves, elf)
			calories = make([]int, 0) // reset
			//fmt.Printf("created new Elf. Calories: %d\n", elf.Total())
		}

		// Add calorie to current count
		calorie, _ := strconv.Atoi(text)
		calories = append(calories, calorie)
	}
	// Add final Elf if calories still in the buffer
	if len(calories) != 0 {
		elf := NewElf(calories)
		elves = append(elves, elf)
	}

	return elves
}

func Main(testmode bool) {
	elves := []Elf{}
	if testmode {
		elves = readInput("testinput.txt")
	} else {
		elves = readInput("day1/day1.txt")

	}

	// Day 1
	// Max calories carried
	max_calorie := 0
	for _, e := range elves {
		total := e.total
		if max_calorie < total {
			max_calorie = total
		}
	}
	fmt.Printf("Max calorie load is: %d\n", max_calorie)

	// Day2
	// Total calories from the top 3 elves
	topThree := make([]Elf, 3)
	for _, e := range elves {
		// TopThree elves will be maintained in sorted order, lowest to highest
		// So we only have to ever check the 1st topThree elf
		if topThree[0].total < e.total {
			topThree[0] = e
			sort.Slice(topThree, func(m, n int) bool {
				// Sort in ascending order
				return topThree[m].total < topThree[n].total
			})
		}
	}
	fmt.Printf("Top 3 elves are: %+v\n", topThree)
	fmt.Printf("Top 3 elves total: %d\n", SumElves(topThree))
}
