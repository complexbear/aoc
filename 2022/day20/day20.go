package day20

import (
	"aoc/util"
	"container/list"
	"fmt"

	"golang.org/x/exp/slices"
)

func parseInput(text []string) []int {
	numbers := make([]int, len(text))
	for i, s := range text {
		numbers[i] = util.StrToInt(s)
	}
	return numbers
}

func printList(l *list.List) {
	// e := l.Front()
	// for e != l.Back() {
	// 	fmt.Printf("%d,", e.Value.(int))
	// 	e = e.Next()
	// }
	// fmt.Printf("%d\n", e.Value.(int))
}

func decode(orig []int) []int {

	// build
	// map values to list nodes for fast lookup
	valueMap := make(map[int]*list.Element, len(orig))
	buffer := list.New()
	for _, n := range orig {
		e := buffer.PushBack(n)
		valueMap[n] = e
	}
	printList(buffer)

	// decode
	for _, n := range orig {
		if n == 0 {
			continue
		}

		// lookup element to move
		srcElem := valueMap[n]
		mark := srcElem

		// find insert point
		for i := 0; i < util.Abs(n); i++ {
			if n > 0 {
				mark = mark.Next()
				if mark == nil {
					mark = buffer.Front()
				}
			} else {
				mark = mark.Prev()
				if mark == nil {
					mark = buffer.Back()
				}
			}
		}
		if mark == buffer.Front() {
			buffer.MoveAfter(srcElem, buffer.Back())
		} else if mark == buffer.Back() {
			buffer.MoveBefore(srcElem, buffer.Front())
		} else if n > 0 {
			buffer.MoveAfter(srcElem, mark)
		} else {
			buffer.MoveBefore(srcElem, mark)
		}
		printList(buffer)
	}

	// convert back to list
	printList(buffer)
	result := make([]int, len(orig))
	l := buffer.Front()
	for i := 0; i < len(orig); i++ {
		result[i] = l.Value.(int)
		l = l.Next()
	}
	return result
}

func calcResult(numbers []int) int {
	result := 0
	idxs := [3]int{1000, 2000, 3000}
	size := len(numbers)
	startIdx := slices.Index(numbers, 0)
	for _, i := range idxs {
		r := numbers[(startIdx+i)%size]
		result += r
	}
	return result
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day20/test.txt", "").Lines
	} else {
		input = util.ReadInput("day20/day20.txt", "").Lines
	}

	numbers := parseInput(input)
	fmt.Println(numbers)

	decoded := decode(numbers)
	fmt.Println(decoded)

	fmt.Println("Result:", calcResult(decoded))
}
