package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func PartOne(input []string) int {
	result := 0

	connections := make(map[string][]string, len(input))
	tees := make([]string, 0, len(input))

	for _, line := range input {
		computers := strings.Split(line, "-")
		connections[computers[0]] = append(connections[computers[0]], computers[1])
		connections[computers[1]] = append(connections[computers[1]], computers[0])
	}

	for computer, connectedComputers := range connections {
		if computer[0] == 't' {
			maybes := make([]string, 0, 4)
			for i, cpu1 := range connectedComputers[0 : len(connectedComputers)-1] {
				for _, cpu2 := range connectedComputers[i:] {
					if slices.Contains(connections[cpu1], cpu2) {
						tmp := []string{computer, cpu1, cpu2}
						slices.Sort(tmp)
						group := ""
						for _, cpu := range tmp {
							group += cpu
						}
						maybes = append(maybes, group)
					}
				}
			}
			for _, maybe := range maybes {
				if !slices.Contains(tees, maybe) {
					tees = append(tees, maybe)
				}
			}
		}
	}

	if *debug {
		slices.Sort(tees)
		fmt.Println(tees)
	}

	result = len(tees)
	return result
}

func PartTwo(input []string) int {
	result := 0
	connections := make(map[string][]string, len(input))
	//password := make([]string, 0, 100)

	for _, line := range input {
		computers := strings.Split(line, "-")
		connections[computers[0]] = append(connections[computers[0]], computers[1])
		connections[computers[1]] = append(connections[computers[1]], computers[0])
	}

	fmt.Println(connections)

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
