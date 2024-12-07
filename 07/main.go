package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TestSet struct {
	ans  int
	nums []int
}

func runOps(ans, acc int, nums []int) bool {
	if len(nums) == 0 && acc == ans {
		return true
	} else if len(nums) == 0 {
		return false
	}

	return runOps(ans, acc+nums[0], nums[1:]) || runOps(ans, acc*nums[0], nums[1:])
}

func runOps2(ans, acc int, nums []int) bool {
	if len(nums) == 0 && acc == ans {
		return true
	} else if len(nums) == 0 {
		return false
	}

	tmp := strconv.Itoa(acc) + strconv.Itoa(nums[0])
	cat, _ := strconv.Atoi(tmp)
	return runOps2(ans, acc+nums[0], nums[1:]) || runOps2(ans, acc*nums[0], nums[1:]) || runOps2(ans, cat, nums[1:])
}

func PartOne(input []string) int {
	result := 0

	testSet := make([]TestSet, len(input))
	for i, line := range input {
		parts := strings.Split(line, ":")
		ans, _ := strconv.Atoi(parts[0])
		numStrs := strings.Fields(parts[1])
		nums := make([]int, len(numStrs))
		for j, num := range numStrs {
			tmp, _ := strconv.Atoi(num)
			nums[j] = tmp
		}
		testSet[i] = TestSet{ans: ans, nums: nums}
	}

	for _, set := range testSet {
		tmp := runOps(set.ans, set.nums[0], set.nums[1:])
		if tmp {
			result += set.ans
		}
	}

	return result
}

func PartTwo(input []string) int {
	result := 0

	testSet := make([]TestSet, len(input))
	for i, line := range input {
		parts := strings.Split(line, ":")
		ans, _ := strconv.Atoi(parts[0])
		numStrs := strings.Fields(parts[1])
		nums := make([]int, len(numStrs))
		for j, num := range numStrs {
			tmp, _ := strconv.Atoi(num)
			nums[j] = tmp
		}
		testSet[i] = TestSet{ans: ans, nums: nums}
	}

	for _, set := range testSet {
		tmp := runOps2(set.ans, set.nums[0], set.nums[1:])
		if tmp {
			result += set.ans
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
