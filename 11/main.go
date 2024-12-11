package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	result := 0

	/*
	   If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
	   If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
	       The left half of the digits are engraved on the new left stone,
	       and the right half of the digits are engraved on the new right stone.
	       (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
	   If none of the other rules apply, the stone is replaced by a new stone;
	       the old stone's number multiplied by 2024 is engraved on the new stone.
	*/

	stones := strings.Fields(input[0])

	for i := 0; i < 25; i++ {
		newStones := make([]string, 0, len(stones))
		for _, stone := range stones {
			switch {
			case stone == "0":
				newStones = append(newStones, "1")
			case len(stone)%2 == 0:
				stone1 := strconv.Itoa((utils.Atoi(stone[0 : len(stone)/2])))
				stone2 := strconv.Itoa((utils.Atoi(stone[len(stone)/2:])))
				newStones = append(newStones, stone1, stone2)
			default:
				newStones = append(newStones, strconv.Itoa(utils.Atoi(stone)*2024))
			}
		}
		stones = newStones
	}

	result = len(stones)

	return result
}

type cacheHeader struct {
	stone int
	depth int
}

var memo = make(map[cacheHeader]int)
var maxDepth = 75

func resetCache() {
	memo = make(map[cacheHeader]int)
}

func processStone(stone int, depth int) int {

	key := cacheHeader{stone: stone, depth: depth}

	if depth == maxDepth {
		memo[key] = 1
		return 1
	}

	stored, ok := memo[key]
	if ok {
		return stored
	}

	result := 0
	if !ok {
		if stone == 0 {
			result += processStone(1, depth+1)
		} else if digits := len(strconv.Itoa(stone)); digits%2 == 0 {
			stone1 := stone / int(math.Pow10(digits/2))
			stone2 := stone % int(math.Pow10(digits/2))
			result += processStone(stone1, depth+1)
			result += processStone(stone2, depth+1)
		} else {
			result += processStone(stone*2024, depth+1)
		}
	}

	memo[key] = result

	return result
}

func PartTwo(input []string) int {
	result := 0
	stones := strings.Fields(input[0])

	for _, stone := range stones {
		stInt := utils.Atoi(stone)
		result += processStone(stInt, 0)
	}

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)

	maxDepth = 25
	resetCache()

	partTwoResult = PartTwo(input)
	fmt.Println("Part One Part Two Style Result:", partTwoResult)
}
