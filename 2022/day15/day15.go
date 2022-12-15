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

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day15/test.txt", "").Lines
	} else {
		input = util.ReadInput("day15/day15.txt", "").Lines
	}

	sensors := parseInput(input)
	// for _, s := range sensors {
	// 	fmt.Printf("%+v\n", s)
	// }
	m := buildMap(sensors)
	print(&m)
}
