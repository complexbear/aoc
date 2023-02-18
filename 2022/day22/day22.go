package day22

import (
	"aoc/util"
	"fmt"
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

func printMap() {
	m := &ROUTE_MAP
	for y, cols := range *m {
		fmt.Printf("%03d: ", y)
		for _, val := range cols {
			fmt.Printf("%s ", string(val))
		}
		fmt.Println()
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
	fmt.Println(path)
	return path
}

func parseMap(text []string) Map {
	MAX_X = len(text[0])
	MAX_Y = len(text)
	routeMap := make([][]byte, MAX_X)
	for i, t := range text {
		routeMap[i] = []byte(t)
	}
	POS = Position{x: strings.Index(text[0], "."), y: 0}
	fmt.Println(POS)
	return routeMap
}

func finalResult() int {
	row := POS.x + 1
	col := POS.y + 1
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

func wrap(pos Position) Position {
	return pos
}

func rotate(direction byte) byte {
	if direction == 'R' {
		switch FACING_MOVES[FACING] {
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
		switch FACING_MOVES[FACING] {
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
	if instruction.move != 0 {
		for i := 0; i < instruction.move; i++ {
			next_pos.add(FACING_MOVES[FACING])
			if next_pos.x >= MAX_X {
				// wrap position
				next_pos = wrap(next_pos)
			}
			val := ROUTE_MAP[next_pos.y][next_pos.x]
			if val == ' ' {
				// wrap position
				next_pos = wrap(next_pos)
				val = ROUTE_MAP[next_pos.y][next_pos.x]
			}
			if val == '#' {
				break
			}
			curr_pos = next_pos
		}
		POS = curr_pos
		ROUTE_MAP[POS.y][POS.x] = FACING
	} else {
		FACING = rotate(instruction.direction)
	}

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
	fmt.Println("Result:", finalResult())
}
