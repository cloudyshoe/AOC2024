package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Gate struct {
	left     string
	gate     string
	right    string
	wire     string
	resolved bool
}

func runGate(left, right int, gate string) int {
	switch gate {
	case "AND":
		return left & right
	case "OR":
		return left | right
	case "XOR":
		return left ^ right
	default:
		panic("invalid op")
	}
}

func PartOne(input []string) int {
	result := 0
	wires := make(map[string]int, len(input))
	gates := make([]Gate, 0, len(input))

	for _, line := range strings.Split(input[0], "\n") {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = utils.Atoi(parts[1])
	}
	for _, line := range strings.Split(input[1], "\n") {
		parts := strings.Fields(line)
		gate := Gate{
			left:     parts[0],
			gate:     parts[1],
			right:    parts[2],
			wire:     parts[4],
			resolved: false,
		}
		gates = append(gates, gate)
	}

	unresolved := true
	for unresolved {
		resolved := 0
		for i, gate := range gates {
			if gate.resolved {
				resolved++
				continue
			}
			left, okL := wires[gate.left]
			right, okR := wires[gate.right]
			if okL && okR && !gate.resolved {
				wires[gate.wire] = runGate(left, right, gate.gate)
				gate.resolved = true
				gates[i] = gate
				break
			}
		}
		if resolved == len(gates) {
			unresolved = false
		}
	}

	tmp := ""
	for i := 99; i >= 0; i-- {
		key := fmt.Sprintf("z%02d", i)
		if num, ok := wires[key]; ok {
			tmp += strconv.Itoa(num)
		}
	}

	result64, _ := strconv.ParseInt(tmp, 2, 64)
	result = int(result64)

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
