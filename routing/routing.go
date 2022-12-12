package routing

import (
	"aoc2022/DataStruct"
)

type Graph interface {
	Neighbours(p DataStruct.Point) []DataStruct.Point
}
