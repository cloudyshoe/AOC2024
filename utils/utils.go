package utils

import (
	"fmt"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Bounds struct {
	Min Point
	Max Point
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Point) In(q Bounds) bool {
	return q.Min.X <= p.X && p.X < q.Max.X &&
		q.Min.Y <= p.Y && p.Y < q.Max.Y
}

func (p Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

type Grid map[Point]rune

type IntGrid map[Point]int

type IntGridCell struct {
	Exists bool
	Val    int
	Point  Point
}

var gridDirs = map[string]Point{
	"n": {X: 0, Y: -1},
	"e": {X: 1, Y: 0},
	"s": {X: 0, Y: 1},
	"w": {X: -1, Y: 0},
}

func (i IntGrid) Dir(p Point, str string) IntGridCell {
	exists := false
	val := 0
	point := p.Add(gridDirs[str])
	val, exists = i[point]

	return IntGridCell{Exists: exists, Val: val, Point: point}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi(x string) int {
	num, _ := strconv.Atoi(x)
	return num
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a) * (Abs(b) / GCD(a, b))
}
