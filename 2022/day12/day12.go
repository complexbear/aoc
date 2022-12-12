package day12

/*
Find shortest path between S, E in a grid with constraints on possible moves.
Always try to move closer to E in each move.

Track positions at which there is >1 possible move.
This will be used to backtrack when:
	* there is a dead end
	* we will revisit an already visited node


I know I should use Dijkstra - will prob need it for the full input
*/

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type HeightMap [][]byte
type Position struct {
	x int
	y int
}
type Visited map[Position]struct{}
type History []Position

var height int
var width int
var start Position
var end Position
var visited Visited = Visited{}
var history History = History{}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMap(input []string) HeightMap {
	height = len(input)
	width = len(input[0])

	m := make([][]byte, width)
	for i := 0; i < width; i++ {
		m[i] = make([]byte, height)
	}
	for y, line := range input {
		for x, v := range line {
			m[x][y] = byte(v)
		}
	}
	return m
}

func printPath() {
	m := make([][]string, height)
	for i := 0; i < height; i++ {
		m[i] = make([]string, width)
		for j := 0; j < width; j++ {
			m[i][j] = ". "
		}
	}
	for i, v := range history {
		val := strconv.Itoa(i)
		if i < 10 {
			val += " "
		}
		m[v.y][v.x] = val
	}
	for i := 0; i < height; i++ {
		fmt.Println(m[i])
	}
}

func findAndReplace(m HeightMap, old byte, new byte) Position {
	for x, col := range m {
		for y, v := range col {
			if v == old {
				m[x][y] = new
				return Position{x, y}
			}
		}
	}
	return Position{}
}

func nextPosition(pos Position, m *HeightMap) Position {
	currentHeight := (*m)[pos.x][pos.y]

	// search for a valid hop
	searchOptions := []Position{
		{pos.x + 1, pos.y},
		{pos.x, pos.y + 1},
		{pos.x - 1, pos.y},
		{pos.x, pos.y - 1},
	}
	// sort by dist from target
	sort.Slice(searchOptions, func(i, j int) bool {

		// find direction to go
		dx_i := float64(end.x - searchOptions[i].x)
		dy_i := float64(end.y - searchOptions[i].y)
		dx_j := float64(end.x - searchOptions[j].x)
		dy_j := float64(end.y - searchOptions[j].y)
		dist_i := float64(0)
		dist_j := float64(0)
		if dx_i == 0 || dy_i == 0 {
			if dx_i == 0 {
				dist_i = dy_i
			}
			if dy_i == 0 {
				dist_i = dx_i
			}
		} else {
			dist_i = float64(math.Sqrt(math.Pow(dx_i, 2) + math.Pow(dy_i, 2)))
		}
		if dx_j == 0 || dy_j == 0 {
			if dx_j == 0 {
				dist_i = dy_j
			}
			if dy_j == 0 {
				dist_i = dx_j
			}
		} else {
			dist_i = float64(math.Sqrt(math.Pow(dx_j, 2) + math.Pow(dy_j, 2)))
		}
		return dist_i < dist_j
	})
	for _, p := range searchOptions {

		// check we've not been here before
		_, exists := visited[p]
		if exists {
			continue
		}

		// check this position is sane
		if p.x < 0 || p.x >= width || p.y < 0 || p.y >= height {
			continue
		}

		// check height of next hop
		nextHeight := (*m)[p.x][p.y]
		if nextHeight-1 > currentHeight {
			continue
		}

		return p
	}
	fmt.Println("PANIC")
	os.Exit(1)
	return Position{}
}

func explore(pos Position, m *HeightMap) {
	fmt.Printf("explore %+v\n", pos)
	visited[pos] = struct{}{}
	history = append(history, pos)
	if pos == end {
		return
	}
	nextPos := nextPosition(pos, m)
	explore(nextPos, m)
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day12/test.txt", "").Lines
	} else {
		input = util.ReadInput("day12/day12.txt", "").Lines
	}
	heightMap := makeMap(input)

	start = findAndReplace(heightMap, 'S', 'a')
	end = findAndReplace(heightMap, 'E', 'z')
	fmt.Printf("W: %d, H: %d, S: %d, E: %d\n", width, height, start, end)

	explore(start, &heightMap)
	fmt.Printf("Path len: %d\n", len(visited)-1)
	// fmt.Println(history)
	printPath()
}
