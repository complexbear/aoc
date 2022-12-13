package day13

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	val     int
	list    List
	divider bool
}
type List []Item
type Packets []List

func readPackets(input *[]string) Packets {
	packets := make(Packets, 0)
	for i := 0; i < len(*input); i += 3 {
		l1 := make(List, 0)
		l2 := make(List, 0)
		t1 := (*input)[i]
		t2 := (*input)[i+1]
		readLists(t1, &l1)
		readLists(t2, &l2)
		packets = append(packets, l1, l2)
	}
	return packets
}

func readLists(text string, l *List) int {
	for i := 1; i < len(text); {

		if text[i] == ',' {
			i++
			continue
		}
		if text[i] == '[' {
			subItem := Item{list: make(List, 0)}
			i += readLists(text[i:], &subItem.list) + 1
			*l = append(*l, subItem)
			continue
		}
		if text[i] == ']' {
			return i
		}
		// read numeric

		end := strings.IndexFunc(text[i:], func(x rune) bool {
			return x == ',' || x == ']'
		})
		val, _ := strconv.Atoi(text[i : end+i])
		item := Item{list: nil, val: val}
		*l = append(*l, item)
		i += end
	}
	return len(text)
}

func print(l *List) {
	fmt.Printf("[")
	for idx, i := range *l {
		if i.list == nil {
			fmt.Print(i.val)
		} else {
			print(&i.list)
		}
		if idx != len(*l)-1 {
			fmt.Print(",")
		}
	}
	fmt.Printf("]")
}

// Ordering is represented as
//  1 - ordered
//  0 - undecided
// -1 - not ordered

func compareVal(l, r int) int {
	return r - l
}

func compareLists(l, r *List) int {
	ordered := 0
	li := 0
	ri := 0

	for {
		if li == len(*l) || ri == len(*r) {
			ordered = compareVal(len(*l), len(*r))
			break
		}
		if (*l)[li].list == nil && (*r)[ri].list == nil {
			ordered = compareVal((*l)[li].val, (*r)[ri].val)
			if ordered != 0 {
				break
			}
		} else if (*l)[li].list != nil || (*r)[ri].list != nil {
			a := (*l)[li].list
			b := (*r)[ri].list
			if (*l)[li].list == nil {
				a = List{{val: (*l)[li].val}}
			}
			if (*r)[ri].list == nil {
				b = List{{val: (*r)[ri].val}}
			}
			ordered = compareLists(&a, &b)
			if ordered != 0 {
				break
			}
		} else if len((*r)[ri+1:]) == 0 && len((*l)[li+1:]) > 0 {
			ordered = -1
			break
		}

		// next items
		li++
		ri++
	}
	return ordered
}

func isOrdered(o int) bool {
	if o < 0 {
		return false
	}
	return true
}

func orderPackets(packets *Packets) {
	sort.Slice(*packets, func(i, j int) bool {
		p1 := (*packets)[i]
		p2 := (*packets)[j]
		return isOrdered(compareLists(&p1, &p2))
	})
}

func Main(testmode bool) {
	var input []string
	if testmode {
		input = util.ReadInput("day13/test.txt", "").Lines
	} else {
		input = util.ReadInput("day13/day13.txt", "").Lines
	}
	packets := readPackets(&input)

	// part 1
	orderedIdx := 0
	for i := 0; i < len(packets)/2; i++ {
		p1 := &packets[i*2]
		p2 := &packets[i*2+1]
		print(p1)
		fmt.Println()
		print(p2)
		fmt.Println()
		ordered := isOrdered(compareLists(p1, p2))
		if ordered {
			orderedIdx += i + 1
		}
		fmt.Printf("%d: Ordered: %v\n", i+1, ordered)
	}
	fmt.Printf("Sum idx: %d\n", orderedIdx)

	// part 2
	extraText := []string{
		"[[2]]",
		"[[6]]",
	}
	extraPackets := readPackets(&extraText)
	extraPackets[0][0].divider = true
	extraPackets[1][0].divider = true
	packets = append(packets, extraPackets...)

	orderPackets(&packets)
	dividers := 1
	for i, p := range packets {
		print(&p)
		if len(p) > 0 && p[0].divider {
			fmt.Printf("  - Divider packet at %d", i+1)
			dividers *= i+1
		}
		fmt.Println()
	}
	fmt.Printf("Dividers product: %d\n", dividers)
}
