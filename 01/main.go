package main

import (
	"fmt"
	"os"
	"strings"
)

func PartOne(input []string) int {
	result := 0

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	inputLen := len(inputFile)
	input := strings.Split(string(inputFile[:inputLen-1]), "\n")

	partOneResult := PartOne(input)
	fmt.Printf("Part One Result: %d\n", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
