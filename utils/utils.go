package utils

import (
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

type Grid map[Point]rune

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
