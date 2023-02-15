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
	size := len(orig)
	origValues := make([]*list.Element, size)
	buffer := list.New()
	for i, n := range orig {
		e := buffer.PushBack(n)
		origValues[i] = e
	}
	printList(buffer)

	// decode
	for j := 0; j < size; j++ {
		// lookup element to move
		srcElem := origValues[j]
		mark := srcElem

		n := srcElem.Value.(int)
		n = n % (size - 1)
		if n == 0 {
			continue
		}

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
		if n > 0 {
			if mark == buffer.Back() {
				buffer.MoveBefore(srcElem, buffer.Front())
			} else {
				buffer.MoveAfter(srcElem, mark)
			}
		} else {
			if mark == buffer.Front() {
				buffer.MoveAfter(srcElem, buffer.Back())
			} else {
				buffer.MoveBefore(srcElem, mark)
			}
		}
		printList(buffer)
	}

	// convert back to list
	printList(buffer)
	result := make([]int, size)
	l := buffer.Front()
	for i := 0; i < size; i++ {
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
		fmt.Printf("%dth value: %d\n", i, r)
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
