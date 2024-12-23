package utils

import "fmt"

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

func (p Coord) Equals(q Coord) bool {
	return p.Row == q.Row && p.Col == q.Col
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

var GridDirs = map[string]Coord{
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
	point := p.Add(GridDirs[str])
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
