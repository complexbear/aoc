package day21

import (
	"aoc/util"
	"fmt"
	"math"
	"strings"
)

type Monkey struct {
	Name  string
	Value int
	Left  string
	Right string
	Op    string
}

type MonkeyMap map[string]*Monkey

func parseMonkeys(input [][]string) MonkeyMap {
	monkeys := make(MonkeyMap, len(input))
	for i := 0; i < len(input); i++ {
		tokens := input[i]
		name := strings.Replace(tokens[0], ":", "", 1)
		m := Monkey{Name: name}
		if len(tokens) == 2 {
			m.Value = util.StrToInt(tokens[1])
		} else {
			m.Left = tokens[1]
			m.Op = tokens[2]
			m.Right = tokens[3]
		}
		monkeys[name] = &m
	}
	return monkeys
}

func calcMonkey(name string, monkeys *MonkeyMap) {
	m := (*monkeys)[name]
	if m.Left != "" && m.Right != "" {
		left := (*monkeys)[m.Left]
		calcMonkey(left.Name, monkeys)
		right := (*monkeys)[m.Right]
		calcMonkey(right.Name, monkeys)
		
		switch m.Op {
		case "+":
			m.Value = left.Value + right.Value
		case "-":
			m.Value = left.Value - right.Value
		case "*":
			m.Value = left.Value * right.Value
		case "/":
			m.Value = left.Value / right.Value
		case "=":
			{
				if left.Value == right.Value {
					m.Value = 1
				} else {
					m.Value = 0
				}
			}
		}
	}
}

func branchSearch(name string, monkeys *MonkeyMap, target string) bool {
	if name == "" {
		return false
	}
	m := (*monkeys)[name]
	if m.Left == target || m.Right == target {
		return true
	}
	return branchSearch(m.Left, monkeys, target) || branchSearch(m.Right, monkeys, target)
}

func branchSolve(name string, monkeys *MonkeyMap, target int) {
	m := (*monkeys)[name]
	humn := (*monkeys)["humn"]
	humn.Value = 0

	a := 0
	b := math.MaxInt
	c := 0

	for true {
		c = (a + b) / 2

		humn.Value = c
		calcMonkey(name, monkeys)
		dx := target - m.Value
		if dx == 0 {
			break
		}

		if dx < 0 {
			a = c
		} else {
			b = c
		}
		// fmt.Println("dx:", dx, "abs:", util.Abs(dx))
	}
}

func part1(monkeys *MonkeyMap) {
	calcMonkey("root", monkeys)
	fmt.Println("Root:", (*monkeys)["root"].Value)
}

func part2(monkeys *MonkeyMap) {
	root := (*monkeys)["root"]
	root.Op = "="
	root.Value = -1

	// which branch is humn not in, this is the expected answer
	humnInLeftBranch := branchSearch(root.Left, monkeys, "humn")
	fmt.Println("humn is in left branch from root:", humnInLeftBranch)
	expectedValue := 0
	if humnInLeftBranch {
		expectedValue = (*monkeys)[root.Right].Value
		branchSolve(root.Left, monkeys, expectedValue)
	} else {
		expectedValue = (*monkeys)[root.Left].Value
		branchSolve(root.Right, monkeys, expectedValue)
	}
	fmt.Println("Part 2 Root expected value:", expectedValue)
	fmt.Println("Part 2 Humn value:", (*monkeys)["humn"].Value)

	part1(monkeys)
}

func Main(testmode bool) {
	var input [][]string
	if testmode {
		input = util.ReadInput("day21/test.txt", " ").Tokens
	} else {
		input = util.ReadInput("day21/day21.txt", " ").Tokens
	}

	monkeys := parseMonkeys(input)

	part1(&monkeys)
	fmt.Println("---------------------")
	part2(&monkeys)
}
