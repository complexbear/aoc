package main

import (
	"aoc/day1"
	"aoc/day10"
	"aoc/day11"
	"aoc/day12"
	"aoc/day13"
	"aoc/day14"
	"aoc/day15"
	"aoc/day17"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"aoc/day9"
	"flag"
	"fmt"
	"os"
)

type MainFunc func(bool)

func main() {
	var testmode bool
	var day int

	flag.IntVar(&day, "day", 0, "Run day n scenario")
	flag.BoolVar(&testmode, "test", false, "Activate test mode")
	flag.Parse()

	fmt.Printf("Test mode: %t\n", testmode)

	packageMap := map[int]MainFunc{
		1:  day1.Main,
		2:  day2.Main,
		3:  day3.Main,
		4:  day4.Main,
		5:  day5.Main,
		6:  day6.Main,
		7:  day7.Main,
		8:  day8.Main,
		9:  day9.Main,
		10: day10.Main,
		11: day11.Main,
		12: day12.Main,
		13: day13.Main,
		14: day14.Main,
		15: day15.Main,
		17: day17.Main,
	}

	f, exists := packageMap[day]
	if exists == false {
		fmt.Println("Invalid day scenario selected")
		os.Exit(1)
	}
	f(testmode)
}
