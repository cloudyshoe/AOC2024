package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func buildPattern(pattern string, towels []string, maxTowelSize int) bool {

	return true
}

func PartOne(input []string) int {
	result := 0

	towels := strings.Split(input[0], ", ")
	patterns := strings.Split(input[1], "\n")

	maxTowelSize := len(slices.MaxFunc(towels, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}))

	if *debug {
		fmt.Println(patterns)
	}

	for _, pattern := range patterns {
		if buildPattern(pattern, towels, maxTowelSize) {
			result += 1
		}
		fmt.Println(result)
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {

	partOne := flag.Bool("1", false, "Run part one")
	partTwo := flag.Bool("2", false, "Run part two")
	flag.Parse()

	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n\n")

	if *partOne {
		partOneResult := PartOne(input)
		fmt.Println("Part One Result:", partOneResult)
		os.Exit(0)
	}

	if *partTwo {
		partTwoResult := PartTwo(input)
		fmt.Println("Part Two Result:", partTwoResult)
		os.Exit(0)
	}

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
