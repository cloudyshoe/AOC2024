package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func PartOne(input []string) int {
	result := 0

	nodes := make(utils.HashGrid[rune], len(input)*len(input[0]))
	antinodes := make(utils.HashGrid[rune], len(input)*len(input[0]))
	frequencies := make(map[rune][]utils.Point)

	for y, line := range input {
		for x, char := range line {
			point := utils.Point{X: x, Y: y}
			nodes[point] = char
			antinodes[point] = '.'
			if char != '.' {
				frequencies[char] = append(frequencies[char], point)
			}
		}
	}

	bounds := utils.Bounds{Min: utils.Point{X: 0, Y: 0}, Max: utils.Point{X: len(input[0]), Y: len(input)}}

	for _, points := range frequencies {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				//fmt.Println(string(frequency), points[j], points[i].Sub(points[j]))
				//fmt.Println(string(frequency), points[i], points[j].Sub(points[i]))
				iNewPoint := points[i].Add(points[i].Sub(points[j]))
				jNewPoint := points[j].Add(points[j].Sub(points[i]))
				if iNewPoint.In(bounds) {
					if antinodes[iNewPoint] != '#' {
						antinodes[iNewPoint] = '#'
						result++
					}
				}
				if jNewPoint.In(bounds) {
					if antinodes[jNewPoint] != '#' {
						antinodes[jNewPoint] = '#'
						result++
					}
				}

			}
		}
	}

	/*
		nodeout := ""
		antinodeout := ""
		for y := range input {
			nodeline := ""
			antinodeline := ""
			for x := range input {
				point := utils.Point{X: x, Y: y}
				nodeline += string(nodes[point])
				antinodeline += string(antinodes[point])
			}
			nodeout += nodeline + "\n"
			antinodeout += antinodeline + "\n"
		}

		fmt.Println(nodeout)
		fmt.Println(antinodeout)
		fmt.Println(frequencies)
	*/

	return result
}

func PartTwo(input []string) int {
	result := 0

	nodes := make(map[utils.Point]rune, len(input)*len(input[0]))
	antinodes := make(map[utils.Point]rune, len(input)*len(input[0]))
	frequencies := make(map[rune][]utils.Point)

	for y, line := range input {
		for x, char := range line {
			point := utils.Point{X: x, Y: y}
			nodes[point] = char
			antinodes[point] = '.'
			if char != '.' {
				frequencies[char] = append(frequencies[char], point)
				antinodes[point] = '#'
				result++
			}
		}
	}

	bounds := utils.Bounds{Min: utils.Point{X: 0, Y: 0}, Max: utils.Point{X: len(input[0]), Y: len(input)}}

	for _, points := range frequencies {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				//fmt.Println(string(frequency), points[j], points[i].Sub(points[j]))
				//fmt.Println(string(frequency), points[i], points[j].Sub(points[i]))
				iDist := points[i].Sub(points[j])
				for iNewPoint := points[i].Add(iDist); iNewPoint.In(bounds); iNewPoint = iNewPoint.Add(iDist) {
					if antinodes[iNewPoint] != '#' {
						antinodes[iNewPoint] = '#'
						result++
					}

				}
				jDist := points[j].Sub(points[i])
				for jNewPoint := points[j].Add(jDist); jNewPoint.In(bounds); jNewPoint = jNewPoint.Add(jDist) {
					if antinodes[jNewPoint] != '#' {
						antinodes[jNewPoint] = '#'
						result++
					}

				}

			}
		}
	}

	/*
		nodeout := ""
		antinodeout := ""
		for y := range input {
			nodeline := ""
			antinodeline := ""
			for x := range input {
				point := utils.Point{X: x, Y: y}
				nodeline += string(nodes[point])
				antinodeline += string(antinodes[point])
			}
			nodeout += nodeline + "\n"
			antinodeout += antinodeline + "\n"
		}

		fmt.Println(nodeout)
		fmt.Println(antinodeout)
		fmt.Println(frequencies)
	*/

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
