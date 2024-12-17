package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Computer struct {
	A  int
	B  int
	C  int
	PC int
}

func (c Computer) ComboVal(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		log.Fatalf("Invalid operand: %d\n", operand)
	}

	return 0
}

func PartOne(input []string) string {

	registerInfo := strings.Split(input[0], "\n")
	computer := Computer{}
	tmp := strings.Split(registerInfo[0], ": ")
	computer.A = utils.Atoi(tmp[1])
	tmp = strings.Split(registerInfo[1], ": ")
	computer.B = utils.Atoi(tmp[1])
	tmp = strings.Split(registerInfo[2], ": ")
	computer.C = utils.Atoi(tmp[1])

	blah := strings.Split(input[1], "Program: ")
	program := strings.Split(blah[1], ",")
	out := ""

	for computer.PC < len(program) {
		opcode := program[computer.PC]
		operand := utils.Atoi(program[computer.PC+1])
		switch opcode {
		case "0":
			computer.A = computer.A / utils.Pow2(computer.ComboVal(operand))
			computer.PC += 2
		case "1":
			computer.B = computer.B ^ operand
			computer.PC += 2
		case "2":
			computer.B = computer.ComboVal(operand) % 8
			computer.PC += 2
		case "3":
			if computer.A == 0 {
				computer.PC += 2
			} else {
				computer.PC = operand
			}
		case "4":
			computer.B = computer.B ^ computer.C
			computer.PC += 2
		case "5":
			out += strconv.Itoa(computer.ComboVal(operand) % 8)
			out += ","
			computer.PC += 2
		case "6":
			computer.B = computer.A / utils.Pow2(computer.ComboVal(operand))
			computer.PC += 2
		case "7":
			computer.C = computer.A / utils.Pow2(computer.ComboVal(operand))
			computer.PC += 2
		default:
			log.Fatalf("Invalid Opcode: %s\n", opcode)
		}
	}

	return out[0 : len(out)-1]

	/*
	       Combo operands 0 through 3 represent literal values 0 through 3.
	       Combo operand 4 represents the value of register A.
	       Combo operand 5 represents the value of register B.
	       Combo operand 6 represents the value of register C.
	       Combo operand 7 is reserved and will not appear in valid programs.

	   	^ = bitwise XOR
	   	** = exponentiation

	   	adv opcode 0 A = A / (2**Op) combo
	   	bxl opcode 1 B = B^Op literal
	   	bst opcode 2 B = Op % 8 combo
	   	jnz opcode 3 If A != 0 then PC += Op else NOP literal
	   	bxc opcode 4 B = B^C (For legacy reasons, this instruction reads an operand but ignores it.)
	   	out opcode 5 Print(Op % 8) combo
	   	bdv opcode 6 B = A / (2**Op) combo
	   	cdv opcode 7 C = A / (2**Op) combo


	*/

}

func PartTwo(input []string) string {

	computer := Computer{}

	blah := strings.Split(input[1], "Program: ")
	program := strings.Split(blah[1], ",")
	out := ""
	outCycle := 0

	for i := 200_000_000_000_000; i < 300_000_000_000_000; i++ {
		fmt.Printf(" A: %d\r", i)
		out = ""
		outCycle = i
		computer = Computer{A: i}
		for computer.PC < len(program) {
			opcode := program[computer.PC]
			operand := utils.Atoi(program[computer.PC+1])
			switch opcode {
			case "0":
				computer.A = computer.A / utils.Pow2(computer.ComboVal(operand))
				computer.PC += 2
			case "1":
				computer.B = computer.B ^ operand
				computer.PC += 2
			case "2":
				computer.B = computer.ComboVal(operand) % 8
				computer.PC += 2
			case "3":
				if computer.A == 0 {
					computer.PC += 2
				} else {
					computer.PC = operand
				}
			case "4":
				computer.B = computer.B ^ computer.C
				computer.PC += 2
			case "5":
				if len(out) > 0 {
					out += ","
				}
				out += strconv.Itoa(computer.ComboVal(operand) % 8)
				computer.PC += 2
			case "6":
				computer.B = computer.A / utils.Pow2(computer.ComboVal(operand))
				computer.PC += 2
			case "7":
				computer.C = computer.A / utils.Pow2(computer.ComboVal(operand))
				computer.PC += 2
			default:
				log.Fatalf("Invalid Opcode: %s\n", opcode)
			}
			lenOut := len(out)
			if lenOut > len(blah[1]) {
				fmt.Println("")
				fmt.Println("Too long!")
				break
			}
			if out != blah[1][:lenOut] {
				break
			}
		}
		if out == blah[1] {
			break
		}
	}
	fmt.Println("")
	fmt.Println(outCycle)
	return out

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
