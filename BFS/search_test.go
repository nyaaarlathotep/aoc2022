package BFS

import (
	"aoc2022/DataStruct"
	"fmt"
	"testing"
)

func TestBFSTestBFS(t *testing.T) {
	g := DataStruct.NewGrid[int](10, 10, DataStruct.Directions4)

	route := MineBFS(g, DataStruct.Point{0, 0}, DataStruct.Point{5, 5})

	fmt.Println(route)
}
