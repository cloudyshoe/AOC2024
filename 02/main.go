package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

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
		if sign(levels[i]-levels[i-1]) != firstSign || utils.Abs(levels[i]-levels[i-1]) > 3 {
			return 0
		}
	}

	return 1
}

func safetyCheckPartTwo(levels []int) int {

	if safetyCheck(levels) == 1 {
		return 1
	}

	for i := range levels {
		newLevels := append([]int{}, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		safe := safetyCheck(newLevels)
		if safe == 1 {
			return 1
		}
	}

	return 0
}

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		levels := strings.Fields(line)
		levelsInt := make([]int, len(levels))

		for i, level := range levels {
			levelsInt[i] = utils.Atoi(level)
		}

		result += safetyCheck(levelsInt)
	}

	return result
}

func PartTwo(input []string) int {
	//result := 0

	var result int32
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			levels := strings.Fields(line)
			levelsInt := make([]int, len(levels))

			for i, level := range levels {
				levelsInt[i] = utils.Atoi(level)
			}
			atomic.AddInt32(&result, int32(safetyCheckPartTwo(levelsInt)))
		}()
	}

	wg.Wait()

	return int(result)
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(input)
	fmt.Printf("Part One Result: %d\n", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
