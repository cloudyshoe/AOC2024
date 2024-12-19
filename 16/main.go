package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/fzipp/astar"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Maze map[utils.Coord]Node

type MazeCell struct {
	Exists bool
	Value  Node
}

var maze = make(Maze)

func (m Maze) ValueAt(c utils.Coord) MazeCell {
	cell, ok := m[c]
	return MazeCell{Exists: ok, Value: cell}
}

type Node struct {
	location utils.Coord
	facing   utils.Coord
	value    rune
}

func rotate(m Maze, n Node) []utils.Coord {
	var a, b, c Node
	if n.facing.Row == 0 {
		aCoord := utils.Coord{Row: n.facing.Col, Col: 0}
		a = m[n.location.Add(aCoord)]
		a.facing = aCoord
		m[a.location] = a

		bCoord := utils.Coord{Row: -n.facing.Col, Col: 0}
		b = m[n.location.Add(bCoord)]
		b.facing = bCoord
		m[b.location] = b
	} else {
		aCoord := utils.Coord{Row: 0, Col: n.facing.Row}
		a = m[n.location.Add(aCoord)]
		a.facing = aCoord
		m[a.location] = a

		bCoord := utils.Coord{Row: 0, Col: -n.facing.Row}
		b = m[n.location.Add(bCoord)]
		b.facing = bCoord
		m[b.location] = b
	}

	c = m[n.location.Add(n.facing)]
	c.facing = n.facing
	maze[c.location] = c

	return []utils.Coord{a.location, b.location, c.location}
}

func (m Maze) Neighbours(c utils.Coord) iter.Seq[utils.Coord] {
	return func(yield func(utils.Coord) bool) {
		fmt.Println(maze[c].location, maze[c].facing)
		rotateDirs := rotate(m, maze[c])
		for _, coord := range rotateDirs {
			new, ok := maze[coord]
			if ok && new.value != '#' {
				if !yield(coord) {
					return
				}
			}
		}
	}
}

func nodeDistance(c, d utils.Coord) (dist float64) {
	return float64(utils.Abs(c.Row-d.Row) + utils.Abs(c.Col-d.Col))
}

func nodeCost(c, d utils.Coord) float64 {
	if !maze[c].facing.Equals(maze[d].facing) {
		return 1000
	}
	return 1
}

func PartOne(input []string) int {
	result := 0
	rows := len(input)
	cols := len(input[0])
	fmt.Println(rows, cols)
	facing := utils.Coord{Row: 0, Col: 1}
	var start, end utils.Coord

	for row, line := range input {
		for col, char := range line {
			coord := utils.Coord{Row: row, Col: col}
			node := Node{location: coord, facing: facing, value: char}
			maze[coord] = node
			if char == 'S' {
				start = coord
			} else if char == 'E' {
				end = coord
			}
		}
	}

	if *debug {
		fmt.Println(start, end)
		out := ""
		for row := range rows {
			for col := range cols {
				out += string(maze[utils.Coord{Row: row, Col: col}].value)
			}
			out += "\n"
		}
		fmt.Println(out)
		fmt.Println(start, end)
	}

	//path := astar.FindPath(maze, start, end, nodeCost, nodeDistance)
	path := astar.FindPath(maze, start, end, nodeCost, nodeCost)

	if *debug {
		fmt.Println(path)
		fmt.Println(len(path))
	}

	for row := range rows {
		pathed := ""
		for col := range cols {
			coord := utils.Coord{Row: row, Col: col}
			if slices.Contains(path, coord) {
				pathed += "O"
			} else {
				pathed += string(maze[coord].value)
			}
		}
		fmt.Println(pathed)
	}
	for i := 1; i < len(path); i++ {
		result++
		if newFacing := path[i].Sub(path[i-1]); newFacing != facing {
			result += 1000
			facing = newFacing
		}
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {

	partOne := flag.Bool("1", false, "Run part one")
	partTwo := flag.Bool("2", false, "Run part two")
	flag.Parse()

	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	if *partOne {
		partOneResult := PartOne(input)
		fmt.Println("Part One Result:", partOneResult)
		os.Exit(0)
	}

	if *partTwo {
		partTwoResult := PartTwo(input)
		fmt.Println("Part Two Result:", partTwoResult)
		os.Exit(0)
	}

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
