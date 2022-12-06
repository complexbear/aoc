package day6

import (
	"aoc/util"
	"fmt"
)

func startOfIndex(text *string, bufSize int) int {
	// return index of start of message
	// marked by bufSize different characters
	fmt.Println(*text)
	startIdx := bufSize

	tracker := make(map[byte]int)
	for i := 0; i < bufSize; i++ {
		tracker[(*text)[i]] += 1
	}
	for i := 0; i < len(*text)-bufSize; i++ {
		isUnique := true
		for _, t := range tracker {
			if t > 1 {
				isUnique = false
				break
			}
		}
		if isUnique {
			return startIdx
		}

		tracker[(*text)[startIdx]] += 1
		tracker[(*text)[i]] -= 1
		startIdx++
	}

	return len(*text)
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = []string{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			"nppdvjthqldpwncqszvftbrmjlhg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		}

	} else {
		input = util.ReadInput("day6/day6.txt", "").Lines
	}

	for _, text := range input {
		start := startOfIndex(&text, 4)
		fmt.Printf("Packet start index: %d\n", start)
		start = startOfIndex(&text, 14)
		fmt.Printf("Message start index: %d\n", start)
		fmt.Println("-------------------------")
	}
}
