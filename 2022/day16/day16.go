package day16

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

type Valve struct {
	location    string
	rate        int
	connections []string
	open        bool
	visited     bool
}

var valves map[string]Valve

func parseInput(input []string) {
	pattern := regexp.MustCompile(`Valve (\w+) has flow rate=(-?\d+); tunnels? leads? to valves? (.*)$`)
	valves = make(map[string]Valve, 0)
	for _, text := range input {
		groups := pattern.FindStringSubmatch(text)
		location := groups[1]
		valves[location] = Valve{
			location:    location,
			rate:        util.StrToInt(groups[2]),
			connections: strings.Split(strings.ReplaceAll(groups[3], " ", ""), ","),
			open:        false,
		}
	}
}

func search(v Valve, time *int, totalFlow *int) {
	v.visited = true
	fmt.Printf("visiting %s\ttime:%d\tflow:%d\t", v.location, *time, *totalFlow)
	if *time <= 0 {
		fmt.Println("time up")
		return
	}
	if !v.open && v.rate != 0 {
		fmt.Println("opening valve")
		v.open = true
		*time--
		flow := v.rate * (*time)
		*totalFlow += flow
	}
	fmt.Println()

	// which valve to visit next, avoid visited valves
	for _, c := range v.connections {
		nextValve := valves[c]
		if nextValve.visited {
			continue
		}
		if nextValve.rate > 0 && nextValve.open == false {
			// go there
			*time--
			search(nextValve, time, totalFlow)
		}
	}

	return
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day16/test.txt", "").Lines
	}

	// minutes := 30
	parseInput(input)
	for _, v := range valves {
		fmt.Printf("%+v\n", v)
	}

	// for each connection move to next highest rate value
	totalFlow := 0
	time := 30
	search(valves["AA"], &time, &totalFlow)
	fmt.Printf("Total flow: %d\n", totalFlow)
}
