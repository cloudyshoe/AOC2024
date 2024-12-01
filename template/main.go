package main

import (
    "os"
    "fmt"
    "strings"
)

func PartOne(input []string) int  {
    result := 0

    return result
}

func PartTwo(input []string) int {
    result := 0

    return result
}

func main () {
    inputFile, _ := os.ReadFile("input.txt")
    input := strings.Split(string(inputFile), "\n")

    partOneResult := PartOne(input)
    fmt.Printf("Part One Result: %d\n", partOneResult)

    partTwoResult := PartTwo(input)
    fmt.Printf("Part Two Result: %d\n", partTwoResult)
}
