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

type HashGrid[T any] map[Point]T

type GridCell[T any] struct {
	Exists bool
	Val    T
	Point  Point
}

var gridDirs = map[string]Point{
	"n":  {X: 0, Y: -1},
	"ne": {X: 1, Y: -1},
	"e":  {X: 1, Y: 0},
	"se": {X: 1, Y: 1},
	"s":  {X: 0, Y: 1},
	"sw": {X: -1, Y: 1},
	"w":  {X: -1, Y: 0},
	"nw": {X: -1, Y: -1},
}

func (i HashGrid[T]) Dir(p Point, str string) GridCell[T] {
	point := p.Add(gridDirs[str])
	val, exists := i[point]

	return GridCell[T]{Exists: exists, Val: val, Point: point}
}

func Abs[T int | int8 | int16 | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi[T string | rune](x T) int {
	xStr := string(x)
	num, _ := strconv.Atoi(xStr)
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
