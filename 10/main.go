package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func climbUpThatHill(grid utils.HashGrid[int], point utils.Coords, origPoint utils.Coords, reachableNines map[utils.Coords][]utils.Coords) {
	dirs := []string{"n", "e", "s", "w"}
	elevation := grid[point]

	for _, dir := range dirs {
		nPoint := grid.Dir(point, dir)
		if nPoint.Exists && nPoint.Val-elevation == 1 && nPoint.Val == 9 {
			if !slices.Contains(reachableNines[origPoint], nPoint.Point) {
				reachableNines[origPoint] = append(reachableNines[origPoint], nPoint.Point)
			}
		} else if nPoint.Exists && nPoint.Val-elevation == 1 {
			climbUpThatHill(grid, nPoint.Point, origPoint, reachableNines)
		}
	}
}

func pathUpThatHill(grid utils.HashGrid[int], point utils.Coords, origPoint utils.Coords, path string, paths map[utils.Coords][]string) {
	dirs := []string{"n", "e", "s", "w"}
	elevation := grid[point]

	for _, dir := range dirs {
		nPoint := grid.Dir(point, dir)
		if nPoint.Exists && nPoint.Val-elevation == 1 && nPoint.Val == 9 {
			paths[origPoint] = append(paths[origPoint], path)
		} else if nPoint.Exists && nPoint.Val-elevation == 1 {
			path += nPoint.Point.String()
			pathUpThatHill(grid, nPoint.Point, origPoint, path, paths)
		}
	}
}

func PartOne(input []string) int {
	result := 0

	grid := make(utils.HashGrid[int])
	cols := len(input)
	rows := len(input[0])
	reachableNines := make(map[utils.Coords][]utils.Coords)

	for x, line := range input {
		for y, char := range strings.Split(line, "") {
			grid[utils.Coords{Col: x, Row: y}] = utils.Atoi(char)
		}
	}

	for row := 0; row < rows; row++ {
		line := ""
		for col := 0; col < cols; col++ {
			point := utils.Coords{Col: row, Row: col}
			elevation := grid[point]
			line += strconv.Itoa(elevation)
			if elevation == 0 {
				climbUpThatHill(grid, point, point, reachableNines)
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
	grid := make(utils.HashGrid[int])
	rows := len(input)
	cols := len(input[0])
	paths := make(map[utils.Coords][]string)

	for x, line := range input {
		for y, char := range strings.Split(line, "") {
			grid[utils.Coords{Col: x, Row: y}] = utils.Atoi(char)
		}
	}

	for x := 0; x < cols; x++ {
		line := ""
		for y := 0; y < rows; y++ {
			point := utils.Coords{Col: x, Row: y}
			elevation := grid[point]
			line += strconv.Itoa(elevation)
			if elevation == 0 {
				pathUpThatHill(grid, point, point, "", paths)
			}
		}
	}

	for _, v := range paths {
		result += len(v)
	}

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
