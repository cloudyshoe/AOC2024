package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func PartOne(input []string) int {
	result := 0
	//var blocks []int
	//var freeSpace []int
	//var usedSpace []int

	size := 0

	tmp := make([]int, len(input[0]))
	for i, char := range input[0] {
		num := utils.Atoi(char)
		tmp[i] = num
		size += num
	}

	blocks := make([]int, 0, size)
	freeSpace := make([]int, 0, size)
	usedSpace := make([]int, 0, size)

	for i, num := range tmp {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				blocks = append(blocks, i/2)
				usedSpace = append(usedSpace, len(blocks)-1)
			}
		} else {
			for j := 0; j < num; j++ {
				blocks = append(blocks, -1)
				freeSpace = append(freeSpace, len(blocks)-1)
			}
		}
	}

	for i, num := range freeSpace {
		usedIndex := usedSpace[len(usedSpace)-1-i]
		if num > usedIndex {
			break
		}
		if blocks[num] == -1 && blocks[usedIndex] != -1 {
			blocks[num] = blocks[usedIndex]
			blocks[usedIndex] = -1
		}
	}

	for i, num := range blocks {
		if num == -1 {
			break
		}
		result += i * num
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	type Block struct {
		start int
		len   int
	}

	//var blocks []int
	//var freeSpace []Block
	//var usedSpace []Block

	size := 0

	tmp := make([]int, len(input[0]))
	for i, char := range input[0] {
		num := utils.Atoi(char)
		tmp[i] = num
		size += num
	}

	blocks := make([]int, 0, size)
	freeSpace := make([]Block, 0, size)
	usedSpace := make([]Block, 0, size)

	for i, num := range tmp {
		if i%2 == 0 {
			usedSpace = append(usedSpace, Block{start: len(blocks), len: num})
			for j := 0; j < num; j++ {
				blocks = append(blocks, i/2)
			}
		} else {
			freeSpace = append(freeSpace, Block{start: len(blocks), len: num})
			for j := 0; j < num; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	for i := range usedSpace {
		blockId := len(usedSpace) - 1 - i
		block := usedSpace[blockId]
		for j, freeBlock := range freeSpace {
			if freeBlock.len >= block.len && freeBlock.start < block.start {
				for k := freeBlock.start; k < freeBlock.start+block.len; k++ {
					blocks[k] = blockId
				}
				for k := block.start; k < block.start+block.len; k++ {
					blocks[k] = -1
				}
				freeSpace[j].len -= block.len
				freeSpace[j].start += block.len
				freeSpace = append(freeSpace, block)
				break
			}
		}
	}

	for i, num := range blocks {
		if num != -1 {
			result += i * num
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
