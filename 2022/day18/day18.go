package day18

import (
	"aoc/util"
	"fmt"
	"strconv"
)

type Droplet struct {
	x int
	y int
	z int
}
type DropletMap map[Droplet]int

var verticies DropletMap = make(DropletMap, 0)
var minDroplet Droplet = Droplet{10000, 10000, 10000}
var maxDroplet Droplet = Droplet{-10000, -10000, -10000}

func cubeArea(d1, d2 Droplet) int {
	dx := d2.x-d1.x 
	dy := d2.y-d1.y
	dz := d2.z-d1.z
	return (dx*dx)*2 + (dy*dy)*2 + (dz*dz)*2 
}

func parseInput(input [][]string) []Droplet {
	droplets := make([]Droplet, len(input))
	for i, tokens := range input {
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])

		minDroplet.x = util.Min(minDroplet.x, x)
		minDroplet.y = util.Min(minDroplet.y, y)
		minDroplet.z = util.Min(minDroplet.z, z)
		maxDroplet.x = util.Max(maxDroplet.x, x)
		maxDroplet.y = util.Max(maxDroplet.y, y)
		maxDroplet.z = util.Max(maxDroplet.z, z)
		droplets[i] = Droplet{x: x, y: y, z: z}
	}
	return droplets
}

func mapVerticies(droplets []Droplet) {
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

func fill(d Droplet, lava *DropletMap, cube *DropletMap) bool {
	_, exists := (*lava)[d]
	if exists {
		return false
	}
	_, exists = (*cube)[d]
	if exists {
		return false
	}
	if d.x < minDroplet.x || d.x > maxDroplet.x {
		return false
	}
	if d.y < minDroplet.y || d.y > maxDroplet.y {
		return false
	}
	if d.z < minDroplet.z || d.z > maxDroplet.z {
		return false
	}
	(*cube)[d] = 1
	return true
}

func floodFill(minD, maxD Droplet, lava *DropletMap) int {
	cubeDropletMap := make(DropletMap, 0)
	searchStack := make([]Droplet, 1)
	searchStack[0] = minD

	for {
		l := len(searchStack)
		if l == 0 {
			break
		}

		d := searchStack[l-1]
		searchStack = searchStack[:l-1]
		if fill(d, lava, &cubeDropletMap) {
			neighbours := make([]Droplet, 6)
			neighbours[0] = Droplet{x: d.x, y: d.y, z: d.z + 1}
			neighbours[1] = Droplet{x: d.x, y: d.y, z: d.z - 1}
			neighbours[2] = Droplet{x: d.x, y: d.y + 1, z: d.z}
			neighbours[3] = Droplet{x: d.x, y: d.y - 1, z: d.z}
			neighbours[4] = Droplet{x: d.x + 1, y: d.y, z: d.z}
			neighbours[5] = Droplet{x: d.x - 1, y: d.y, z: d.z}
			searchStack = append(neighbours, searchStack...)
		}
	}

	cubeDroplets := make([]Droplet, len(cubeDropletMap))
	verticies = make(DropletMap)

	idx := 0
	for d, _ := range cubeDropletMap {
		cubeDroplets[idx] = d
		idx++
	}
	mapVerticies(cubeDroplets)

	area := 0
	for _, v := range verticies {
		area += v
	}
	return area - cubeArea(minD, maxD)
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
	fmt.Println(verticies)

	// Part 1
	area := 0
	for _, v := range verticies {
		area += v
	}
	fmt.Println(area)

	// Part 2
	// start outside the lava and floodfill the space
	// substrct exterior cube faces from total surface area of the
	// mold generated
	mapDroplets := make(DropletMap, len(droplets))
	for _, d := range droplets {
		mapDroplets[d] = verticies[d]
	}

	fmt.Println(minDroplet)
	fmt.Println(maxDroplet)

	interiorSurface := floodFill(minDroplet, maxDroplet, &mapDroplets)

	fmt.Printf("Total area: %d\tExterior area: %d\n", area, interiorSurface)
}
