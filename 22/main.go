package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func next(secret int) int {
	secret ^= secret * 64
	secret %= 16777216
	secret ^= secret / 32
	secret %= 16777216
	secret ^= secret * 2048
	secret = secret % 16777216
	return secret
}

func next2k(secret int, c chan int) {
	for range 2000 {
		secret = next(secret)
	}
	c <- secret
}

func PartOne(input []string) int {
	result := 0

	c := make(chan int)
	for _, line := range input {
		secret := utils.Atoi(line)
		go next2k(secret, c)
	}

	for range input {
		result += <-c
	}

	return result
}

func PartTwo(input []string) int {
	result := 0
	allChanges := make([]map[[4]int]int, len(input))
	combinedChanges := make(map[[4]int]int)
	c := make(chan map[[4]int]int, len(input))

	for _, line := range input {
		go func() {
			changeSeq := [4]int{}
			secret := utils.Atoi(line)
			changes := make(map[[4]int]int)
			prevBananas := secret % 10
			for i := 0; i < 2000; i++ {
				secret = next(secret)
				bananas := secret % 10
				diff := bananas - prevBananas
				if i < 3 {
					changeSeq[i] = diff
				} else if i == 3 {
					changeSeq[i] = diff
					changes[changeSeq] = bananas
				} else {
					changeSeq = [4]int{changeSeq[1], changeSeq[2], changeSeq[3], diff}
					if _, ok := changes[changeSeq]; !ok {
						changes[changeSeq] = bananas
					}
				}
				prevBananas = bananas
			}
			c <- changes
		}()
	}

	for range input {
		allChanges = append(allChanges, <-c)
		for k, v := range allChanges[len(allChanges)-1] {
			combinedChanges[k] += v
		}
	}

	for k, v := range combinedChanges {
		if v > result {
			if *debug {
				fmt.Println(k, v)
			}
			result = v
		}
	}

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
