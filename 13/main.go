package main

import (
	"fmt"
	"os"
	"strings"
)

type Button struct {
	dx int
	dy int
}

type Prize struct {
	x int
	y int
}

func PartOne(input []string) int {
	result := 0
	aButtons := make([]Button, 0, len(input)/4)
	bButtons := make([]Button, 0, len(input)/4)
	prizes := make([]Prize, 0, len(input)/4)

	for _, group := range input {
		entries := strings.Split(group, "\n")
		buttonA := Button{}
		fmt.Sscanf(entries[0], "Button A: X+%d, Y+%d", &buttonA.dx, &buttonA.dy)
		aButtons = append(aButtons, buttonA)
		buttonB := Button{}
		fmt.Sscanf(entries[1], "Button B: X+%d, Y+%d", &buttonB.dx, &buttonB.dy)
		bButtons = append(bButtons, buttonB)
		prize := Prize{}
		fmt.Sscanf(entries[2], "Prize: X=%d, Y=%d", &prize.x, &prize.y)
		prizes = append(prizes, prize)
	}

	for i := 0; i < len(aButtons); i++ {
		minTokens := 1000
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				if (aButtons[i].dx*a+bButtons[i].dx*b) == prizes[i].x &&
					(aButtons[i].dy*a+bButtons[i].dy*b) == prizes[i].y {
					tokensUsed := 3*a + b
					if tokensUsed < minTokens {
						minTokens = tokensUsed
					}
				}
			}
		}
		if minTokens < 1000 {
			result += minTokens
		}
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	return result
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
