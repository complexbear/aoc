package day2
// https://adventofcode.com/2022/day/2

import (
	"aoc/util"
	"fmt"
)

type Option struct {
	name  string
	value int
}

var OptionValue = map[byte]Option{
	'A': {name: "rock", value: 1},
	'B': {name: "paper", value: 2},
	'C': {name: "scissors", value: 3},
}

var TargetResult = map[byte]int{
	'X': 0, // loss
	'Y': 3, // draw
	'Z': 6, // win
}

type Strategy struct {
	challenge byte
	response  byte
}

type StrategyGuide []Strategy

func (s Strategy) ScorePart1() int {
	// The response option is considered to be the choice the player should make
	// eg rock, paper or scissors

	// convert response into same symbol as challenge for convenience
	symbolMap := map[byte]byte{
		'X': 'A',
		'Y': 'B',
		'Z': 'C',
	}
	response := symbolMap[s.response]
	//fmt.Printf("%s vs %s: ", OptionValue[s.challenge].name, OptionValue[response].name)

	// score from option selection
	score := OptionValue[response].value

	// note the order rock, paper scissors represents the win order
	if s.challenge == response {
		score += 3 // draw
	} else if s.challenge == 'Z' && response == 'X' {
		// special loss case due to challenge of Z
		score += 6
	} else if s.challenge+1 == response {
		// normal case we can see if our response is next 1 option on from the challenge
		score += 6 // win
	}
	return score
}

func (s Strategy) ScorePart2() int {
	// The response option is considered to be the result of the round
	// eg loss, draw, win
	score := 0
	adjustment := 0
	response := byte('X')

	if s.response == 'Y' {
		// target is a draw
		score = 3
		response = s.challenge
	} else {
		// target is a loss, response must be 1 less than challenge
		// otherwise it should be +1
		if s.response == 'X' {
			score = 0
			adjustment = -1
		} else {
			score = 6
			adjustment = 1
		}

		// apply adjustment and handle when we fall off either end of the alphabet range
		response = byte(int(s.challenge) + adjustment)
		if response == '@' {
			response = 'C'
		}
		if response == 'D' {
			response = 'A'
		}
	}
	// fmt.Printf("target: %s,  %s vs %s\n", string(s.response), OptionValue[s.challenge].name, OptionValue[response].name)
	return score + OptionValue[response].value
}

func readInput(filename string) StrategyGuide {
	input := util.ReadInput(filename, " ")

	guide := StrategyGuide{}
	for _, s := range input.Tokens {
		guide = append(guide, Strategy{byte(s[0][0]), byte(s[1][0])})
	}
	return guide
}

func sanityCheck() {
	result := Strategy{'A', 'X'}.ScorePart1()
	fmt.Printf("score: %d\n", result)

	result = Strategy{'A', 'Z'}.ScorePart1()
	fmt.Printf("score: %d\n", result)

	result = Strategy{'C', 'Y'}.ScorePart1()
	fmt.Printf("score: %d\n", result)

	result = Strategy{'A', 'Y'}.ScorePart1()
	fmt.Printf("score: %d\n", result)

	result = Strategy{'B', 'Z'}.ScorePart1()
	fmt.Printf("score: %d\n", result)

	result = Strategy{'C', 'X'}.ScorePart1()
	fmt.Printf("score: %d\n", result)
}

func Main(testmode bool) {
	strategyGuide := StrategyGuide{}
	total := 0
	if testmode {
		sanityCheck()
		strategyGuide = StrategyGuide{
			Strategy{'A', 'Y'},
			Strategy{'B', 'X'},
			Strategy{'C', 'Z'},
		}
	} else {
		strategyGuide = readInput("day2/day2.txt")
	}

	total = 0
	for _, s := range strategyGuide {
		total += s.ScorePart1()
	}
	fmt.Printf("Part 1 total score: %d\n", total)

	total = 0
	for _, s := range strategyGuide {
		total += s.ScorePart2()
	}
	fmt.Printf("Part 2 total score: %d\n", total)
}
