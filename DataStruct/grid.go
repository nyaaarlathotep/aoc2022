package DataStruct

import (
	"strings"

	"golang.org/x/exp/slices"
)

var Directions4 = []Point{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

var Directions8 = []Point{
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
}

type Grid[T any] struct {
	xLen      int64
	yLen      int64
	state     [][]T
	movements []Point
}

func NewGrid[T any](xLen, yLen int64, movements []Point) *Grid[T] {
	state := make([][]T, yLen)
	for y := int64(0); y < yLen; y++ {
		state[y] = make([]T, xLen)
	}
	return &Grid[T]{
		xLen:      xLen,
		yLen:      yLen,
		state:     state,
		movements: movements,
	}
}

func (g *Grid[T]) isValid(x, y int64) bool {
	switch {
	case x < 0, x >= g.xLen, y < 0, y >= g.yLen:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) isValidPoint(p Point) bool {
	switch {
	case p.X < 0, p.X >= g.xLen, p.Y < 0, p.Y >= g.yLen:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) Neighbours(p Point) []Point {
	var ret []Point

	for _, m := range g.movements {
		np := p.Add(m)
		if g.isValidPoint(np) {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *Grid[T]) XLen() int64 {
	return g.xLen
}

func (g *Grid[T]) YLen() int64 {
	return g.yLen
}

func (g *Grid[T]) GetSliceToEdge(x, y int64, movement Point) []T {
	var ret []T
	p := Point{x, y}
	for ; g.isValidPoint(p); p = p.Add(movement) {
		ret = append(ret, g.GetState(p.Y, p.X))
	}

	return ret
}

func (g *Grid[T]) SetState(x, y int64, state T) {
	if g.isValid(x, y) {
		g.state[y][x] = state
	}
}

func (g *Grid[T]) GetState(y, x int64) T {
	return g.state[y][x]
}

func (g *Grid[T]) GetPointState(p Point) T {
	return g.state[p.Y][p.X]
}

func (g *Grid[T]) StateString(f func(t T) string) string {
	var ret strings.Builder

	for i, y := range g.state {
		if i > 40 {
			break
		}
		for _, x := range y {
			ret.WriteString(f(x))
		}
		ret.WriteRune('\n')
	}

	return ret.String()
}

func (g *Grid[T]) StateMap() map[Point]T {
	ret := make(map[Point]T)

	for y, l := range g.state {
		for x, s := range l {
			ret[Point{int64(x), int64(y)}] = s
		}
	}

	return ret
}

func (g *Grid[T]) StateMapWhere(f func(T) bool) map[Point]T {
	ret := make(map[Point]T)

	for y, l := range g.state {
		for x, s := range l {
			if f(s) {
				ret[Point{int64(x), int64(y)}] = s
			}
		}
	}

	return ret
}

func (g *Grid[T]) Clone() *Grid[T] {
	ng := Grid[T]{
		xLen:      g.xLen,
		yLen:      g.yLen,
		movements: slices.Clone(g.movements),
		state:     make([][]T, g.yLen),
	}

	for yi := range g.state {
		ng.state[yi] = slices.Clone(g.state[yi])
	}

	return &ng
}
