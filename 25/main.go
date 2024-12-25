package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func PartOne(input []string) int {
	result := 0
	locks := make([][5]int, 0, len(input)/2)
	keys := make([][5]int, 0, len(input)/2)

	for _, schematics := range input {
		lines := strings.Split(schematics, "\n")
		var heights [5]int
		if lines[0][0] == '#' {
			for row := 1; row < len(lines); row++ {
				for col := 0; col < len(lines[row]); col++ {
					if lines[row][col] == '#' {
						heights[col]++
					}
				}
			}
			locks = append(locks, heights)
		} else {
			for row := 0; row < len(lines)-1; row++ {
				for col := 0; col < len(lines[row]); col++ {
					if lines[row][col] == '#' {
						heights[col]++
					}
				}
			}
			keys = append(keys, heights)
		}
	}

	if *debug {
		fmt.Println(keys)
		fmt.Println(locks)
	}

	for _, key := range keys {
		for _, lock := range locks {
			fit := true
			for i := range key {
				if key[i]+lock[i] > 5 {
					fit = false
					break
				}
			}
			if fit {
				result += 1
			}
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
