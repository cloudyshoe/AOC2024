package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		secret := utils.Atoi(line)
		if *debug {
			fmt.Println(secret)
		}
		i := 0
		for i < 2000 {
			secret ^= secret * 64
			secret %= 16777216
			secret ^= secret / 32
			secret %= 16777216
			secret ^= secret * 2048
			secret = secret % 16777216
			i++
		}
		if *debug {
			fmt.Println(secret)
		}
		result += secret
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
	input := strings.Split(string(inputFile), "\n")

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
