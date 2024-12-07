package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func PartOne(input []string) int {
	//result := 0

	rules := strings.Split(input[0], "\n")
	updatesList := strings.Split(input[1], "\n")
	updates := make([][]string, len(updatesList))

	for i, update := range updatesList {
		updates[i] = strings.Split(update, ",")
	}

	var result int64
	var wg sync.WaitGroup
	for _, update := range updates {
		wg.Add(1)
		go func() {
			defer wg.Done()
			correct := true
			for i, page := range update {
				for j := 0; j < i; j++ {
					if !slices.Contains(rules, update[j]+"|"+page) {
						correct = false
						break
					}
				}
				for j := i + 1; j < len(update); j++ {
					if !slices.Contains(rules, page+"|"+update[j]) {
						correct = false
						break
					}
				}
				if !correct {
					break
				}
			}
			if correct {
				tmp, _ := strconv.Atoi(update[len(update)/2])
				atomic.AddInt64(&result, int64(tmp))

			}
		}()
	}

	wg.Wait()

	return int(result)
}

func PartOnePartTwoStyle(input []string) int {
	//result := 0

	type Rules struct {
		before []string
		after  []string
	}

	rulesList := strings.Split(input[0], "\n")
	rules := make(map[string]*Rules)
	updatesList := strings.Split(input[1], "\n")
	updates := make([][]string, len(updatesList))

	for _, rule := range rulesList {
		parts := strings.Split(rule, "|")

		if _, ok := rules[parts[0]]; !ok {
			rules[parts[0]] = &Rules{before: []string{}, after: []string{}}
		}
		rules[parts[0]].before = append(rules[parts[0]].before, parts[1])

		if _, ok := rules[parts[1]]; !ok {
			rules[parts[1]] = &Rules{before: []string{}, after: []string{}}
		}
		rules[parts[1]].after = append(rules[parts[1]].after, parts[0])

	}
	for i, update := range updatesList {
		updates[i] = strings.Split(update, ",")
	}

	var result int64
	var wg sync.WaitGroup

	for _, update := range updates {
		wg.Add(1)
		go func() {
			defer wg.Done()
			correct := true
			for i, page := range update {
				for j := 0; j < i; j++ {
					if !slices.Contains(rules[page].after, update[j]) {
						correct = false
						break
					}
				}
				for j := i + 1; j < len(update); j++ {
					if !slices.Contains(rules[page].before, update[j]) {
						correct = false
						break
					}
				}
			}
			if correct {
				tmp, _ := strconv.Atoi(update[len(update)/2])
				atomic.AddInt64(&result, int64(tmp))

			}
		}()
	}

	wg.Wait()

	return int(result)
}

func PartTwo(input []string) int {
	//result := 0

	type Rules struct {
		before []string
		after  []string
	}

	rulesList := strings.Split(input[0], "\n")
	rules := make(map[string]*Rules)
	updatesList := strings.Split(input[1], "\n")
	updates := make([][]string, len(updatesList))

	for _, rule := range rulesList {
		parts := strings.Split(rule, "|")

		if _, ok := rules[parts[0]]; !ok {
			rules[parts[0]] = &Rules{before: []string{}, after: []string{}}
		}
		rules[parts[0]].before = append(rules[parts[0]].before, parts[1])

		if _, ok := rules[parts[1]]; !ok {
			rules[parts[1]] = &Rules{before: []string{}, after: []string{}}
		}
		rules[parts[1]].after = append(rules[parts[1]].after, parts[0])

	}
	for i, update := range updatesList {
		updates[i] = strings.Split(update, ",")
	}

	reorder := func(blah []string) int {
		slices.SortFunc(blah, func(a, b string) int {
			var aCount, bCount int
			for _, bleh := range blah {
				if slices.Contains(rules[a].after, bleh) {
					aCount++
				}
				if slices.Contains(rules[b].after, bleh) {
					bCount++
				}
			}
			return aCount - bCount
		})
		tmp, _ := strconv.Atoi(blah[len(blah)/2])
		return tmp
	}

	var result int64
	var wg sync.WaitGroup

	for _, update := range updates {
		wg.Add(1)
		go func() {
			defer wg.Done()
			correct := true
			for i, page := range update {
				for j := 0; j < i; j++ {
					if !slices.Contains(rules[page].after, update[j]) {
						correct = false
						atomic.AddInt64(&result, int64(reorder(update)))
						break
					}
				}
				for j := i + 1; j < len(update); j++ {
					if !slices.Contains(rules[page].before, update[j]) {
						correct = false
						atomic.AddInt64(&result, int64(reorder(update)))
						break
					}
				}
				if !correct {
					break
				}
			}
		}()
	}

	wg.Wait()

	return int(result)
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
