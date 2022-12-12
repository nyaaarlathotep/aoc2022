package main

import (
	"aoc2022/BFS"
	"aoc2022/DataStruct"
	"aoc2022/c"
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	inputM := util.GetInput("12")
	heightMap := util.Get2dString(inputM, "\n", "")
	//g, _, _ := initialHeightRuneMap(heightMap)
	g, s, e := initialHeightRuneMap(heightMap)
	log.Printf("%+v", g)
	log.Printf("%+v", HeightMap{g}.Neighbours(DataStruct.Point{
		X: 0,
		Y: 0,
	}))

	log.Printf("%v", len(BFS.MineBFS(HeightMap{g}, s, e))-1)
	//var start, end aoc.Point
	//data := aoc.SliceFromFile("day12.txt", func(l string) []rune {
	//	return []rune(l)
	//})
	//input := grid.New[rune](int64(len(data[0])), int64(len(data)), grid.Directions4)
	//for y, l := range data {
	//	for x, s := range l {
	//		switch s {
	//		case 'S':
	//			start = aoc.Point{int64(x), int64(y)}
	//			s = 'a'
	//		case 'E':
	//			end = aoc.Point{int64(x), int64(y)}
	//			s = 'z'
	//		}
	//		input.SetState(int64(x), int64(y), s)
	//	}
	//}
	//log.Printf("%+v", input)
	//log.Printf("%+v", HeightMap{input}.Neighbours(aoc.Point{
	//	X: 0,
	//	Y: 0,
	//}))
	//
	//fmt.Println(do1(HeightMap{input}, start, end))
	elapsed := time.Now().Sub(startTime)
	log.Println("该函数执行完成耗时：", elapsed)
}

//func main() {
//	var start, end aoc.Point
//	data := aoc.SliceFromFile("day12.txt", func(l string) []rune {
//		return []rune(l)
//	})
//	input := DataStruct.NewGrid[rune](int64(len(data[0])), int64(len(data)), grid.Directions4)
//	for y, l := range data {
//		for x, s := range l {
//			switch s {
//			case 'S':
//				start = aoc.Point{int64(x), int64(y)}
//				s = 'a'
//			case 'E':
//				end = aoc.Point{int64(x), int64(y)}
//				s = 'z'
//			}
//			input.SetState(int64(x), int64(y), s)
//		}
//	}
//
//	fmt.Println(do1(HeightMap{input}, start, end))
//	fmt.Println(do2(HeightMap{input}, end))
//}

//func do1(in HeightMap, start, end aoc.Point) int {
//	return len(search.BFS(in, start, end)) - 1
//}

//func do2(in HeightMap, end aoc.Point) int {
//	return c.Min(c.Select(c.Map(maps.Keys(in.StateMapWhere(func(i rune) bool { return i == 'a' })), func(p aoc.Point) int {
//		return len(BFS.BFS(in, p, end))
//	}), func(i int) bool {
//		return i > 1
//	})) - 1
//}

func initialStepMap(heightMap [][]string) [][]int {
	stepsMap := make([][]int, len(heightMap))
	for i := range heightMap {
		stepsMap[i] = make([]int, len(heightMap[0]))
		for j := range heightMap[0] {
			stepsMap[i][j] = -1
		}
	}
	return stepsMap
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
	log.Printf("%+v", h.Grid.Neighbours(p))
	return c.Select(h.Grid.Neighbours(p), func(p DataStruct.Point) bool {
		return h.GetState(p.Y, p.X) <= height+1
	})
}

//
//type HeightMap struct {
//	*grid.Grid[rune]
//}
//
//func (h HeightMap) Neighbours(p aoc.Point) []aoc.Point {
//	height := h.GetState(p.Y, p.X)
//	return c.Select(h.Grid.Neighbours(p), func(x aoc.Point) bool {
//		return h.GetState(x.Y, x.X) <= height+1
//	})
//}
