package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Pins struct {
	lockOrKey string
	heights   [5]int
}

func countPins(schematic string) (pins Pins) {
	lines := strings.Split(schematic, "\n")
	var startRow, endRow int
	if lines[0][0] == '#' {
		pins.lockOrKey = "lock"
		startRow = 1
		endRow = len(lines)
	} else {
		pins.lockOrKey = "key"
		startRow = 0
		endRow = len(lines) - 1
	}
	for row := startRow; row < endRow; row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == '#' {
				pins.heights[col]++
			}
		}
	}
	return
}

func checkFit(key, lock [5]int) int {
	fit := true
	for i := range key {
		if key[i]+lock[i] > 5 {
			fit = false
			break
		}
	}
	if fit {
		return 1
	} else {
		return 0
	}

}

func PartOne(input []string) int {
	result := 0
	locks := make([][5]int, 0, len(input)/2)
	keys := make([][5]int, 0, len(input)/2)

	for _, schematic := range input {
		pins := countPins(schematic)
		if pins.lockOrKey == "lock" {
			locks = append(locks, pins.heights)
		} else {
			keys = append(keys, pins.heights)
		}
	}

	if *debug {
		fmt.Println(keys)
		fmt.Println(locks)
	}

	for _, key := range keys {
		for _, lock := range locks {
			result += checkFit(key, lock)
		}
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
