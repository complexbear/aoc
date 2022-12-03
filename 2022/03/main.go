package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	return input
}

func Priority(item byte) int {
	// convert item to byte value
	// [a-z] => [1..26]
	// [A-Z] => [27..52]
	ref := 0
	if item > 'Z' {
		ref = int('a')
	} else {
		ref = int('A') - 26
	}
	return int(item) - ref + 1
}

func DuplicateItemInSack(text string) byte {
	// split input in halves
	size := len(text) / 2
	c1 := text[:size]
	c2 := text[size:]

	// put c1 items into a map, and then check c2 to see if any are repeated
	reference := make(map[byte]struct{})
	for _, i := range []byte(c1) {
		reference[i] = struct{}{}
	}
	for _, i := range []byte(c2) {
		_, exists := reference[i]
		if exists {
			return i
		}
	}
	return '0'
}

func DuplicateItemInSacks(sacks []string) byte {
	// use map with counter to determine which item is in all 3 sacks
	groupReference := make(map[byte]int)
	for _, s := range sacks {
		reference := make(map[byte]struct{})
		for _, i := range []byte(s) {
			_, exists := reference[i]
			if exists {
				continue
			}
			// mark this item in this sack to avoid double counting item types
			reference[i] = struct{}{}
			// group sack item type count update
			groupReference[i] += 1
		}
	}
	// fmt.Printf("Map: %v\n", groupReference)
	badge := byte('0')
	for i, count := range groupReference {
		if count == 3 {
			badge = i
			break
		}
	}
	return badge
}

func AssessSacksPart1(input []string) {

	mistakes := make([]byte, len(input))
	prioritySum := 0
	for i, text := range input {
		mistakes[i] = DuplicateItemInSack(text)
		priority := Priority(mistakes[i])
		prioritySum += priority
		fmt.Printf("Sack %d, mistake = %s, priority = %d\n", i, string(mistakes[i]), priority)
	}
	fmt.Printf("Total priority: %d\n", prioritySum)
}

func AssessSacksPart2(input []string) {
	// groups of 3 sacks
	prioritySum := 0
	for i := 0; i < len(input); i += 3 {
		groupSacks := input[i : i+3]
		// fmt.Printf("Group: %v\n", groupSacks)
		groupItem := DuplicateItemInSacks(groupSacks)
		priority := Priority(groupItem)
		prioritySum += priority
		fmt.Printf("Group item: %s, priority: %d\n", string(groupItem), priority)
	}
	fmt.Printf("Total priority: %d\n", prioritySum)
}

func main() {
	test := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	input := readInput("day3.txt")

	// AssessSacksPart1(test)
	// AssessSacksPart1(input)

	AssessSacksPart2(test)
	AssessSacksPart2(input)

}
