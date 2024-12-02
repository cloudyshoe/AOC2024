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
    fmt.Println("Part One Result:", partOneResult)

    partTwoResult := PartTwo(input)
    fmt.Println("Part Two Result:", partTwoResult)
}
