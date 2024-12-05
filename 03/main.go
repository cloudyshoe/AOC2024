package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOneSplit(input []string) int {
	result := 0

	mulSplit := strings.Split(input[0], "mul(")
	parenSplit := make([][]string, len(mulSplit))
	for i, mul := range mulSplit {
		parenSplit[i] = strings.Split(mul, ")")
	}

	for _, out := range parenSplit {
		for _, in := range out {
			tmp := strings.Split(in,",")
			if len(tmp) == 2 {
				tmp := strings.Split(in,",")
				tmp1, err1 := strconv.Atoi(tmp[0])
				tmp2, err2 := strconv.Atoi(tmp[1])

				if err1 == nil && err2 == nil {
					result += tmp1 * tmp2
				}
			}

		}
	}

	return result
}

func PartTwoSplit(input []string) int {
	result := 0

	doSplit := strings.Split(input[0], "do()")

	for _, line := range doSplit {
		tmp := strings.Split(line,"don't()")
		result += PartOneSplit([]string{tmp[0]})
	}

	return result
}

func PartOneParser(input []string) int {
	result := 0

	type State int

	const (
		CharScan State = iota
		MulLeft
		MulRight
	)

	state := CharScan

	i := 0
	leftStr, rightStr := "", ""

	for i < len(input[0]) {

		char := string(input[0][i])

		switch state {
		case CharScan:
			if i+4 < len(input[0]) && input[0][i:i+4] == "mul(" {
				state = MulLeft
				i += 4
			} else {
				i++
			}
		case MulLeft:
			if strings.ContainsAny(char, "0123456789") {
				leftStr += char
				i++
			} else if char == "," && len(leftStr) != 0 {
				state = MulRight
				i++
			} else {
				state = CharScan
				leftStr = ""
				i++
			}
		case MulRight:
			if strings.ContainsAny(char, "0123456789") {
				rightStr += char
				i++
			} else if char == ")" && len(rightStr) != 0 {
				state = CharScan
				tmpLeft, _ := strconv.Atoi(leftStr)
				tmpRight, _ := strconv.Atoi(rightStr)
				result += tmpLeft * tmpRight
				leftStr, rightStr = "", ""
				i++
			} else {
				state = CharScan
				leftStr, rightStr = "", ""
				i++
			}
		default:
			log.Fatal("invalid state")
		}
	}

	return result
}

func PartTwoParser(input []string) int {
	result := 0

	type State int

	const (
		CharScan State = iota
		MulLeft
		MulRight
	)

	state := CharScan
	enabled := true

	i := 0
	leftStr, rightStr := "", ""

	for i < len(input[0]) {

		char := string(input[0][i])
		switch state {
		case CharScan:
			switch char {
			case "m":
				if enabled {
					if i+4 < len(input[0]) && input[0][i:i+4] == "mul(" {
						state = MulLeft
						i += 4
					} else {
						i++
					}
				} else {
					i++
				}
			case "d":
				if i+4 < len(input[0]) && input[0][i:i+4] == "do()" {
					enabled = true
					i += 4
				} else if i+7 < len(input[0]) && input[0][i:i+7] == "don't()" {
					enabled = false
					i += 7
				} else {
					i++
				}
			default:
				i++
			}
		case MulLeft:
			if strings.ContainsAny(char, "0123456789") {
				leftStr += char
				i++
			} else if char == "," && len(leftStr) != 0 {
				state = MulRight
				i++
			} else {
				state = CharScan
				leftStr = ""
				i++
			}
		case MulRight:
			if strings.ContainsAny(char, "0123456789") {
				rightStr += char
				i++
			} else if char == ")" && len(rightStr) != 0 {
				state = CharScan
				tmpLeft, _ := strconv.Atoi(leftStr)
				tmpRight, _ := strconv.Atoi(rightStr)
				result += tmpLeft * tmpRight
				leftStr, rightStr = "", ""
				i++
			} else {
				state = CharScan
				leftStr, rightStr = "", ""
				i++
			}
		default:
			log.Fatal("invalid state")
		}
	}

	return result
}

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

	partOneParserResult := PartOneParser(input)
	fmt.Println("Part One Result:", partOneParserResult)

	partOneSplitResult := PartOneSplit(input)
	fmt.Println("Part One Result:", partOneSplitResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)

	partTwoParserResult := PartTwoParser(input)
	fmt.Println("Part Two Result:", partTwoParserResult)

	partTwoSplitResult := PartTwoSplit(input)
	fmt.Println("Part Two Result:", partTwoSplitResult)
}
