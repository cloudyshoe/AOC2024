package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

	var wg sync.WaitGroup
	var acc atomic.Int64
	for _, line := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			secret := utils.Atoi(line)
			for range 2000 {
				secret = next(secret)
			}
			acc.Add(int64(secret))
		}()
	}

	wg.Wait()

	result = int(acc.Load())

	return result
}

func numToBananas(secret int) int {
	bananaStr := strconv.Itoa(secret)
	bananaLen := len(bananaStr)
	return utils.Atoi(bananaStr[bananaLen-1 : bananaLen])
}

func rotateChanges(changeSeq [4]int, bananas, prevBananas int) [4]int {
	changeSeq[0] = changeSeq[1]
	changeSeq[1] = changeSeq[2]
	changeSeq[2] = changeSeq[3]
	changeSeq[3] = bananas - prevBananas
	return changeSeq
}

func PartTwo(input []string) int {
	result := 0
	allChanges := make([]map[[4]int]int, len(input))
	combinedChanges := make(map[[4]int]int)

	var wg sync.WaitGroup

	for i, line := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()

			changeSeq := [4]int{}
			secret := utils.Atoi(line)
			changes := make(map[[4]int]int)
			prevBananas := numToBananas(secret)
			for i := 0; i < 2000; i++ {
				secret = next(secret)
				bananas := numToBananas(secret)
				if i < 3 {
					changeSeq[i] = bananas - prevBananas
				} else if i == 3 {
					changeSeq[i] = bananas - prevBananas
					changes[changeSeq] = bananas
				} else {
					changeSeq = rotateChanges(changeSeq, bananas, prevBananas)
					if _, ok := changes[changeSeq]; !ok {
						changes[changeSeq] = bananas
					}
				}
				prevBananas = bananas
			}
			allChanges[i] = changes
		}()
	}

	wg.Wait()

	for _, changes := range allChanges {
		for k, v := range changes {
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
