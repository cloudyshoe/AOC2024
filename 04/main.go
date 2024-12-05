package main

import (
	"fmt"
	"os"
	"strings"
)

func PartOne(input []string) int {
	result := 0
	rows := len(input)
	cols := len(input[0])

	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'X' {
				if j-3 >= 0 {
					if input[i][j-3:j] == "SAM" {
						result++
					}
				}
				if j+4 <= cols {
					if input[i][j+1:j+4] == "MAS" {
						result++
					}
				}
				if i-3 >= 0 {
					if input[i-1][j] == 'M' &&
						input[i-2][j] == 'A' &&
						input[i-3][j] == 'S' {
						result++
					}
				}
				if i+3 < rows {
					if input[i+1][j] == 'M' &&
						input[i+2][j] == 'A' &&
						input[i+3][j] == 'S' {
						result++
					}
				}
				if i-3 >= 0 && j-3 >= 0 {
					if input[i-1][j-1] == 'M' &&
						input[i-2][j-2] == 'A' &&
						input[i-3][j-3] == 'S' {
						result++
					}
				}
				if i-3 >= 0 && j+3 < cols {
					if input[i-1][j+1] == 'M' &&
						input[i-2][j+2] == 'A' &&
						input[i-3][j+3] == 'S' {
						result++
					}
				}
				if i+3 < rows && j-3 >= 0 {
					if input[i+1][j-1] == 'M' &&
						input[i+2][j-2] == 'A' &&
						input[i+3][j-3] == 'S' {
						result++
					}
				}
				if i+3 < rows && j+3 < cols {
					if input[i+1][j+1] == 'M' &&
						input[i+2][j+2] == 'A' &&
						input[i+3][j+3] == 'S' {
						result++
					}
				}
			}
		}
	}
	return result
}

func PartTwo(input []string) int {
	result := 0
	rows := len(input)
	cols := len(input[0])

	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'A' {
				if i+1 < rows && i-1 >= 0 &&
					j+1 < cols && j-1 >= 0 {
					if input[i-1][j-1] == 'M' &&
						input[i-1][j+1] == 'M' &&
						input[i+1][j-1] == 'S' &&
						input[i+1][j+1] == 'S' {
						result++
					}
					if input[i-1][j-1] == 'S' &&
						input[i-1][j+1] == 'S' &&
						input[i+1][j-1] == 'M' &&
						input[i+1][j+1] == 'M' {
						result++
					}
					if input[i-1][j-1] == 'M' &&
						input[i-1][j+1] == 'S' &&
						input[i+1][j-1] == 'M' &&
						input[i+1][j+1] == 'S' {
						result++
					}
					if input[i-1][j-1] == 'S' &&
						input[i-1][j+1] == 'M' &&
						input[i+1][j-1] == 'S' &&
						input[i+1][j+1] == 'M' {
						result++
					}
				}
			}
		}
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
