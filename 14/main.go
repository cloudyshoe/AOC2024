package main

import (
	"aoc/utils"
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
	maxCoord := utils.Coord{Row: maxRows, Col: maxCols}

	for i, line := range input {
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.loc.Col, &robot.loc.Row, &robot.vel.Col, &robot.vel.Row)
		robot.vel.Row *= 100
		robot.vel.Col *= 100
		robot.loc = robot.loc.AddMod(robot.vel, maxCoord)
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

func createVis(robots []Robot, maxRows, maxCols int) (string, bool) {
	stop := false
	out := ""
	for row := 0; row < maxRows; row++ {
		consecutive := 0
		for col := 0; col < maxCols; col++ {
			found := false
			for _, robot := range robots {
				if robot.loc.Row == row && robot.loc.Col == col {
					out += "*"
					found = true
					if consecutive++; !stop && consecutive > 9 {
						stop = true
					}
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
	return out, stop
}

func PartTwo(input []string, maxRows int, maxCols int) int {
	result := 0
	robots := make([]Robot, len(input))
	maxCoord := utils.Coord{Row: maxRows, Col: maxCols}
	seconds := 1

	for i, line := range input {
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.loc.Col, &robot.loc.Row, &robot.vel.Col, &robot.vel.Row)
		robots[i] = robot
		if *debug {
			fmt.Println(robot)
		}
	}

	for {
		for i, robot := range robots {
			robot.loc = robot.loc.AddMod(robot.vel, maxCoord)
			robots[i] = robot
		}

		out, stop := createVis(robots, maxRows, maxCols)

		if stop {
			fmt.Print(out)
			fmt.Println(seconds)
			//key := bufio.NewReader(os.Stdin)
			//blah, _, _ := key.ReadLine()

			//if string(blah) == "EXIT" {
			//	break
			result = seconds
			return result
		}
		seconds++
	}

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
