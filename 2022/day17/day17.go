package day17

import (
	"aoc/util"
	"bytes"
	"fmt"
	"sort"
)

var log util.Logger = util.Logger{}
var width int = 7
var highestBlock int = 0
var chamber [][]byte
var shapeUse []int

var startPatternIdx int = 0
var patternSize int = 0
var shapesInPattern int = 0

type Shape struct {
	blocks [][]byte
	x      int
	y      int
	width  int
	height int
}

func (s Shape) Print() {
	for _, r := range s.blocks {
		log.Print(string(r))
	}
}

func (s Shape) Collision() bool {
	// check for collision with other blocks
	collision := false
	for i, r := range s.blocks {
		for j, b := range r {
			if s.y+i >= len(chamber) {
				continue
			}
			if b == '#' && chamber[s.y+i][s.x+j] == '#' {
				collision = true
				goto collision
			}
		}
	}
collision:
	return collision
}

func (s *Shape) Move(dir byte) {
	// fmt.Printf("%s|", string(dir))
	move := 0
	if dir == '>' && s.x+s.width+1 <= width {
		move = 1
	}
	if dir == '<' && s.x-1 >= 0 {
		move = -1
	}
	if move != 0 {
		next := Shape(*s)
		next.x += move
		if next.Collision() == false {
			s.x = next.x
		}
	}
}

func (s *Shape) Drop() bool {
	next := Shape(*s)
	next.y--
	collision := next.Collision()
	if collision == false {
		s.y--
	} else {
		s.Settle()
	}
	return collision == false
}

func (s Shape) Settle() {
	// update chamber with shape position
	for i, r := range s.blocks {
		for j, b := range r {
			if b == '#' {
				chamber[s.y+i][s.x+j] = b
				if s.y+i > highestBlock {
					highestBlock = s.y + i
				}
			}
		}
	}
	shapeUse[s.y] += 1
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func parseInput(input []string) []Shape {
	shapes := make([]Shape, 0)
	shape := Shape{}
	for _, text := range input {
		if text == "" {
			shape.width = len(shape.blocks[0])
			shape.height = len(shape.blocks)
			shapes = append(shapes, shape)
			shape = Shape{}
		} else {
			blocks := []byte(text)
			shape.blocks = append([][]byte{blocks}, shape.blocks...)
		}
	}
	shape.width = len(shape.blocks[0])
	shape.height = len(shape.blocks)
	shapes = append(shapes, shape)
	return shapes
}

func newBlock(s Shape, moves []byte) []byte {
	s.x = 2
	s.y = highestBlock + 4
	for len(chamber) <= s.y {
		chamber = append(chamber, make([]byte, width))
		shapeUse = append(shapeUse, 0)
	}
	for midx := 0; midx < len(moves); midx++ {
		m := moves[midx]

		// push
		s.Move(m)

		// drop
		dropped := s.Drop()

		// stuck
		if !dropped {
			return moves[midx+1:]
		}
	}
	s.Settle()
	return []byte{}
}

func compareChambers(a, b [][]byte) bool {
	// return true if a and b match
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if bytes.Compare(a[i], b[i]) != 0 {
			return false
		}
	}
	return true
}

func findPattern() {
	// to manage memory, detect when block pattern repeats
	// and discard the repetitions
	//
	// method: start with top of chamber and search down until
	// same line pattern found. then cut chamber in 2 and compare halves.
	// if they match then the pattern has repeated.
	//
	// we can then set the pattern size and start index

	// bottom up search
	for y := 1; y < highestBlock; y++ {
		patternBlock := chamber[y:highestBlock]
		if len(patternBlock) < 20 {
			continue
		}
		midwayIdx := len(patternBlock) / 2
		patternA := patternBlock[:midwayIdx]
		patternB := patternBlock[midwayIdx:]
		if compareChambers(patternA, patternB) {
			// pattern found,
			startPatternIdx = y + midwayIdx
			patternSize = midwayIdx
			// how many shapes in the pattern
			for i := startPatternIdx; i < startPatternIdx+midwayIdx; i++ {
				shapesInPattern += shapeUse[i]
			}
			return
		}
	}
}

func PrintChamber() {
	for y := len(chamber) - 1; y >= 0; y-- {
		text := fmt.Sprintf("%02d:", y)
		for _, b := range chamber[y] {
			val := b
			if val == 0 {
				val = '.'
			}
			text += string(val)
		}
		log.Print(text)
	}
	log.Print("------------------------")
}

func Main(testmode bool) {
	var input []string
	var shapeCount int
	if testmode {
		input = util.ReadInput("day17/test.txt", "").Lines
		shapeCount = 2022
		log.Debug = false
	} else {
		input = util.ReadInput("day17/day17.txt", "").Lines
		shapeCount = 1000000000000
	}

	orig_moves := []byte(input[0])
	moves := orig_moves

	// init chamber
	chamber = make([][]byte, 5)
	shapeUse = make([]int, 5)
	for i := 0; i < len(chamber); i++ {
		chamber[i] = make([]byte, width)
	}
	for i := 0; i < width; i++ {
		chamber[0][i] = '#'
	}

	input = util.ReadInput("day17/shapes.txt", "").Lines
	shapes := parseInput(input)
	for _, s := range shapes {
		s.Print()
	}

	jumped := false
	cumHighestBlock := 0
	for sidx := 0; sidx < shapeCount; sidx++ {
		s := shapes[sidx%len(shapes)]

		log.Print("Shape %d\n", sidx+1)
		s.Print()
		log.Print(string(moves))
		log.Print("******")

		moves = newBlock(s, moves)
		PrintChamber()

		// add new moves if needed
		if len(moves) < 10 {
			log.Print("extending moves")
			moves = append(moves, orig_moves...)
		}

		// start looking for repeated pattern
		if sidx > len(shapes) && patternSize == 0 {
			findPattern()
		}
		// pattern found, so we should jump ahead as far as we can
		if patternSize != 0 && jumped == false {
			remainingShapes := shapeCount - sidx
			patternRepeats := remainingShapes / shapesInPattern
			cumHighestBlock += patternRepeats * patternSize
			sidx = shapeCount - (remainingShapes % shapesInPattern)
			jumped = true
		}
	}

	log.Debug = true
	// PrintChamber()
	cumHighestBlock += highestBlock
	fmt.Printf("Pattern detected at idx:%d, size:%d\n", startPatternIdx, patternSize)
	fmt.Printf("Height: %d\n", cumHighestBlock)
}
