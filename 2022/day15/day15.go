package day15

import (
	"aoc/util"
	"fmt"
	"regexp"
)

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
var intersectLine int
var intersectCoverage map[int]byte = make(map[int]byte, 0)
var bMin int
var bMax int

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
		sensors[i] = sensor

		if sensor.y == intersectLine {
			intersectCoverage[sensor.x] = 'S'
		}
		if sensor.beacon.y == intersectLine {
			intersectCoverage[sensor.beacon.x] = 'B'
		}
	}
	return sensors
}

func plotCoverage(s Sensor) {
	// plot area of manhattan distance beacon is from sensor
	dx := util.Abs(s.beacon.x - s.x)
	dy := util.Abs(s.beacon.y - s.y)
	dist := dx + dy

	// is intersect line within dist?
	for x := minX - dist; x <= maxX+dist; x++ {
		intersectDist := util.Abs(x-s.x) + util.Abs(intersectLine-s.y)
		if intersectDist <= dist && intersectCoverage[x] == 0 {
			intersectCoverage[x] = '#'
		}
	}
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day15/test.txt", "").Lines
		intersectLine = 10
		bMax = 20
	} else {
		input = util.ReadInput("day15/day15.txt", "").Lines
		intersectLine = 2000000
		bMax = 4000000
	}

	sensors := parseInput(input)

	// part 1
	// count covered positions at y=intersetLine
	for _, s := range sensors {
		plotCoverage(s)
	}
	count := 0
	for _, v := range intersectCoverage {
		if v == '#' {
			count++
		}
	}
	fmt.Printf("\nCovered pos at y=%d: %d\n", intersectLine, count)

	// part 2
	minX = 0
	var x int
	var y int
	for y = 0; y <= bMax; y++ {
		fmt.Printf("testing line %d\n", y)
		intersectLine = y
		intersectCoverage = make(map[int]byte, 0)
		for _, s := range sensors {
			plotCoverage(s)
		}
		for x = 0; x <= bMax; x++ {
			_, exists := intersectCoverage[x]
			if !exists {
				goto done
			}
		}
	}
done:
	fmt.Printf("x=%d, y=%d\n", x, y)
}

