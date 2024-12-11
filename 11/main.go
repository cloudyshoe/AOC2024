package main

import (
	"aoc/utils"
	"fmt"
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
				stone1 := strconv.Itoa((utils.Atoi[string](stone[0 : len(stone)/2])))
				stone2 := strconv.Itoa((utils.Atoi[string](stone[len(stone)/2:])))
				newStones = append(newStones, stone1)
				newStones = append(newStones, stone2)
			default:
				newStones = append(newStones, strconv.Itoa(utils.Atoi(stone)*2024))
			}
		}
		stones = newStones
	}

	result = len(stones)

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
