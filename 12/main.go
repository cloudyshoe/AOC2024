package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

var visited = make(utils.HashGrid[struct{}])

func Neighbors(coord utils.Coord, char rune, garden utils.HashGrid[rune]) []utils.Coord {
	neighbors := make([]utils.Coord, 0)
	dirs := []string{"n", "e", "w", "s"}

	for _, dir := range dirs {
		neighbor := garden.Dir(coord, dir)
		if neighbor.Exists && neighbor.Val == char {
			neighbors = append(neighbors, neighbor.Point)
		}
	}

	return neighbors
}

func Regionize(coord utils.Coord, char rune, garden utils.HashGrid[rune], regions utils.HashGrid[int], regionNum int, visited utils.HashGrid[struct{}]) {
	visited[coord] = struct{}{}
	regions[coord] = regionNum
	neighbors := Neighbors(coord, char, garden)
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
	visited := make(utils.HashGrid[struct{}], cells)
	areas := make(map[int]int)
	perimeters := make(map[int]int)
	regionNum := 0
	rows := len(input)
	cols := len(input[0])

	for row, line := range input {
		for col, char := range line {
			coord := utils.Coord{Row: row, Col: col}
			garden[coord] = char
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
	//		c := utils.Coord{Row: row, Col: col}
	//		if num := regions[c]; num == 1 {
	//			fmt.Print(neighbors[c])
	//		} else {
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Print("\n")
	//}

	for row := range rows {
		for col := range cols {
			c := utils.Coord{Row: row, Col: col}
			r := regions[c]
			n := neighbors[c]
			areas[r]++
			perimeters[r] += 4 - n
		}
	}

	for i := 0; i < regionNum; i++ {
		result += areas[i] * perimeters[i]
	}

	return result
}

func Sides(coord utils.Coord, regions utils.HashGrid[int]) int {
	sides := 0

	return sides
}

func PartTwo(input []string) int {
	result := 0
	cells := len(input) * len(input[0])
	garden := make(utils.HashGrid[rune], cells)
	neighbors := make(utils.HashGrid[int], cells)
	regions := make(utils.HashGrid[int], cells)
	regionStart := make([]utils.Coord, 0, 100)
	visited := make(utils.HashGrid[struct{}], cells)
	areas := make(map[int]int)
	regionNum := 0
	rows := len(input)
	cols := len(input[0])

	for row, line := range input {
		for col, char := range line {
			coord := utils.Coord{Row: row, Col: col}
			garden[coord] = char
		}
	}

	for k, v := range garden {
		neighbors[k] = len(Neighbors(k, v, garden))
		if _, ok := visited[k]; !ok {
			regionStart = append(regionStart, k)
			Regionize(k, v, garden, regions, regionNum, visited)
			regionNum++
		}
	}

	//for row := range len(input) {
	//	for col := range len(input[0]) {
	//		c := utils.Coords{Row: row, Col: col}
	//		if num := regions[c]; num == 0 {
	//			fmt.Print(neighbors[c])
	//		} else {
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Print("\n")
	//}

	for row := range rows {
		for col := range cols {
			c := utils.Coord{Row: row, Col: col}
			r := regions[c]
			areas[r]++
		}
	}

	for i := 0; i < regionNum; i++ {
		result += areas[i] * Sides(regionStart[i], regions)
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
