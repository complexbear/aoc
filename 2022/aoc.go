package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"flag"
	"fmt"
	"os"
)

func main() {
	var testmode bool
	var day int

	flag.IntVar(&day, "day", 0, "Run day n scenario")
	flag.BoolVar(&testmode, "test", false, "Activate test mode")
	flag.Parse()

	fmt.Printf("Test mode: %t\n", testmode)

	switch day {
	case 1:
		day1.Main(testmode)
	case 2:
		day2.Main(testmode)
	case 3:
		day3.Main(testmode)
	case 4:
		day4.Main(testmode)
	default:
		fmt.Println("Invalid day scenario selected")
		os.Exit(1)
	}

}
