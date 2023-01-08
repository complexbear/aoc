package day18

import (
	"aoc/util"
	"fmt"
	"strconv"
)

var verticies map[Droplet]int = make(map[Droplet]int, 0)
var minX map[Droplet]int = map[Droplet]int{}
var maxX map[Droplet]int = map[Droplet]int{}
var minY map[Droplet]int = map[Droplet]int{}
var maxY map[Droplet]int = map[Droplet]int{}
var minZ map[Droplet]int = map[Droplet]int{}
var maxZ map[Droplet]int = map[Droplet]int{}

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

func mapVerticiesPart1(droplets []Droplet) {
	for _, d := range droplets {
		// generate neighbour positions
		neighbours := make([]Droplet, 6)
		neighbours[0] = Droplet{x: d.x, y: d.y, z: d.z + 1}
		neighbours[1] = Droplet{x: d.x, y: d.y, z: d.z - 1}
		neighbours[2] = Droplet{x: d.x, y: d.y + 1, z: d.z}
		neighbours[3] = Droplet{x: d.x, y: d.y - 1, z: d.z}
		neighbours[4] = Droplet{x: d.x + 1, y: d.y, z: d.z}
		neighbours[5] = Droplet{x: d.x - 1, y: d.y, z: d.z}

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

func calcMinMax(droplets []Droplet) {
	for _, d := range droplets {
		// X
		refD := Droplet{y:d.y, z:d.z}
		minD := minX[refD]
		maxD := maxX[refD]
		minX[refD] = util.Min(minD, d.x)
		maxX[refD] = util.Max(maxD, d.x)

		// Y
		refD = Droplet{x:d.x, z:d.z}
		minD = minY[refD]
		maxD = maxY[refD]
		minY[refD] = util.Min(minD, d.y)
		maxY[refD] = util.Max(maxD, d.y)

		// X
		refD = Droplet{x:d.x, y:d.y}
		minD = minZ[refD]
		maxD = maxZ[refD]
		minZ[refD] = util.Min(minD, d.z)
		maxZ[refD] = util.Max(maxD, d.z)
	}
}

func mapVerticiesPart2(droplets []Droplet) {
	for _, d := range droplets {
		// generate neighbour positions
		neighbours := make([]Droplet, 6)
		neighbours[0] = Droplet{x: d.x, y: d.y, z: d.z + 1}
		neighbours[1] = Droplet{x: d.x, y: d.y, z: d.z - 1}
		neighbours[2] = Droplet{x: d.x, y: d.y + 1, z: d.z}
		neighbours[3] = Droplet{x: d.x, y: d.y - 1, z: d.z}
		neighbours[4] = Droplet{x: d.x + 1, y: d.y, z: d.z}
		neighbours[5] = Droplet{x: d.x - 1, y: d.y, z: d.z}

		for _, n := range neighbours {
			refD := Droplet{y:n.y, z:n.z}
			xInside := n.x > minX[refD] && n.x < maxX[refD]
			refD = Droplet{x:n.x, z:n.z}
			yInside := n.y > minY[refD] && n.y < maxY[refD]
			refD = Droplet{x:n.x,y:n.y}
			zInside := n.z > minZ[refD] && n.z < maxZ[refD]
			if xInside && yInside && zInside {
				verticies[n] = 0
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

	mapVerticiesPart1(droplets)
	fmt.Println(verticies)

	// Part 1
	area := 0
	for _, v := range verticies {
		area += v
	}
	fmt.Println(area)

	// Part 2
	// look over the vertices, and see which lie within the min/max of each x,y,z plane
	// for those that do, reduce their value to 0 and re-sum the area
	mapDroplets := make(map[Droplet]int, len(droplets))
	for _, d := range droplets {
		mapDroplets[d] = verticies[d]
	}

	calcMinMax(droplets)
	mapVerticiesPart2(droplets)
	// fmt.Println(verticies)
	interiorArea := 0
	for _, v := range verticies {
		interiorArea += v
	}
	fmt.Printf("Total area: %d\tExterior area: %d\n", area, interiorArea)
}
