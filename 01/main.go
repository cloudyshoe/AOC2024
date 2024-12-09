package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"aoc/utils"
)

func PartOne(input []string) int {
	result := 0
	listLen := len(input)
	left := make([]int, listLen)
	right := make([]int, listLen)

	for i, line := range input {
		tmp := strings.Fields(line)
		left[i], _ = strconv.Atoi(tmp[0])
		right[i], _ = strconv.Atoi(tmp[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		result += utils.Abs(left[i] - right[i])
	}

	return result
}

func PartTwo(input []string) int {
	result := 0
	listLen := len(input)
	counts := make(map[int]int, listLen)

	left := make([]int, listLen)

	for i, line := range input {
		tmp := strings.Fields(line)
		left[i], _ = strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		counts[right]++
	}

	for _, num := range left {
		result += num * counts[num]
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
