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

func PartOne(input []string) int {
	result := 0

	for _, line := range input {
		secret := utils.Atoi(line)
		if *debug {
			fmt.Println(secret)
		}
		i := 0
		for i < 2000 {
			secret ^= secret * 64
			secret %= 16777216
			secret ^= secret / 32
			secret %= 16777216
			secret ^= secret * 2048
			secret = secret % 16777216
			i++
		}
		if *debug {
			fmt.Println(secret)
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
		bananaStr := strconv.Itoa(secret)
		bananaLen := len(bananaStr)
		prevBananas := utils.Atoi(bananaStr[bananaLen-1 : bananaLen])
		i := 0
		for i < 2000 {
			bananas := 0
			secret ^= secret * 64
			secret %= 16777216
			secret ^= secret / 32
			secret %= 16777216
			secret ^= secret * 2048
			secret = secret % 16777216
			str := strconv.Itoa(secret)
			lenStr := len(str)
			bananas = utils.Atoi(str[lenStr-1 : lenStr])
			if i > 3 {
				changeSeq[0] = changeSeq[1]
				changeSeq[1] = changeSeq[2]
				changeSeq[2] = changeSeq[3]
				changeSeq[3] = bananas - prevBananas
				changes.tryAdd(changeSeq, bananas)
			} else if i == 3 {
				changeSeq[i] = bananas - prevBananas
				changes.tryAdd(changeSeq, bananas)
			} else {
				changeSeq[i] = bananas - prevBananas
			}
			prevBananas = bananas
			i++
		}
		allChanges = append(allChanges, changes)
		if *debug {
			fmt.Println(secret)
		}
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
