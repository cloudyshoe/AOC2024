package utils

import (
	"strconv"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Point) In(min Point, max Point) bool {
	return min.X <= p.X && p.X < max.X &&
		min.Y <= p.Y && p.Y < max.Y
}

type Grid map[Point]string

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
