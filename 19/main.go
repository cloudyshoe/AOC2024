package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func connectSegments(idx, patternLen int, usefulTowels map[int][]int) bool {

	if idx == patternLen {
		return true
	}

	segments, ok := usefulTowels[idx]

	if !ok {
		return false
	}

	for _, segmentLen := range segments {
		connected := connectSegments(idx+segmentLen, patternLen, usefulTowels)
		if connected {
			return true
		}
	}

	return false
}

func buildPattern(pattern string, towels []string) bool {
	usefulTowels := make(map[int][]int)
	patternLen := len(pattern)

	for i := 0; i < len(pattern); i++ {
		for j := i + 1; j <= len(pattern); j++ {
			segment := pattern[i:j]
			if slices.Contains(towels, segment) {
				usefulTowels[i] = append(usefulTowels[i], len(segment))
			}
		}
	}

	_, ok := usefulTowels[0]
	if !ok {
		return false
	}

	connected := connectSegments(0, patternLen, usefulTowels)

	if connected {
		return true
	}

	if *debug {
		fmt.Println(pattern, usefulTowels)
	}

	return false
}

type Header struct {
	idx        int
	patternLen int
}

var cache = make(map[Header]int)

func connectSegmentsCount(idx, patternLen int, usefulTowels map[int][]int) int {
	result := 0

	key := Header{idx: idx, patternLen: patternLen}

	if idx == patternLen {
		cache[key] = 1
		return 1
	}

	stored, ok := cache[key]

	if ok {
		return stored
	}

	segments, ok := usefulTowels[idx]

	if ok {
		for _, segmentLen := range segments {
			result += connectSegmentsCount(idx+segmentLen, patternLen, usefulTowels)
		}
	}

	cache[key] = result

	return result
}

func combinations(pattern string, towels []string) int {
	usefulTowels := make(map[int][]int)
	patternLen := len(pattern)
	result := 0

	for i := 0; i < len(pattern); i++ {
		for j := i + 1; j <= len(pattern); j++ {
			segment := pattern[i:j]
			if slices.Contains(towels, segment) {
				usefulTowels[i] = append(usefulTowels[i], len(segment))
			}
		}
	}

	cache = make(map[Header]int)
	result = connectSegmentsCount(0, patternLen, usefulTowels)

	if *debug {
		fmt.Println(pattern, usefulTowels)
	}

	return result
}

func PartOne(input []string) int {
	result := 0

	towels := strings.Split(input[0], ", ")
	patterns := strings.Split(input[1], "\n")

	if *debug {
		fmt.Println(patterns)
	}

	for _, pattern := range patterns {
		if buildPattern(pattern, towels) {
			result += 1
		}
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	towels := strings.Split(input[0], ", ")
	patterns := strings.Split(input[1], "\n")

	if *debug {
		fmt.Println(patterns)
	}

	for _, pattern := range patterns {
		result += combinations(pattern, towels)
	}

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
