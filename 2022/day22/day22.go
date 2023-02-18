package day22

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Map [][]byte
type Instruction struct {
	move      int
	direction byte
}
type Path []Instruction
type Position struct {
	x int
	y int
}

var MAX_X int
var MAX_Y int
var POS Position
var ROUTE_MAP Map
var PATH Path
var FACING byte = '>'

// Facing direction encoded like
//
//	x:1,  y:0  => right
//	x:-1, y:0  => left
//	x:0,  y:1  => down
//	x:0,  y:-1 => up
var RIGHT = Position{x: 1, y: 0}
var LEFT = Position{x: -1, y: 0}
var UP = Position{x: 0, y: -1}
var DOWN = Position{x: 0, y: 1}
var FACING_MOVES = map[byte]Position{
	'>': RIGHT,
	'<': LEFT,
	'^': UP,
	'v': DOWN,
}

func (p *Position) add(other Position) {
	p.x += other.x
	p.y += other.y
}

func (p *Position) subtract(other Position) {
	p.x -= other.x
	p.y -= other.y
}

func mapVal(p *Position) byte {
	return ROUTE_MAP[p.y][p.x]
}

func printMap() {
	f, _ := os.Create("map.txt")
	w := bufio.NewWriter(f)
	defer f.Close()

	m := &ROUTE_MAP
	for y, cols := range *m {
		fmt.Printf("%03d: ", y)
		fmt.Fprintf(w, "%03d: ", y)
		for _, val := range cols {
			fmt.Printf("%s", string(val))
			fmt.Fprintf(w, "%s", string(val))
		}
		fmt.Println()
		fmt.Fprintln(w)
		w.Flush()
	}
	fmt.Println("-----------------")
}

func parseInput(text []string) {
	var routeText []string
	var pathText string

	for i := 0; i < len(text); i++ {
		if text[i] == "" {
			routeText = text[:i]
			pathText = text[i+1]
		}
	}

	PATH = parsePath(pathText)
	ROUTE_MAP = parseMap(routeText)
}

func parsePath(text string) Path {
	path := Path{}
	var lenText []byte
	for _, c := range []byte(text) {
		if c == 'L' || c == 'R' {
			if len(lenText) > 0 {
				path = append(
					path,
					Instruction{move: util.StrToInt(string(lenText))},
				)
				lenText = []byte{}
			}
			path = append(
				path,
				Instruction{direction: c},
			)
		} else {
			lenText = append(lenText, c)
		}
	}
	if len(lenText) > 0 {
		path = append(
			path,
			Instruction{move: util.StrToInt(string(lenText))},
		)
		lenText = []byte{}
	}
	for i, p := range path {
		fmt.Println(i, ":", p.move, string(p.direction))
	}
	return path
}

func parseMap(text []string) Map {
	MAX_X = 0
	for _, t := range text {
		MAX_X = util.Max(MAX_X, len(t))
	}
	MAX_Y = len(text)
	routeMap := make([][]byte, MAX_Y)
	for i, t := range text {
		row := make([]byte, MAX_X)
		for j := 0; j < MAX_X; j++ {
			if j < len(t) {
				row[j] = t[j]
			} else {
				row[j] = ' '
			}

		}
		routeMap[i] = row
	}
	POS = Position{x: strings.Index(text[0], "."), y: 0}
	return routeMap
}

func finalResult() int {
	row := POS.y + 1
	col := POS.x + 1
	facing := -1
	switch FACING {
	case '>':
		facing = 0
	case 'v':
		facing = 1
	case '<':
		facing = 2
	case '^':
		facing = 3
	}

	return 1000*row + 4*col + facing
}

func step(pos *Position) (Position, bool) {
	next_pos := *pos
	hitwall := false

	for true {
		// handle all cases where position needs to be wrapped
		next_pos.add(FACING_MOVES[FACING])

		if next_pos.x < 0 {
			next_pos.x = MAX_X - 1
		}
		if next_pos.y < 0 {
			next_pos.y = MAX_Y - 1
		}

		// if we run off the end of the map
		next_pos.x %= MAX_X
		next_pos.y %= MAX_Y

		// pos now should be legal
		switch mapVal(&next_pos) {
		case ' ':
			break // run on
		case '#':
			{
				hitwall = true
				return *pos, hitwall
			}
		case '.':
			return next_pos, hitwall
		}
	}
	return *pos, hitwall
}

func rotate(facing byte, direction byte) byte {
	if direction == 'R' {
		switch FACING_MOVES[facing] {
		case LEFT:
			return '^'
		case RIGHT:
			return 'v'
		case UP:
			return '>'
		case DOWN:
			return '<'
		}
	} else {
		switch FACING_MOVES[facing] {
		case LEFT:
			return 'v'
		case RIGHT:
			return '^'
		case UP:
			return '<'
		case DOWN:
			return '>'
		}
	}
	panic("bad rotate")
}

func move(instruction Instruction) {
	curr_pos := POS
	next_pos := POS
	hitwall := false
	if instruction.move != 0 {
		for i := 0; i < instruction.move; i++ {
			// printMap()
			next_pos, hitwall = step(&curr_pos)
			if hitwall {
				break
			}
			curr_pos = next_pos
			ROUTE_MAP[curr_pos.y][curr_pos.x] = FACING
		}
		POS = curr_pos
	} else {
		FACING = rotate(FACING, instruction.direction)
	}
	ROUTE_MAP[POS.y][POS.x] = FACING
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day22/test.txt", "").Lines
	} else {
		input = util.ReadInput("day22/day22.txt", "").Lines
	}

	parseInput(input)
	ROUTE_MAP[POS.y][POS.x] = FACING

	printMap()
	for _, i := range PATH {
		move(i)
	}
	printMap()
	fmt.Println("Result:", finalResult())
}
