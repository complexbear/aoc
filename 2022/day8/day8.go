package day8

import (
	"aoc/util"
	"fmt"
	"strconv"

	"github.com/golang-demos/chalk"
)

type Grid []int

var height int = 0
var width int = 0

func print(trees *Grid, visible *Grid) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			idx := (h * width) + w
			tree := (*trees)[idx]
			text := strconv.Itoa(tree)
			if visible != nil && (*visible)[idx] > 0 {
				fmt.Print(chalk.Red(text))
			} else {
				fmt.Print(text)
			}
		}
		fmt.Println("")
	}
}

func readGrid(input *[]string) *Grid {
	height = len(*input)
	width = len((*input)[0])
	grid := make(Grid, height*width)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			val, _ := strconv.Atoi(string((*input)[h][w]))
			grid[(h*width)+w] = val
		}
	}
	return &grid
}

func rowIdx(h, w int) int {
	return (h * width) + w
}

func colIdx(h, w int) int {
	return (w * height) + h
}

func visibleView(trees *Grid, visible *Grid, direction string) {
	var lastHighestTree int
	var idx func(int, int) int
	if direction == "left" || direction == "right" {
		idx = rowIdx
	} else {
		idx = colIdx
	}
	if direction == "left" || direction == "top" {
		for i := 0; i < height; i++ {
			lastHighestTree = -1
			for j := 0; j < width; j++ {
				i := idx(i, j)
				t := (*trees)[i]
				if t > lastHighestTree {
					lastHighestTree = t
					(*visible)[i] = 1
				}
			}
		}
	} else {
		for i := height - 1; i >= 0; i-- {
			lastHighestTree = -1
			for j := width - 1; j >= 0; j-- {
				i := idx(i, j)
				t := (*trees)[i]
				if t > lastHighestTree {
					lastHighestTree = t
					(*visible)[i] = 1
				}
			}
		}
	}
}

func findVisibleTrees(trees, visible *Grid) {
	visibleView(trees, visible, "left")
	visibleView(trees, visible, "top")
	visibleView(trees, visible, "right")
	visibleView(trees, visible, "bottom")
}

func treeScore(trees *Grid, h, w int) int {
	origH := h
	origW := w
	origTree := (*trees)[rowIdx(h, w)]
	count := 0
	
	
}

func calcScenicScores(trees, scores *Grid) {
	for h:=0; h<height; h++ {
		for w:=0; w<width; w++ {
			(*scores)[rowIdx(h, w)] = treeScore(trees, h, w)
		}
	}
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = []string{
			"30373",
			"25512",
			"65332",
			"33549",
			"35390",
		}
	} else {
		input = util.ReadInput("day8/day8.txt", "").Lines
	}
	trees := readGrid(&input)
	visible := make(Grid, height*width)

	// find visible trees
	findVisibleTrees(trees, &visible)

	totalCount := 0
	for _, val := range visible {
		totalCount += val
	}
	fmt.Printf("Visible trees: %d\n", totalCount)
	print(trees, &visible)
	fmt.Println("-------------------")

	scores := make(Grid, height*width)
	calcScenicScores(trees, &scores)
	print(&scores, nil)

	maxScore := 0
	for _, val := range scores {
		if val > maxScore {
			maxScore = val
		}
	}
	fmt.Printf("Max scenic score: %d\n", maxScore)
}
