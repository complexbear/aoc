package day7

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name     string
	files    []File
	children map[string]*Dir
	parent   *Dir
}

func NewDir(name string, parent *Dir) *Dir {
	path := name
	if parent != nil {
		path = parent.name + "/" + path
	}
	d := Dir{
		name:     path,
		parent:   parent,
		children: make(map[string]*Dir),
	}
	return &d
}

func (d *Dir) TotalSize(sizes *map[string]int) int {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, c := range d.children {
		size += c.TotalSize(sizes)
	}
	// fmt.Printf("Dir %s total size: %d\n", d.name, size)
	(*sizes)[d.name] = size
	return size
}

func (d *Dir) Print(level int) {
	indent := make([]byte, level)
	for i := 0; i < level; i++ {
		indent[i] = '\t'
	}
	for _, f := range d.files {
		fmt.Printf("%s%d\t%s\n", indent, f.size, f.name)
	}
	for _, dir := range d.children {
		fmt.Printf("%sdir\t%s\n", indent, dir.name)
		dir.Print(level + 1)
	}
}

// Globals
var root *Dir = NewDir("", nil)
var currentDir *Dir = root

func execCmd(cmdargs []string) {
	switch cmdargs[0] {
	case "cd":
		{
			if cmdargs[1] == "/" {
				currentDir = root
			} else if cmdargs[1] == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.children[cmdargs[1]]
			}
		}
	case "ls":
		{
			// No Op - will read in the main loop
		}
	default:
		fmt.Printf("Unknown cmd: %v\n", cmdargs)
		os.Exit(1)
	}
}

func generateFilesystem(input *[][]string) {

	for _, text := range *input {
		if text[0] == "$" {
			// Exec command
			execCmd(text[1:])
		} else {
			// Read response
			if text[0] == "dir" {
				name := text[1]
				(*currentDir).children[name] = NewDir(name, currentDir)
			} else {
				size, _ := strconv.Atoi(text[0])
				f := File{size: size, name: text[1]}
				(*currentDir).files = append((*currentDir).files, f)
			}
		}
	}
}

func Main(testmode bool) {
	var input [][]string
	if testmode {
		input = util.ReadInput("day7/test.txt", " ").Tokens
	} else {
		input = util.ReadInput("day7/day7.txt", " ").Tokens
	}

	generateFilesystem(&input)
	root.Print(0)

	// Calc dir total sizes
	sizes := map[string]int{}
	root.TotalSize(&sizes)
	fmt.Println(sizes)

	// Find those < 100000 and generate sum
	totalSize := 0
	for _, size := range sizes {
		if size <= 100000 {
			totalSize += size
		}
	}
	fmt.Printf("Dir size total: %d\n", totalSize)
}
