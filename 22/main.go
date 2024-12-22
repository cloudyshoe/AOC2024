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

func next(secret int) int {
	secret ^= secret * 64
	secret %= 16777216
	secret ^= secret / 32
	secret %= 16777216
	secret ^= secret * 2048
	secret = secret % 16777216
	return secret
}

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		secret := utils.Atoi(line)
		for range 2000 {
			secret = next(secret)
		}
		result += secret
	}

	return result
}

type ChangeSeq [4]int
type WormSet map[ChangeSeq]int

func (w *WormSet) tryAdd(change ChangeSeq, bananas int) {
	worm := *w
	if _, ok := worm[change]; !ok {
		worm[change] = bananas
	}
}

func numToBananas(secret int) int {
	bananaStr := strconv.Itoa(secret)
	bananaLen := len(bananaStr)
	return utils.Atoi(bananaStr[bananaLen-1 : bananaLen])
}

func rotateChanges(changeSeq ChangeSeq, bananas, prevBananas int) ChangeSeq {
	changeSeq[0] = changeSeq[1]
	changeSeq[1] = changeSeq[2]
	changeSeq[2] = changeSeq[3]
	changeSeq[3] = bananas - prevBananas
	return changeSeq
}

func PartTwo(input []string) int {
	result := 0
	allChanges := make([]WormSet, 0, len(input))

	for _, line := range input {
		changeSeq := ChangeSeq{}
		secret := utils.Atoi(line)
		if *debug {
			fmt.Println(secret)
		}
		changes := WormSet{}
		prevBananas := numToBananas(secret)
		for i := 0; i < 2000; i++ {
			secret = next(secret)
			bananas := numToBananas(secret)
			if i > 3 {
				changeSeq = rotateChanges(changeSeq, bananas, prevBananas)
				changes.tryAdd(changeSeq, bananas)
			} else if i == 3 {
				changeSeq[i] = bananas - prevBananas
				changes.tryAdd(changeSeq, bananas)
			} else {
				changeSeq[i] = bananas - prevBananas
			}
			prevBananas = bananas
		}
		allChanges = append(allChanges, changes)
	}

	combinedChanges := WormSet{}

	for _, set := range allChanges {
		for k, v := range set {
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
