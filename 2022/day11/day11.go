package day11

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation []string
	test      int
	ifTrue    int
	ifFalse   int
	inspects  int
}

func readMonkeys(input *[]string) []Monkey {
	monkeys := make([]Monkey, 0)
	var monkey Monkey
	for _, text := range *input {
		if text == "" {
			monkeys = append(monkeys, monkey)
			continue
		}
		if strings.Contains(text, "Monkey") {
			monkey = Monkey{}
			continue
		}
		if strings.Contains(text, "Starting items") {
			itemText := strings.Split(text, ":")[1]
			itemText = strings.ReplaceAll(itemText, " ", "")
			itemsStr := strings.Split(itemText, ",")
			items := make([]int, len(itemsStr))
			for i, s := range itemsStr {
				val, _ := strconv.Atoi(s)
				items[i] = val
			}
			monkey.items = items
			continue
		}
		if strings.Contains(text, "Operation") {
			opText := strings.Split(text, " ")
			monkey.operation = []string{opText[5], opText[6], opText[7]}
			continue
		}
		if strings.Contains(text, "Test") {
			testText := strings.Split(text, " ")
			val, _ := strconv.Atoi(testText[len(testText)-1])
			monkey.test = val
			continue
		}
		if strings.Contains(text, "If true") {
			val, _ := strconv.Atoi(string(text[len(text)-1]))
			monkey.ifTrue = val
			continue
		}
		if strings.Contains(text, "If false") {
			val, _ := strconv.Atoi(string(text[len(text)-1]))
			monkey.ifFalse = val
			continue
		}
	}
	// add last monkey
	monkeys = append(monkeys, monkey)
	return monkeys
}

func applyOperation(opStr []string, old int) int {
	a := old
	b := old
	if opStr[0] != "old" {
		a, _ = strconv.Atoi(opStr[0])
	}
	if opStr[2] != "old" {
		b, _ = strconv.Atoi(opStr[2])
	}
	switch opStr[1] {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	}
	return 0
}

func inspect(m *Monkey, monkeys *[]Monkey) {
	for _, item := range m.items {
		(*m).inspects++
		// bump worry
		item = applyOperation(m.operation, item)
		// divide by 3
		item /= 3
		// test
		target := 0
		if item%m.test == 0 {
			target = m.ifTrue
		} else {
			target = m.ifFalse
		}
		// throw item
		(*monkeys)[target].items = append((*monkeys)[target].items, item)
		(*m).items = make([]int, 0)
	}
}

func round(monkeys *[]Monkey) {
	for i := 0; i < len(*monkeys); i++ {
		inspect(&(*monkeys)[i], monkeys)
	}
}

func print(monkeys *[]Monkey) {
	fmt.Println("-------------------")
	for _, m := range *monkeys {
		fmt.Printf("%+v\n", m)
	}
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day11/test.txt", "").Lines
	} else {
		input = util.ReadInput("day11/day11.txt", "").Lines
	}

	monkeys := readMonkeys(&input)
	print(&monkeys)

	// Part 1
	for i := 0; i < 20; i++ {
		round(&monkeys)
	}
	print(&monkeys)
	// Most active monkeys
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspects
	}
	sort.Ints(inspections)
	fmt.Printf("Monkey business: %d\n", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}
