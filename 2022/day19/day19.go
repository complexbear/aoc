package day19

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mitchellh/hashstructure/v2"
)

var DURATION = 0

type Cost struct {
	Name  int
	Value int
}

type Robot struct {
	Name  int
	Costs []Cost
}

type Blueprint struct {
	Robots [4]Robot
}

type State struct {
	Robots [4]int
	Stock  [4]int
	Time   int
}

func (s *State) hash() uint64 {
	hash, err := hashstructure.Hash(*s, hashstructure.FormatV2, nil)
	if err != nil {
		panic("doh")
	}
	return hash
}

const ORE = 0
const CLAY = 1
const OBSIDIAN = 2
const GEODE = 3

var NAMES = [4]int{3, 2, 1, 0}
var NAMES_STR = [4]string{"ORE", "CLAY", "OBSIDIAN", "GEODE"}

func nameIdx(name string) int {
	switch name {
	case "ore":
		return ORE
	case "clay":
		return CLAY
	case "obsidian":
		return OBSIDIAN
	case "geode":
		return GEODE
	}
	return -1
}

func createRobot(input string) Robot {
	pattern1 := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+)$`)
	pattern2 := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+) and (\d+) (\w+)$`)

	tokens := pattern1.FindStringSubmatch(input)
	if tokens == nil {
		tokens = pattern2.FindStringSubmatch(input)
	}
	r := Robot{}
	r.Name = nameIdx(tokens[1])
	tokens = tokens[2:]
	r.Costs = make([]Cost, len(tokens)/2)
	for i, _ := range r.Costs {
		r.Costs[i].Name = nameIdx(tokens[i*2+1])
		r.Costs[i].Value = util.StrToInt(tokens[i*2])
	}
	return r
}

func parseInput(text []string) []Blueprint {
	blueprints := make([]Blueprint, len(text))
	for i, t := range text {
		t = strings.Split(t, ":")[1]
		items := strings.Split(t, ".")

		b := &blueprints[i]
		b.Robots[ORE] = createRobot(items[ORE])
		b.Robots[CLAY] = createRobot(items[CLAY])
		b.Robots[OBSIDIAN] = createRobot(items[OBSIDIAN])
		b.Robots[GEODE] = createRobot(items[GEODE])
		fmt.Println(*b)
	}
	fmt.Println("Blueprints:", len(blueprints))
	return blueprints
}

var STATES = []State{}
var STATES_CACHE = map[uint64]struct{}{}
var MAX_GEODES int

func buildRobot(state *State, name int, costs []Cost) State {
	newState := *state
	newState.Robots[name]++
	for _, c := range costs {
		newState.Stock[c.Name] -= c.Value
	}
	return newState
}

func harvest(state *State) {
	state.Time++
	for i, s := range state.Robots {
		state.Stock[i] += s
	}
}

func addState(s *State) {
	hash := s.hash()
	_, ok := STATES_CACHE[hash]
	if !ok {
		STATES = append(STATES, *s)
		STATES_CACHE[hash] = struct{}{}
	}
}

func calcLimits(b *Blueprint) [4]int {
	limits := [4]int{}
	for _, r := range b.Robots {
		for _, c := range r.Costs {
			limits[c.Name] = util.Max(c.Value, limits[c.Name])
		}
	}
	return limits
}

func evolve(b *Blueprint, state *State, limits [4]int) {
	// using this state what is max geodes that could be produced in remaining time
	timeLeft := DURATION - state.Time
	geodes := state.Stock[GEODE] + timeLeft*state.Robots[GEODE]
	if geodes > MAX_GEODES {
		// fmt.Println(*state, geodes)
		MAX_GEODES = geodes
	}

	if state.Time >= DURATION {
		return
	}

	// what if optimisitic geode building cannot beat our best obvservation
	extraGeodes := ((timeLeft * timeLeft) + timeLeft) / 2
	if extraGeodes+state.Stock[GEODE] < MAX_GEODES-1 {
		return
	}

	// the do nothing case
	newState := *state
	harvest(&newState)
	addState(&newState)

	// try to find better results by building different robots
	for _, robot := range NAMES {
		// time taken to build this robot
		costReq := b.Robots[robot].Costs
		canBuild := true
		newRobotState := newState

		for _, c := range costReq {
			if state.Robots[c.Name] == 0 {
				canBuild = false
				break
			}
			if state.Stock[c.Name]-c.Value < 0 {
				canBuild = false
				break
			}
		}

		// if state.Robots[robot] > limits[robot] {
		// 	// no point building more of a robot if we have enough of them for most expensive robot
		// 	canBuild = false
		// }

		if canBuild {
			// build robot and add new
			r := buildRobot(&newRobotState, robot, costReq)
			addState(&r)

			// if we can build a geode robot, no need to build anything else
			if robot == GEODE {
				break
			}
		}
	}
}

func runBlueprint(blueprint Blueprint) int {
	MAX_GEODES = 0
	STATES_CACHE = map[uint64]struct{}{}
	STATES = make([]State, 1)
	STATES[0] = State{Robots: [4]int{1, 0, 0, 0}, Time: 0}

	limits := calcLimits(&blueprint)
	// fmt.Println("Limits:", limits)

	for len(STATES) > 0 {
		latestState := STATES[len(STATES)-1]
		STATES = STATES[:len(STATES)-1]
		evolve(&blueprint, &latestState, limits)
	}
	return MAX_GEODES
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day19/test.txt", "").Lines
	} else {
		input = util.ReadInput("day19/day19.txt", "").Lines
	}

	blueprints := parseInput(input)
	result := 0

	// part 1
	DURATION = 24
	for i, b := range blueprints {
		t := time.Now()
		g := runBlueprint(b)
		r := (i + 1) * g
		result += r
		fmt.Printf(
			"Blueprint %d geodes: %d, score: %d, perf: %s\n",
			i,
			g,
			r,
			time.Since(t),
		)
	}
	fmt.Println("Part 1 result:", result)

	// part 2
	DURATION = 32
	result = 1
	for i, b := range blueprints {
		t := time.Now()
		g := runBlueprint(b)
		result *= g
		fmt.Printf(
			"Blueprint %d geodes: %d, perf: %s\n",
			i,
			g,
			time.Since(t),
		)
	}
	fmt.Println("Part 2 result:", result)
}
