package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	result := 0

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	matches := re.FindAllStringSubmatch(input[0], -1)

	for i := range matches {
		tmp1, _ := strconv.Atoi(matches[i][1])
		tmp2, _ := strconv.Atoi(matches[i][2])

		result += tmp1 * tmp2
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(([0-9]+),([0-9]+)\)`)

	indexes := re.FindAllStringIndex(input[0], -1)

	enabled := true

	for _, index := range indexes {
		instruction := input[0][index[0] : index[0]+3]
		if instruction == "do(" {
			enabled = true
		} else if instruction == "don" {
			enabled = false
		} else if instruction == "mul" && enabled {
			tmp := strings.Split(input[0][index[0]+4:index[1]-1], ",")
			tmp1, _ := strconv.Atoi(tmp[0])
			tmp2, _ := strconv.Atoi(tmp[1])
			result += tmp1 * tmp2
		}
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
}
