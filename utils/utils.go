package utils

import (
	"fmt"
	"strconv"
)

type Coords struct {
	X int
	Y int
}

type Bounds struct {
	Min Coords
	Max Coords
}

func (p Coords) Add(q Coords) Coords {
	return Coords{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Coords) Sub(q Coords) Coords {
	return Coords{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Coords) In(q Bounds) bool {
	return q.Min.X <= p.X && p.X < q.Max.X &&
		q.Min.Y <= p.Y && p.Y < q.Max.Y
}

func (p Coords) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

type HashGrid[T any] map[Coords]T

type GridCell[T any] struct {
	Exists bool
	Val    T
	Point  Coords
}

var gridDirs = map[string]Coords{
	"n":  {X: 0, Y: -1},
	"ne": {X: 1, Y: -1},
	"e":  {X: 1, Y: 0},
	"se": {X: 1, Y: 1},
	"s":  {X: 0, Y: 1},
	"sw": {X: -1, Y: 1},
	"w":  {X: -1, Y: 0},
	"nw": {X: -1, Y: -1},
}

func (i HashGrid[T]) Dir(p Coords, str string) GridCell[T] {
	point := p.Add(gridDirs[str])
	val, exists := i[point]

	return GridCell[T]{Exists: exists, Val: val, Point: point}
}

type Grid[T any] [][]T

func (g Grid[T]) Dir(p Coords, str string) GridCell[T] {
	cell := GridCell[T]{}
	cell.Point = p.Add(gridDirs[str])
	if len(g) > 0 && len(g[0]) > 0 &&
		p.X >= 0 && p.X <= len(g) && p.Y >= 0 && p.Y <= len(g[0]) {
		cell.Exists = true
		cell.Val = g[p.Y][p.X]
	} else {
		cell.Exists = false
	}

	return cell
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
