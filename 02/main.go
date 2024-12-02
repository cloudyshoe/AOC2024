package main

import (
	"fmt"
	"os"
	"slices"
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

func safetyCheck(levels []int) int {

	firstSign := sign(levels[1] - levels[0])

	if firstSign == 0 {
		return 0
	}

	for i := 1; i < len(levels); i++ {
		if sign(levels[i]-levels[i-1]) != firstSign || abs(levels[i]-levels[i-1]) > 3 {
			return 0
		}
	}

	return 1
}

func safetyCheckPartTwo(levels []int) int {

	if safetyCheck(levels) == 1 {
		return 1
	}

	for i := 1; i < len(levels); i++ {
		newLevels := make([]int, len(levels))
		copy(newLevels, levels)
		newLevels = slices.Delete(newLevels, i-1, i)
		blah := safetyCheck(newLevels)
		if blah == 1 {
			return 1
		}

		if i == len(levels)-1 {
			newLevels = levels[:i]
			blah = safetyCheck(newLevels)
			if blah == 1 {
				return 1
			}
			return 0
		}
	}

	return 1
}

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		levels := strings.Split(line, " ")
		levelsInt := make([]int, len(levels))

		for i, level := range levels {
			levelsInt[i], _ = strconv.Atoi(level)
		}
		result += safetyCheck(levelsInt)
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	for _, line := range input {
		levels := strings.Split(line, " ")
		levelsInt := make([]int, len(levels))

		for i, level := range levels {
			levelsInt[i], _ = strconv.Atoi(level)
		}
		result += safetyCheckPartTwo(levelsInt)
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
