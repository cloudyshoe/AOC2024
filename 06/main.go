package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Heading struct {
	pos Coords
	dir Coords
}

var visited = make(map[Coords]struct{})

func PartOne(input []string) int {
	result := 0

	pos := Coords{}
	posFound := false
	output := make(map[string]struct{})

	for y, row := range input {
		for x, char := range row {
			if char == '^' {
				coords := strconv.Itoa(x) + "," + strconv.Itoa(y)
				output[coords] = struct{}{}
				pos.x, pos.y = x, y
				posFound = true
				break
			}
		}
		if posFound {
			break
		}
	}

	dirs := []Coords{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	dirIndex := 0
	dir := dirs[dirIndex]

	for pos.x != 0 && pos.y != 0 && pos.x != len(input[0])-1 && pos.y != len(input)-1 {
		for input[pos.y+dir.y][pos.x+dir.x] == '#' {
			dirIndex = (dirIndex + 1) % len(dirs)
			dir = dirs[dirIndex]
		}

		pos.x += dir.x
		pos.y += dir.y
		coords := strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
		output[coords] = struct{}{}
		visited[pos] = struct{}{}
	}

	for y, row := range input {
		blah := ""
		for x := range row {
			coords := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if _, ok := output[coords]; ok {
				blah += "X"
			} else {
				blah += "."
			}
		}
	}

	result = len(output)

	return result
}

func PartTwo(input []string) int {
	result := 0

	pos := Coords{}
	posFound := false

	for y, row := range input {
		for x, char := range row {
			if char == '^' {
				pos.x, pos.y = x, y
				posFound = true
				break
			}
		}
		if posFound {
			break
		}
	}
	checkObstacleLocation := func(pos Coords, obs Coords, origInput []string) int {

		input := make([]string, len(origInput))
		copy(input, origInput)

		strSlice := []byte(input[obs.y])
		strSlice[obs.x] = '#'
		input[obs.y] = string(strSlice)

		dirs := []Coords{
			{x: 0, y: -1},
			{x: 1, y: 0},
			{x: 0, y: 1},
			{x: -1, y: 0},
		}

		dirIndex := 0
		dir := dirs[dirIndex]
		history := make(map[Heading]struct{})
		heading := Heading{pos: pos, dir: dir}
		history[heading] = struct{}{}

		for pos.x != 0 && pos.y != 0 && pos.x != len(input[0])-1 && pos.y != len(input)-1 {
			for input[pos.y+dir.y][pos.x+dir.x] == '#' {
				dirIndex = (dirIndex + 1) % len(dirs)
				dir = dirs[dirIndex]
			}

			pos.x += dir.x
			pos.y += dir.y
			heading.pos = pos
			heading.dir = dir

			if _, ok := history[heading]; ok {
				return 1
			}

			history[heading] = struct{}{}
		}

		return 0
	}

	for obs := range visited {
		if !(pos.x == obs.x && pos.y == obs.y) {
			result += checkObstacleLocation(pos, obs, input)
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
