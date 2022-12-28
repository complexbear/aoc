package day18

import (
	"aoc/util"
	"fmt"
	"strconv"
)

var verticies map[Droplet]int = make(map[Droplet]int, 0)

type Droplet struct {
	x int
	y int
	z int
}

func (d Droplet) Area() int {
	return (2 * d.x) + (2 * d.y) + (2 * d.z)
}

func parseInput(input [][]string) []Droplet {
	droplets := make([]Droplet, len(input))
	for i, tokens := range input {
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		droplets[i] = Droplet{x: x, y: y, z: z}
	}
	return droplets
}

func mapVerticies(droplets []Droplet) {
	for _, d := range droplets {
		// generate neighbour positions
		neighbours := make([]Droplet, 6)
		neighbours[0] = Droplet{x:d.x, y:d.y, z:d.z+1}
		neighbours[1] = Droplet{x:d.x, y:d.y, z:d.z-1}
		neighbours[2] = Droplet{x:d.x, y:d.y+1, z:d.z}
		neighbours[3] = Droplet{x:d.x, y:d.y-1, z:d.z}
		neighbours[4] = Droplet{x:d.x+1, y:d.y, z:d.z}
		neighbours[5] = Droplet{x:d.x-1, y:d.y, z:d.z}

		// are any neighbours in the verticies map
		verticies[d] = 6
		for _, n := range neighbours {
			_, exists := verticies[n]
			if exists {
				verticies[d] -= 1
				verticies[n] -= 1
			}
		}
	}
}

func Main(testmode bool) {
	var input [][]string
	if testmode {
		input = util.ReadInput("day18/test.txt", ",").Tokens
	} else {
		input = util.ReadInput("day18/day18.txt", ",").Tokens
	}

	droplets := parseInput(input)
	// fmt.Println(droplets)

	mapVerticies(droplets)
	// fmt.Println(verticies)

	area := 0
	for _, v := range verticies {
		area += v
	}
	fmt.Println(area)
}
