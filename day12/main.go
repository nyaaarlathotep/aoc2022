package main

import (
	"aoc2022/BFS"
	"aoc2022/DataStruct"
	"aoc2022/c"
	"aoc2022/util"
	"golang.org/x/exp/maps"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	inputM := util.GetInput("12")
	heightMap := util.Get2dString(inputM, "\n", "")
	g, s, e := initialHeightRuneMap(heightMap)

	partOne(g, s, e)
	partTwo(g, e)

	elapsed := time.Now().Sub(startTime)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partTwo(g *DataStruct.Grid[rune], e DataStruct.Point) {
	starts := g.StateMapWhere(func(r rune) bool { return r == 'a' })
	startPoints := maps.Keys(starts)
	pathMap := c.Map(startPoints, func(p DataStruct.Point) int {
		return len(BFS.BFS(HeightMap{g}, p, e))
	})
	pathMap = c.Select(pathMap, func(path int) bool { return path > 1 })
	log.Printf("%+v", c.Min(pathMap)-1)
}

func partOne(g *DataStruct.Grid[rune], s DataStruct.Point, e DataStruct.Point) {
	log.Printf("%v", len(BFS.BFS(HeightMap{g}, s, e))-1)
}

func initialHeightRuneMap(heightMap [][]string) (*DataStruct.Grid[rune], DataStruct.Point, DataStruct.Point) {

	input := DataStruct.NewGrid[rune](int64(len(heightMap[0])), int64(len(heightMap)), DataStruct.Directions4)
	var start, end DataStruct.Point
	for y, l := range heightMap {
		for x, s := range l {
			switch s {
			case "S":
				start = DataStruct.Point{X: int64(x), Y: int64(y)}
				s = "a"
			case "E":
				end = DataStruct.Point{X: int64(x), Y: int64(y)}
				s = "z"
			}
			input.SetState(int64(x), int64(y), rune(s[0]))
		}
	}
	return input, start, end
}

type HeightMap struct {
	*DataStruct.Grid[rune]
}

func (h HeightMap) Neighbours(p DataStruct.Point) []DataStruct.Point {
	height := h.GetState(p.Y, p.X)
	return c.Select(h.Grid.Neighbours(p), func(p DataStruct.Point) bool {
		return h.GetState(p.Y, p.X) <= height+1
	})
}
