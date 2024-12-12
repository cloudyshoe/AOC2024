package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

var visited = make(utils.HashGrid[struct{}])

func Neighbors(coords utils.Coords, char rune, garden utils.HashGrid[rune]) []utils.Coords {
	neighbors := make([]utils.Coords, 0)
	dirs := []string{"n", "e", "w", "s"}

	for _, dir := range dirs {
		neighbor := garden.Dir(coords, dir)
		if neighbor.Exists && neighbor.Val == char {
			neighbors = append(neighbors, neighbor.Point)
		}
	}

	return neighbors
}

func Regionize(coords utils.Coords, char rune, garden utils.HashGrid[rune], regions utils.HashGrid[int], regionNum int, visited utils.HashGrid[struct{}]) {
	visited[coords] = struct{}{}
	regions[coords] = regionNum
	neighbors := Neighbors(coords, char, garden)
	for _, neighbor := range neighbors {
		if _, ok := visited[neighbor]; !ok {
			Regionize(neighbor, char, garden, regions, regionNum, visited)
		}
	}
}

func PartOne(input []string) int {
	result := 0
	cells := len(input) * len(input[0])
	garden := make(utils.HashGrid[rune], cells)
	neighbors := make(utils.HashGrid[int], cells)
	regions := make(utils.HashGrid[int], cells)
	visited := make(utils.HashGrid[struct{}])
	regionNum := 1
	rows := len(input)
	cols := len(input[0])
	//areas := make([]int, 0, 100)
	//perimeters := make([]int, 0, 100)

	for row, line := range input {
		for col, char := range line {
			coords := utils.Coords{Row: row, Col: col}
			garden[coords] = char
		}
	}

	for k, v := range garden {
		neighbors[k] = len(Neighbors(k, v, garden))
		if _, ok := visited[k]; !ok {
			Regionize(k, v, garden, regions, regionNum, visited)
			regionNum++
		}
	}

	//for row := range len(input) {
	//	for col := range len(input[0]) {
	//		fmt.Print(regions[utils.Coords{Row: row, Col: col}])
	//	}
	//	fmt.Print("\n")
	//}

	for i := 1; i < regionNum; i++ {
		area, perimeter := 0, 0
		for row := range rows {
			for col := range cols {
				c := utils.Coords{Row: row, Col: col}
				v := regions[c]
				if v == i {
					area++
					perimeter += 4 - neighbors[c]
				}
			}
		}
		result += area * perimeter
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
