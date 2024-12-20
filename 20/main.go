package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"iter"
	"os"
	"strings"

	"github.com/fzipp/astar"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Maze struct {
	data utils.HashGrid[rune]
}

func PrintMaze(maze utils.HashGrid[rune], rows, cols int) {
	out := ""
	for row := range rows {
		for col := range cols {
			coord := utils.Coord{Row: row, Col: col}
			out += string(maze[coord])
		}
		out += "\n"
	}
	fmt.Print(out)
}

func distance(a, b utils.Coord) (result float64) {
	result = 1
	return
}

func (m Maze) Neighbours(a utils.Coord) iter.Seq[utils.Coord] {
	return func(yield func(utils.Coord) bool) {
		dirs := []string{"n", "e", "s", "w"}
		for _, dir := range dirs {
			coord := a.Add(utils.GridDirs[dir])
			var tmp rune
			var ok bool
			if a.Equals(cheatStart) || a.Equals(cheatEnd) {
				tmp, ok = cheatMaze.data[coord]
			} else {
				tmp, ok = maze.data[coord]
			}
			if ok && tmp != '#' {
				if !yield(coord) {
					return
				}
			}
		}
	}
}

var maze = Maze{}
var cheatMaze = Maze{}
var cheatStart, cheatEnd utils.Coord

func PartOne(input []string) int {
	result := 0
	rows := len(input)
	cols := len(input[0])
	var start, end utils.Coord

	maze.data = make(utils.HashGrid[rune], rows*cols)
	cheatMaze.data = make(utils.HashGrid[rune], rows*cols)

	for row, line := range input {
		if row == 0 || row == rows-1 {
			continue
		}
		for col, char := range line {
			if col == 0 || col == cols-1 {
				continue
			}
			coord := utils.Coord{Row: row, Col: col}
			maze.data[coord] = char
			if char == '#' {
				cheatMaze.data[coord] = '.'
			} else {
				cheatMaze.data[coord] = char
			}
			if char == 'S' {
				start = coord
			}
			if char == 'E' {
				end = coord
			}
		}
	}

	if *debug {
		PrintMaze(maze.data, rows, cols)
		PrintMaze(cheatMaze.data, rows, cols)
	}

	boringPath := astar.FindPath(maze, start, end, distance, distance)
	boringLen := len(boringPath)

	dirs := []string{"n", "e", "s", "w"}
	for i, cell := range boringPath {
		for _, dir := range dirs {
			tmp := cell.Add(utils.GridDirs[dir])
			for _, dir2 := range dirs {
				tmp2 := tmp.Add(utils.GridDirs[dir2])
				for j := i + 101; j < boringLen; j++ {
					if tmp2.Equals(boringPath[j]) {
						result++
						break
					}
				}
			}
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
