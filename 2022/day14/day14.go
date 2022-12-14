package day14

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}
type Map [][]byte

var maxY int
var maxX int = 1000
var part2 bool = false

func print(m *Map) {
	for y := 0; y < len((*m)[0]); y++ {
		for x := 0; x < len(*m); x++ {
			val := (*m)[x][y]
			if val == 0 {
				val = '.'
			}
			fmt.Printf(" %s", string(val))
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
}

func buildMap(input *[][]string) Map {
	// read coords of lines
	lines := make([][]Point, 0)
	for _, tokens := range *input {
		line := make([]Point, 0)
		for _, t := range tokens {
			point := strings.Split(t, ",")
			x, _ := strconv.Atoi(point[0])
			y, _ := strconv.Atoi(point[1])
			if y > maxY {
				maxY = y
			}
			line = append(line, Point{x: x, y: y})
		}
		lines = append(lines, line)
	}
	if part2 {
		maxY += 2
	}

	// build map
	m := make(Map, maxX)
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, maxY+1)
	}
	// draw rocks
	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			a := line[i-1]
			b := line[i]
			m[a.x][a.y] = '#'
			m[b.x][b.y] = '#'

			dx := b.x - a.x
			if dx != 0 {
				dx /= util.Abs(dx)
			}
			dy := b.y - a.y
			if dy != 0 {
				dy /= util.Abs(dy)
			}
			for b != a {
				a.x += dx
				a.y += dy
				m[a.x][a.y] = '#'
			}
		}
	}
	if part2 {
		for i := 0; i < len(m); i++ {
			m[i][maxY] = '#'
		}
	}
	return m
}

func validMove(pt Point, m *Map) bool {
	if pt.x < 0 || pt.x >= len(*m) || pt.y > maxY {
		return false
	}
	if (*m)[pt.x][pt.y] == 0 {
		return true
	}
	return false
}

func offMap(pt Point) bool {
	if pt.x < 0 || pt.x >= maxX || pt.y > maxY {
		return true
	}
	return false
}

func dropGrain(m *Map) bool {
	origin := Point{x: 500, y: 0}
	pt := origin
	for {
		down := Point{x: pt.x, y: pt.y + 1}
		if validMove(down, m) {
			pt = down
			continue
		}
		left := Point{x: pt.x - 1, y: pt.y + 1}
		if validMove(left, m) {
			pt = left
			continue
		}
		right := Point{x: pt.x + 1, y: pt.y + 1}
		if validMove(right, m) {
			pt = right
			continue
		}

		// part 1
		if part2 == false && (offMap(down) || offMap(right) || offMap(left)) {
			return false
		}

		// at rest
		(*m)[pt.x][pt.y] = 'o'

		// part 2
		if pt == origin {
			return false
		}
		return true
	}
}

func pourSand(m *Map) int {
	i := 0
	for {
		landed := dropGrain(m)
		// print(m)
		if !landed {
			break
		}
		i++
	}
	return i
}

func Main(testmode bool) {
	var input [][]string
	if testmode {
		input = util.ReadInput("day14/test.txt", " -> ").Tokens
	} else {
		input = util.ReadInput("day14/day14.txt", " -> ").Tokens
	}

	// part 1
	m := buildMap(&input)
	// print(&m)
	part1 := pourSand(&m)
	// print(&m)

	// part 2
	part2 = true
	m = buildMap(&input)
	// print(&m)
	part2 := pourSand(&m) + 1 // doesn't count the origin grain
	// print(&m)

	fmt.Printf("Grains of sand part 1: %d\n", part1)
	fmt.Printf("Grains of sand part 2: %d\n", part2)
}
