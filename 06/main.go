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

var visited = make(map[Coords]bool)

func PartOne(input []string) int {
	result := 0

	pos := Coords{}
	posFound := false
	output := make(map[string]bool)

	for y, row := range input {
		for x, char := range row {
			if char == '^' {
				coords := strconv.Itoa(x) + "," + strconv.Itoa(y)
				output[coords] = true
				pos.x, pos.y = x, y
				posFound = true
				break
			}
		}
		if posFound {
			break
		}
	}

	dir := Coords{x: 0, y: -1}

	for pos.x != 0 && pos.y != 0 && pos.x != len(input[0])-1 && pos.y != len(input)-1 {
		if input[pos.y+dir.y][pos.x+dir.x] == '#' {
			if dir.y == -1 {
				dir.x = 1
				dir.y = 0
			} else if dir.x == 1 {
				dir.x = 0
				dir.y = 1
			} else if dir.y == 1 {
				dir.x = -1
				dir.y = 0
			} else {
				dir.x = 0
				dir.y = -1
			}
		}

		pos.x += dir.x
		pos.y += dir.y
		coords := strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
		output[coords] = true
		visited[pos] = true
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
		output := make(map[Coords]string)

		dir := Coords{x: 0, y: -1}
		history := make(map[Heading]bool)
		heading := Heading{pos: pos, dir: dir}
		history[heading] = true

		for pos.x != 0 && pos.y != 0 && pos.x != len(input[0])-1 && pos.y != len(input)-1 {
			for input[pos.y+dir.y][pos.x+dir.x] == '#' {
				if dir.y == -1 {
					dir.x = 1
					dir.y = 0
				} else if dir.x == 1 {
					dir.x = 0
					dir.y = 1
				} else if dir.y == 1 {
					dir.x = -1
					dir.y = 0
				} else {
					dir.x = 0
					dir.y = -1
				}
			}

			pos.x += dir.x
			pos.y += dir.y
			heading.pos = pos
			heading.dir = dir

			if _, ok := history[heading]; ok {
				//fmt.Println(result, obs.y, obs.x)
				return 1
			}

			history[heading] = true
			output[pos] = "X"
		}

		/*
			for y, row := range input {
				line := ""
				for x := range row {
					if _, ok := output[Coords{x: x, y: y}]; ok {
						if input[y][x] == '#' {
							line += "?"
						} else {
							line += "X"
						}
					} else {
						line += string(input[y][x])
					}
				}
				line += " " + strconv.Itoa(y)
				fmt.Println(line)
			}
		*/

		return 0
	}

	/*
		for y, row := range input {
			for x := range row {
				obs := Coords{
					x: x,
					y: y,
				}
				result += checkObstacleLocation(pos, obs, input)
			}
		}
	*/

	for obs := range visited {
		result += checkObstacleLocation(pos, obs, input)
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
