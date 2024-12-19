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

type Grid struct {
	grid utils.HashGrid[string]
}

func (h Grid) Neighbours(p utils.Coord) iter.Seq[utils.Coord] {
	dirs := []string{"n", "e", "s", "w"}
	return func(yield func(utils.Coord) bool) {
		for _, dir := range dirs {
			new := h.grid.Dir(p, dir)
			if new.Exists && new.Val != "#" {
				if !yield(new.Point) {
					return
				}
			}
		}
	}
}

func PrintGrid(grid utils.HashGrid[string], rows, cols int) {
	out := ""

	for row := range rows {
		for col := range cols {
			coord := utils.Coord{Row: row, Col: col}
			out += grid[coord]
		}
		out += "\n"
	}

	fmt.Println(out)

}

func pointDist(p, q utils.Coord) float64 {
	return 1
}

func PartOne(input []string, rows, cols, bytes int) int {
	result := 0
	grid := Grid{}
	grid.grid = make(utils.HashGrid[string], rows*cols)
	start := utils.Coord{Row: 0, Col: 0}
	end := utils.Coord{Row: rows - 1, Col: cols - 1}

	for i, line := range input {
		if i == bytes {
			break
		}
		coords := strings.Split(line, ",")
		coord := utils.Coord{Row: utils.Atoi(coords[1]), Col: utils.Atoi(coords[0])}
		grid.grid[coord] = "#"
	}

	for row := range rows {
		for col := range cols {
			coord := utils.Coord{Row: row, Col: col}
			if grid.grid[coord] == "" {
				grid.grid[coord] = "."
			}
		}
	}

	if *debug {
		PrintGrid(grid.grid, rows, cols)
	}

	path := astar.FindPath(grid, start, end, pointDist, pointDist)

	result = len(path) - 1

	return result
}

func PartTwo(input []string, rows, cols int) (result string) {
	grid := Grid{}
	grid.grid = make(utils.HashGrid[string], rows*cols)
	start := utils.Coord{Row: 0, Col: 0}
	end := utils.Coord{Row: rows - 1, Col: cols - 1}

	for i, line := range input {
		coords := strings.Split(line, ",")
		coord := utils.Coord{Row: utils.Atoi(coords[1]), Col: utils.Atoi(coords[0])}
		grid.grid[coord] = "#"

		for row := range rows {
			for col := range cols {
				coord := utils.Coord{Row: row, Col: col}
				if grid.grid[coord] == "" {
					grid.grid[coord] = "."
				}
			}
		}

		if *debug {
			PrintGrid(grid.grid, rows, cols)
			fmt.Printf("Placed %d of %d\r", i+1, len(input))
		}

		path := astar.FindPath(grid, start, end, pointDist, pointDist)
		//if *debug {
		//	fmt.Println(len(path) - 1)
		//}

		if path == nil {
			result = line
			break
		}
	}

	return result
}

func main() {

	partOne := flag.Bool("1", false, "Run part one")
	partTwo := flag.Bool("2", false, "Run part two")
	flag.Parse()

	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	if *partOne {
		partOneResult := PartOne(input, 71, 71, 1024)
		fmt.Println("Part One Result:", partOneResult)
		os.Exit(0)
	}

	if *partTwo {
		partTwoResult := PartTwo(input, 71, 71)
		fmt.Println("Part Two Result:", partTwoResult)
		os.Exit(0)
	}

	partOneResult := PartOne(input, 71, 71, 1024)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input, 71, 71)
	fmt.Println("Part Two Result:", partTwoResult)
}
