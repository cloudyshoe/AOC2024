package utils

import (
	"fmt"
	"strconv"
)

type Coord struct {
	Col int
	Row int
}

type Bounds struct {
	Min Coord
	Max Coord
}

func (p Coord) Add(q Coord) Coord {
	return Coord{Col: p.Col + q.Col, Row: p.Row + q.Row}
}

func (p Coord) AddMod(q, m Coord) Coord {
	return Coord{Col: ((p.Col+q.Col)%m.Col + m.Col) % m.Col, Row: ((p.Row+q.Row)%m.Row + m.Row) % m.Row}
}

func (p Coord) Sub(q Coord) Coord {
	return Coord{Col: p.Col - q.Col, Row: p.Row - q.Row}
}

func (p Coord) In(q Bounds) bool {
	return q.Min.Col <= p.Col && p.Col < q.Max.Col &&
		q.Min.Row <= p.Row && p.Row < q.Max.Row
}

func (p Coord) String() string {
	return fmt.Sprintf("{Row: %d, Col: %d}", p.Row, p.Col)
}

type HashGrid[T any] map[Coord]T

type GridCell[T any] struct {
	Exists bool
	Val    T
	Point  Coord
}

var gridDirs = map[string]Coord{
	"n":    {Row: -1, Col: 0},
	"ne":   {Row: -1, Col: 1},
	"e":    {Row: 0, Col: 1},
	"se":   {Row: 1, Col: 1},
	"s":    {Row: 1, Col: 0},
	"sw":   {Row: 1, Col: -1},
	"w":    {Row: 0, Col: -1},
	"nw":   {Row: -1, Col: -1},
	"this": {Row: 0, Col: 0},
}

func (h HashGrid[T]) Dir(p Coord, str string) GridCell[T] {
	point := p.Add(gridDirs[str])
	val, exists := h[point]

	return GridCell[T]{Exists: exists, Val: val, Point: point}
}

func (h HashGrid[T]) Clone() HashGrid[T] {
	hashGrid := make(HashGrid[T], len(h))
	for k, v := range h {
		hashGrid[k] = v
	}
	return hashGrid
}

type Grid[T any] [][]T

func (g Grid[T]) Dir(c Coord, str string) GridCell[T] {
	cell := GridCell[T]{}
	cell.Point = c.Add(gridDirs[str])
	if len(g) > 0 && len(g[0]) > 0 &&
		cell.Point.Col >= 0 && cell.Point.Col < len(g[0]) &&
		cell.Point.Row >= 0 && cell.Point.Row < len(g) {
		cell.Exists = true
		cell.Val = g[c.Row][c.Col]
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

func Min[T int | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Pow2[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](exp T) T {
	return (1 << exp)
}
