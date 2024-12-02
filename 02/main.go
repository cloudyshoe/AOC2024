package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func safetyCheck(levels []string) int {
	levelsInt := make([]int, len(levels))

	for i, level := range levels {
		levelsInt[i], _ = strconv.Atoi(level)
	}

	firstSign := sign(levelsInt[1] - levelsInt[0])

	for i := 1; i < len(levelsInt); i++ {
		if sign(levelsInt[i]-levelsInt[i-1]) != firstSign || abs(levelsInt[i]-levelsInt[i-1]) > 3 {
			return 0
		}
	}

	return 1
}

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		levels := strings.Split(line, " ")
		result += safetyCheck(levels)
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	for _, line := range input {
		levels := strings.Split(line, " ")
		result += safetyCheck(levels)
	}

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(input)
	fmt.Printf("Part One Result: %d\n", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
