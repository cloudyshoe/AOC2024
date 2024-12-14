package main

import (
    "os"
	"flag"
    "fmt"
    "strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func PartOne(input []string) int  {
    result := 0

    return result
}

func PartTwo(input []string) int {
    result := 0

    return result
}

func main () {

	flag.Parse()

    inputFile, _ := os.ReadFile("input.txt")
    input := strings.Split(string(inputFile), "\n")

    partOneResult := PartOne(input)
    fmt.Println("Part One Result:", partOneResult)

    partTwoResult := PartTwo(input)
    fmt.Println("Part Two Result:", partTwoResult)
}
