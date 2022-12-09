package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
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
		1: day1.Main,
		2: day2.Main,
		3: day3.Main,
		4: day4.Main,
		5: day5.Main,
		6: day6.Main,
		7: day7.Main,
		8: day8.Main,
	}

	f, exists := packageMap[day]
	if exists == false {
		fmt.Println("Invalid day scenario selected")
		os.Exit(1)
	}
	f(testmode)
}
