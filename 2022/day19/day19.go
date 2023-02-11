package day19

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

const DURATION int = 10

type Cost struct {
	name  string
	value int
}

type Robot struct {
	name  string
	costs []Cost
}

type Blueprint struct {
	robots map[string]Robot
}

func createRobot(name string, input string) Robot {
	pattern1 := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+)$`)
	pattern2 := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+) and (\d+) (\w+)$`)

	tokens := pattern1.FindStringSubmatch(input)
	if tokens == nil {
		tokens = pattern2.FindStringSubmatch(input)
	}
	r := Robot{}
	r.name = tokens[1]
	tokens = tokens[2:]
	r.costs = make([]Cost, len(tokens)/2)
	for i, _ := range r.costs {
		r.costs[i].name = tokens[i*2+1]
		r.costs[i].value = util.StrToInt(tokens[i*2])
	}
	return r
}

func parseInput(text []string) []Blueprint {
	blueprints := make([]Blueprint, len(text))
	for i, t := range text {
		t = strings.Split(t, ":")[1]
		items := strings.Split(t, ".")

		b := &blueprints[i]
		b.robots = make(map[string]Robot, 4)
		b.robots["ore"] = createRobot("ore", items[0])
		b.robots["clay"] = createRobot("clay", items[1])
		b.robots["obsidian"] = createRobot("obsidian", items[2])
		b.robots["geode"] = createRobot("geode", items[3])
		fmt.Println(*b)
	}
	return blueprints
}

var ROBOTS = map[string]int{}
var STOCK = map[string]int{}
var FACTORY = map[string]int{}
var PRIORITY = [4]string{"geode", "obsidian", "clay", "ore"}

func timeStep(b *Blueprint) {
	fmt.Println("ROOBTS", ROBOTS)
	fmt.Println("STOCK", STOCK)
	fmt.Println("FACTORY", FACTORY)

	// has factory finished any robots?
	for name, _ := range FACTORY {
		for FACTORY[name] > 0 {
			fmt.Printf("finished %s robot\n", name)
			FACTORY[name]--
			ROBOTS[name]++
		}
	}

	// start building new robots, update stock
	for _, name := range PRIORITY {
		availStock := true
		for availStock {
			r := b.robots[name]
			for _, c := range r.costs {
				availStock = STOCK[c.name] >= c.value
			}
			if availStock {
				// build new robot
				fmt.Printf("building %s robot\n", r.name)
				FACTORY[r.name] += 1
				for _, c := range r.costs {
					STOCK[c.name] -= c.value
				}
			}
		}
	}

	// collecting robots add to stock
	for name, number := range ROBOTS {
		STOCK[name] += number
	}
}

func runBlueprint(blueprint Blueprint) int {
	ROBOTS = map[string]int{
		"ore":      1,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	STOCK = map[string]int{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	FACTORY = map[string]int{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	for i := 0; i < DURATION; i++ {
		fmt.Printf("Time: %d\n", i)
		timeStep(&blueprint)
	}
	return STOCK["geodes"]
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day19/test.txt", "").Lines
	} else {
		input = util.ReadInput("day19/day19.txt", "").Lines
	}

	blueprints := parseInput(input)
	geodes := make([]int, len(blueprints))

	for i, b := range blueprints[:1] {
		geodes[i] = runBlueprint(b)
		fmt.Printf("Blueprint %d geodes: %d\n", i, geodes[i])
	}
}
