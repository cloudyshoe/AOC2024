package main

import (
	"aoc/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

type Robot struct {
	loc     utils.Coord
	vel     utils.Coord
	counted bool
}

func PartOne(input []string, maxRows int, maxCols int) int {
	result := 1
	robots := make([]Robot, len(input))

	for i, line := range input {
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.loc.Col, &robot.loc.Row, &robot.vel.Col, &robot.vel.Row)
		seconds := 0
		for seconds < 100 {
			robot.loc = robot.loc.Add(robot.vel)
			if robot.loc.Col < 0 {
				robot.loc.Col = maxCols + robot.loc.Col
			}
			if robot.loc.Col >= maxCols {
				robot.loc.Col = robot.loc.Col - maxCols
			}
			if robot.loc.Row < 0 {
				robot.loc.Row = maxRows + robot.loc.Row
			}
			if robot.loc.Row >= maxRows {
				robot.loc.Row = robot.loc.Row - maxRows
			}
			seconds++
		}
		robots[i] = robot
		if *debug {
			fmt.Println(robot)
		}
	}

	quadrants := []utils.Bounds{
		{Min: utils.Coord{Row: 0, Col: 0}, Max: utils.Coord{Row: maxRows / 2, Col: maxCols / 2}},
		{Min: utils.Coord{Row: 0, Col: maxCols/2 + 1}, Max: utils.Coord{Row: maxRows / 2, Col: maxCols}},
		{Min: utils.Coord{Row: maxRows/2 + 1, Col: 0}, Max: utils.Coord{Row: maxRows, Col: maxCols / 2}},
		{Min: utils.Coord{Row: maxRows/2 + 1, Col: maxCols/2 + 1}, Max: utils.Coord{Row: maxRows, Col: maxCols}},
	}

	for _, quandrant := range quadrants {
		if *debug {
			fmt.Println(quandrant)
		}
		count := 0
		for _, robot := range robots {
			if !robot.counted {
				if robot.loc.In(quandrant) {
					count++
					robot.counted = true
				}

			}
		}
		result *= count
	}

	if *debug {
		for row := 0; row < maxRows; row++ {
			for col := 0; col < maxCols; col++ {
				found := false
				for _, robot := range robots {
					if robot.loc.Row == row && robot.loc.Col == col {
						fmt.Print("#")
						found = true
						break
					}
				}
				if !found {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
	}

	return result
}

func PartTwo(input []string, maxRows int, maxCols int) int {
	result := 0
	robots := make([]Robot, len(input))
	seconds := 1

	for i, line := range input {
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.loc.Col, &robot.loc.Row, &robot.vel.Col, &robot.vel.Row)
		robots[i] = robot
		if *debug {
			fmt.Println(robot)
		}
	}

	key := bufio.NewReader(os.Stdin)

	for {
		for i, robot := range robots {
			robot.loc = robot.loc.Add(robot.vel)
			if robot.loc.Col < 0 {
				robot.loc.Col = maxCols + robot.loc.Col
			}
			if robot.loc.Col >= maxCols {
				robot.loc.Col = robot.loc.Col - maxCols
			}
			if robot.loc.Row < 0 {
				robot.loc.Row = maxRows + robot.loc.Row
			}
			if robot.loc.Row >= maxRows {
				robot.loc.Row = robot.loc.Row - maxRows
			}
			robots[i] = robot
		}

		stop := false
		out := ""
		for row := 0; row < maxRows; row++ {
			consecutive := 0
			for col := 0; col < maxCols; col++ {
				found := false
				for _, robot := range robots {
					if robot.loc.Row == row && robot.loc.Col == col {
						out += "*"
						consecutive += 1
						if consecutive > 9 {
							stop = true
						}
						found = true
						break
					}
				}
				if !found {
					consecutive = 0
					out += " "
				}
			}
			out += "\n"
		}
		if stop {
			fmt.Print(out)
			fmt.Println(seconds)
			blah, _, _ := key.ReadLine()

			if string(blah) == "EXIT" {
				break
			}
		}
		seconds++
	}

	return result
}

func main() {

	flag.Parse()

	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	//exampleFile, _ := os.ReadFile("example.txt")
	//example := strings.Split(string(exampleFile), "\n")

	partOneResult := PartOne(input, 103, 101)
	fmt.Println("Part One Result:", partOneResult)

	//partTwoTest := PartTwo(example, 7, 11)
	//fmt.Println("Part Two Result:", partTwoTest)

	partTwoResult := PartTwo(input, 103, 101)
	fmt.Println("Part Two Result:", partTwoResult)
}
