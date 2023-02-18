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

func (p *Position) subtract(other Position) {
	p.x -= other.x
	p.y -= other.y
}

func mapVal(p *Position) byte {
	return ROUTE_MAP[p.y][p.x]
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
	emptyRow := make([]byte, MAX_X)
	for i := 0; i < MAX_X; i++ {
		emptyRow[i] = ' '
	}
	routeMap := make([][]byte, MAX_Y)
	for i, t := range text {
		routeMap[i] = emptyRow
		copy([]byte(t), routeMap[i])
	}
	POS = Position{x: strings.Index(text[0], "."), y: 0}
	fmt.Println(POS)
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

func wrap(pos Position) Position {
	// expect pos to be off the map
	if pos.x >= MAX_X {
		pos.x = 0
		for mapVal(&pos) == ' ' {
			pos.x++
		}
	} else if pos.x >= MAX_X {
		pos.x = MAX_X - 1
		for mapVal(&pos) == ' ' {
			pos.x--
		}
	} else if pos.y == -1 {
		pos.y = MAX_Y - 1
		for mapVal(&pos) == ' ' {
			pos.y--
		}
	} else if pos.y == MAX_Y {
		pos.y = 0
		for mapVal(&pos) == ' ' {
			pos.y++
		}
	} else if mapVal(&pos) == ' ' {
		move := FACING_MOVES[FACING]
		pos.subtract(move)
		for mapVal(&pos) != ' ' {
			pos.subtract(move)
		}
		pos.add(move)
	}
	return pos
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
	if instruction.move != 0 {
		for i := 0; i < instruction.move; i++ {
			// printMap()
			next_pos.add(FACING_MOVES[FACING])
			if next_pos.x >= MAX_X || next_pos.x == -1 || next_pos.y == -1 || next_pos.y == MAX_Y {
				// wrap position
				next_pos = wrap(next_pos)
			}
			val := mapVal(&next_pos)
			if val == ' ' {
				// wrap position
				next_pos = wrap(next_pos)
				val = mapVal(&next_pos)
			}
			if val == '#' {
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
	fmt.Println("Result:", finalResult())
}
