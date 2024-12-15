package main

import (
	. "aoc/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool = flag.Bool("debug", false, "Print debug statements")

func tryRobotMove(robot Coord, dir Coord, boxes, walls HashGrid[struct{}]) Coord {
	pos := robot.Add(dir)
	_, isWall := walls[pos]
	_, isBox := boxes[pos]
	if isWall {
		return robot
	} else if isBox {
		moved := tryBoxMove(pos, dir, boxes, walls)
		if *debug {
			//fmt.Println(moved)
		}
		if !moved {
			return robot
		}
		return pos
	}

	return pos
}

func tryBoxMove(box Coord, dir Coord, boxes, walls HashGrid[struct{}]) bool {
	pos := box.Add(dir)

	if _, ok := walls[pos]; ok {
		return false
	}

	if _, ok := boxes[pos]; !ok {
		boxes[pos] = struct{}{}
		delete(boxes, box)
		return true
	}

	moved := tryBoxMove(pos, dir, boxes, walls)

	if moved {
		boxes[pos] = struct{}{}
		delete(boxes, box)
		return true
	}

	return false
}

func tryRobotMoveTwo(robot Coord, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) Coord {
	pos := robot.Add(dir)
	boxIdx, isBox := boxes[pos]

	_, isWall := walls[pos]
	if isWall {
		return robot
	}

	if !isBox {
		return pos
	}

	var moved bool
	if dir.Col == -1 {
		moved = canBigBoxMoveLeft(boxIdx, dir, boxInfo, boxes, walls)
		if moved {
			moveBigBoxLeft(boxIdx, dir, boxInfo, boxes, walls)
			return pos
		}
		return robot
	} else if dir.Col == 1 {
		moved = canBigBoxMoveRight(boxIdx, dir, boxInfo, boxes, walls)
		if moved {
			moveBigBoxRight(boxIdx, dir, boxInfo, boxes, walls)
			return pos
		}
		return robot
	} else {
		moved = canBigBoxMoveUpOrDown(boxIdx, dir, boxInfo, boxes, walls)
		if moved {
			moveBigBoxUpOrDown(boxIdx, dir, boxInfo, boxes, walls)
			return pos
		}
		return robot
	}

}

func canBigBoxMoveRight(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) bool {
	boxPair := boxInfo[boxIdx]
	rBox := boxPair[1]

	rPos := rBox.Add(dir)
	if _, ok := walls[rPos]; ok {
		return false
	}

	rrBox, isBox := boxes[rPos]

	if !isBox {
		return true
	}

	return canBigBoxMoveRight(rrBox, dir, boxInfo, boxes, walls)
}

func moveBigBoxRight(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) {
	boxPair := boxInfo[boxIdx]
	lBox := boxPair[0]
	rBox := boxPair[1]

	lPos := lBox.Add(dir)
	rPos := rBox.Add(dir)

	rrBox, isBox := boxes[rPos]

	if isBox {
		moveBigBoxRight(rrBox, dir, boxInfo, boxes, walls)
	}

	delete(boxes, rBox)
	delete(boxes, lBox)
	boxes[rPos] = boxIdx
	boxes[lPos] = boxIdx
	boxInfo[boxIdx] = [2]Coord{lPos, rPos}

}

func canBigBoxMoveLeft(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) bool {
	boxPair := boxInfo[boxIdx]
	lBox := boxPair[0]

	lPos := lBox.Add(dir)
	if _, ok := walls[lPos]; ok {
		return false
	}

	llBox, isBox := boxes[lPos]

	if !isBox {
		return true
	}

	return canBigBoxMoveLeft(llBox, dir, boxInfo, boxes, walls)
}

func moveBigBoxLeft(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) {
	boxPair := boxInfo[boxIdx]
	lBox := boxPair[0]
	rBox := boxPair[1]

	lPos := lBox.Add(dir)
	rPos := rBox.Add(dir)

	llBox, isBox := boxes[lPos]

	if isBox {
		moveBigBoxLeft(llBox, dir, boxInfo, boxes, walls)
	}

	delete(boxes, rBox)
	delete(boxes, lBox)
	boxes[rPos] = boxIdx
	boxes[lPos] = boxIdx
	boxInfo[boxIdx] = [2]Coord{lPos, rPos}

}

func canBigBoxMoveUpOrDown(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) bool {
	boxPair := boxInfo[boxIdx]
	lBox := boxPair[0]
	rBox := boxPair[1]

	lPos := lBox.Add(dir)
	rPos := rBox.Add(dir)
	_, isWallL := walls[lPos]
	_, isWallR := walls[rPos]

	if isWallL || isWallR {
		return false
	}

	llBox, isBoxL := boxes[lPos]
	rrBox, isBoxR := boxes[rPos]

	if !isBoxL && !isBoxR {
		return true
	}

	if llBox == rrBox {
		return canBigBoxMoveUpOrDown(llBox, dir, boxInfo, boxes, walls)
	}

	if !isBoxL {
		return canBigBoxMoveUpOrDown(rrBox, dir, boxInfo, boxes, walls)
	}

	if !isBoxR {
		return canBigBoxMoveUpOrDown(llBox, dir, boxInfo, boxes, walls)
	}

	return canBigBoxMoveUpOrDown(llBox, dir, boxInfo, boxes, walls) && canBigBoxMoveUpOrDown(rrBox, dir, boxInfo, boxes, walls)
}

func moveBigBoxUpOrDown(boxIdx int, dir Coord, boxInfo [][2]Coord, boxes HashGrid[int], walls HashGrid[struct{}]) {
	boxPair := boxInfo[boxIdx]
	lBox := boxPair[0]
	rBox := boxPair[1]

	lPos := lBox.Add(dir)
	rPos := rBox.Add(dir)

	llBox, isBoxL := boxes[lPos]
	rrBox, isBoxR := boxes[rPos]

	if isBoxL || isBoxR {
		if isBoxL && isBoxR {
			same := llBox == rrBox
			moveBigBoxUpOrDown(llBox, dir, boxInfo, boxes, walls)
			if !same {
				moveBigBoxUpOrDown(rrBox, dir, boxInfo, boxes, walls)
			}
		} else if isBoxL {
			moveBigBoxUpOrDown(llBox, dir, boxInfo, boxes, walls)
		} else if isBoxR {
			moveBigBoxUpOrDown(rrBox, dir, boxInfo, boxes, walls)
		}
	}

	delete(boxes, rBox)
	delete(boxes, lBox)
	boxes[rPos] = boxIdx
	boxes[lPos] = boxIdx
	boxInfo[boxIdx] = [2]Coord{lPos, rPos}

}

func PartOne(input []string) int {
	result := 0
	walls := make(HashGrid[struct{}])
	boxes := make(HashGrid[struct{}])
	robot := Coord{}

	layout := strings.Split(input[0], "\n")
	instuctions := strings.Split(input[1], "\n")

	for row, line := range layout {
		for col, char := range line {
			switch char {
			case '#':
				walls[Coord{Row: row, Col: col}] = struct{}{}
			case 'O':
				boxes[Coord{Row: row, Col: col}] = struct{}{}
			case '@':
				robot = Coord{Row: row, Col: col}
			}
		}
	}

	for _, line := range instuctions {
		for _, char := range line {
			switch char {
			case '^':
				robot = tryRobotMove(robot, Coord{Row: -1, Col: 0}, boxes, walls)
			case '>':
				robot = tryRobotMove(robot, Coord{Row: 0, Col: 1}, boxes, walls)
			case 'v':
				robot = tryRobotMove(robot, Coord{Row: 1, Col: 0}, boxes, walls)
			case '<':
				robot = tryRobotMove(robot, Coord{Row: 0, Col: -1}, boxes, walls)
			}
			//if *debug {
			//	fmt.Println(robot)
			//}
		}
	}

	for k := range boxes {
		result += k.Row*100 + k.Col
	}

	return result
}

func PartTwo(input []string) int {
	result := 0
	walls := make(HashGrid[struct{}])
	boxes := make(HashGrid[int])
	boxInfo := make([][2]Coord, 0, 1000)
	robot := Coord{}

	layout := strings.Split(input[0], "\n")
	instuctions := strings.Split(input[1], "\n")

	for row, line := range layout {
		for col, char := range line {
			switch char {
			case '#':
				walls[Coord{Row: row, Col: col * 2}] = struct{}{}
				walls[Coord{Row: row, Col: col*2 + 1}] = struct{}{}
			case 'O':
				l := Coord{Row: row, Col: col * 2}
				r := Coord{Row: row, Col: col*2 + 1}
				pair := [2]Coord{l, r}
				boxInfo = append(boxInfo, pair)
				boxIdx := len(boxInfo) - 1
				boxes[l] = boxIdx
				boxes[r] = boxIdx
			case '@':
				robot = Coord{Row: row, Col: col * 2}
			}
		}
	}

	for _, line := range instuctions {
		for _, char := range line {
			dir := Coord{}
			if *debug && false {
				for row := 0; row < len(layout); row++ {
					for col := 0; col < len(layout[0])*2; col++ {
						loc := Coord{Row: row, Col: col}
						if _, ok := walls[loc]; ok {
							fmt.Print("#")
						} else if _, ok := boxes[loc]; ok {
							fmt.Print("O")
						} else if loc.Row == robot.Row && loc.Col == robot.Col {
							fmt.Print("@")
						} else {
							fmt.Print(".")
						}
					}
					fmt.Print("\n")
				}
			}
			switch char {
			case '^':
				dir.Row, dir.Col = -1, 0
			case '>':
				dir.Row, dir.Col = 0, 1
			case 'v':
				dir.Row, dir.Col = 1, 0
			case '<':
				dir.Row, dir.Col = 0, -1
			}
			robot = tryRobotMoveTwo(robot, dir, boxInfo, boxes, walls)
			if *debug && false {
				fmt.Println(string(char))
			}
		}
	}

	if *debug {
		for row := 0; row < len(layout); row++ {
			for col := 0; col < len(layout[0])*2; col++ {
				loc := Coord{Row: row, Col: col}
				if _, ok := walls[loc]; ok {
					fmt.Print("#")
				} else if _, ok := boxes[loc]; ok {
					fmt.Print("O")
				} else if loc.Row == robot.Row && loc.Col == robot.Col {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
	}

	for _, box := range boxInfo {
		lBox := box[0]
		rBox := box[1]
		result += Min(lBox.Row, rBox.Row)*100 + Min(lBox.Col, rBox.Col)
	}
	return result
}

func main() {

	flag.Parse()

	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n\n")

	partOneResult := PartOne(input)
	fmt.Println("Part One Result:", partOneResult)

	partTwoResult := PartTwo(input)
	fmt.Println("Part Two Result:", partTwoResult)
}
