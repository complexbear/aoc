package day15

import (
	"aoc/util"
	"fmt"
	"regexp"
)

type Map [][]byte

type Beacon struct {
	x int
	y int
}

type Sensor struct {
	x      int
	y      int
	beacon Beacon
}

var minX int
var maxX int
var minY int
var maxY int

func parseInput(input []string) []Sensor {
	pattern := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	sensors := make([]Sensor, len(input))
	for i, text := range input {
		coords := util.StringsToInts(pattern.FindStringSubmatch(text))
		sensor := Sensor{
			x: coords[1],
			y: coords[2],
			beacon: Beacon{
				x: coords[3],
				y: coords[4],
			},
		}
		minX = util.Min(minX, sensor.x, sensor.beacon.x)
		maxX = util.Max(maxX, sensor.x, sensor.beacon.x)
		minY = util.Min(minY, sensor.y, sensor.beacon.y)
		maxY = util.Max(maxY, sensor.y, sensor.beacon.y)
		sensors[i] = sensor
	}
	return sensors
}

func buildMap(sensors []Sensor) Map {
	// allocate
	fmt.Printf("Map dim: x=%d-%d, y=%d-%d\n", minX, maxX, minY, maxY)
	m := make(Map, maxX-minX+1)
	for x := 0; x < len(m); x++ {
		m[x] = make([]byte, maxY-minY+1)
		for y := 0; y < len(m[x]); y++ {
			m[x][y] = '.'
		}
	}
	// plot
	for _, s := range sensors {
		m[s.x-minX][s.y-minY] = 'S'
		m[s.beacon.x-minX][s.beacon.y-minY] = 'B'
	}
	return m
}

func print(m *Map) {
	fmt.Println("-----------------")
	for y := 0; y < len((*m)[0]); y++ {
		for x := 0; x < len(*m); x++ {
			fmt.Print(string((*m)[x][y]))
		}
		fmt.Println()
	}
}

func markPoint(x, y int, m *Map) {
	if x < 0 || x >= len(*m) || y < 0 || y >= len((*m)[0]) {
		return
	}
	if (*m)[x][y] == '.' {
		(*m)[x][y] = '#'
	}
}

func plotCoverage(s Sensor, m *Map) {
	fmt.Printf("Plot %+v\n", s)
	// plot area of manhattan distance beacon is from sensor
	dx := util.Abs(s.beacon.x - s.x)
	dy := util.Abs(s.beacon.y - s.y)
	dist := dx + dy

	y := 0
	for x := dist; x >= 0; x-- {
		for x1 := s.x + x - minX; x1 > s.x-x+1; x1-- {
			markPoint(x1, s.y+y, m)
			markPoint(x1, s.y-y, m)
		}
		y++
	}
}

func Main(testmode bool) {
	var input []string
	var testpos int
	if testmode {
		input = util.ReadInput("day15/test.txt", "").Lines
		testpos = 10
	} else {
		input = util.ReadInput("day15/day15.txt", "").Lines
		testpos = 2000000
	}

	sensors := parseInput(input)

	// part 1
	m := buildMap(sensors)
	// print(&m)
	for _, s := range sensors {
		plotCoverage(s, &m)
	}
	// print(&m)
	// count covered positions at y=testpos
	count := 0
	for x := 0; x < len(m); x++ {
		if m[x][testpos] == '#' {
			count++
		}
	}
	fmt.Printf("Covered pos at y=%d: %d\n", testpos, count)
}
