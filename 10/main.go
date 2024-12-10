package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func floodFill(grid utils.IntGrid, point utils.Point, origPoint utils.Point) {
	dirs := []string{"n", "e", "s", "w"}
	elevation := grid[point]

	for _, dir := range dirs {
		nPoint := grid.Dir(point, dir)
		if nPoint.Exists && nPoint.Val-elevation == 1 && nPoint.Val == 9 {
			if !slices.Contains(reachableNines[origPoint], nPoint.Point) {
				reachableNines[origPoint] = append(reachableNines[origPoint], nPoint.Point)
			}
		} else if nPoint.Exists && nPoint.Val-elevation == 1 {
			floodFill(grid, nPoint.Point, origPoint)
		}
	}
}

var reachableNines = make(map[utils.Point][]utils.Point)

func PartOne(input []string) int {
	result := 0

	grid := make(utils.IntGrid)
	rows := len(input)
	cols := len(input[0])

	for x, line := range input {
		for y, char := range strings.Split(line, "") {
			grid[utils.Point{X: x, Y: y}] = utils.Atoi(char)
		}
	}

	for x := 0; x < cols; x++ {
		line := ""
		for y := 0; y < rows; y++ {
			point := utils.Point{X: x, Y: y}
			elevation := grid[point]
			line += strconv.Itoa(elevation)
			if elevation == 0 {
				floodFill(grid, point, point)
			}
		}
	}

	for _, v := range reachableNines {
		result += len(v)
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
