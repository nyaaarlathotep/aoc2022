package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"fmt"
	"time"
)

const blank = ' '
const ground = '.'

func main() {
	start := time.Now()
	input := util.GetInput("22")
	lines := util.GetStringSlice(input, "\n")
	g := initG(lines)
	route := util.GetInputFile("22", "route")
	var posPoint *DataStruct.Point
	for x := int64(0); x < g.XLen(); x++ {
		if g.GetState(0, x) == ground {
			posPoint = &DataStruct.Point{
				X: x,
				Y: 0,
			}
			break
		}
	}
	numStr := ""
	dirIndex := int64(1)
	for _, r := range route {
		if r != 'L' && r != 'R' {
			numStr = numStr + string(r)
			continue
		}
		posPoint = move(numStr, posPoint, dirIndex, g)
		numStr = ""
		if r == 'L' {
			dirIndex = (dirIndex + 1) % 4
			continue
		}
		if r == 'R' {
			dirIndex = (dirIndex + 3) % 4
			continue
		}
	}
	posPoint = move(numStr, posPoint, dirIndex, g)
	//fmt.Printf("%+v\n", posPoint)
	fmt.Printf("%+v\n", 1000*(posPoint.Y+1)+4*(posPoint.X+1)+getDirValue(dirIndex))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func move(numStr string, posPoint *DataStruct.Point, dirIndex int64, g *DataStruct.Grid[rune]) *DataStruct.Point {
	steps := util.ParseInt(numStr)
	lastPoint := DataStruct.Point{
		X: posPoint.X,
		Y: posPoint.Y,
	}
	//fmt.Printf("%+v:%v\n", string(getDirRune(dirIndex)), steps)
	for i := 0; i < steps; i++ {
		nextPoint := nextLegalPos(lastPoint, dirIndex, g)
		nextPosState := g.GetPointState(nextPoint)
		for nextPosState == blank {
			nextPoint = nextLegalPos(nextPoint, dirIndex, g)
			nextPosState = g.GetPointState(nextPoint)
		}
		if nextPosState == '#' {
			break
		}
		g.SetState(lastPoint.X, lastPoint.Y, getDirRune(dirIndex))
		lastPoint = nextPoint
	}
	posPoint = &lastPoint
	g.SetState(lastPoint.X, lastPoint.Y, getDirRune(dirIndex))
	//printG(g)
	return posPoint
}

func nextLegalPos(lastPoint DataStruct.Point, dirIndex int64, g *DataStruct.Grid[rune]) DataStruct.Point {
	nextPoint := lastPoint.Add(DataStruct.Directions4[dirIndex])
	if nextPoint.X > g.XLen()-1 {
		nextPoint.X = 0
	}
	if nextPoint.X < 0 {
		nextPoint.X = g.XLen() - 1
	}
	if nextPoint.Y > g.YLen()-1 {
		nextPoint.Y = 0
	}
	if nextPoint.Y < 0 {
		nextPoint.Y = g.YLen() - 1
	}
	return nextPoint
}

func initG(lines []string) *DataStruct.Grid[rune] {
	maxX := 0
	for _, line := range lines {
		if len(line) > maxX {
			maxX = len(line)
		}
	}

	g := DataStruct.NewGrid[rune](int64(maxX), int64(len(lines)), DataStruct.Directions4)
	for y, line := range lines {
		if len(line) > maxX {
			maxX = len(line)
		}
		for x := 0; x < maxX; x++ {
			if x < len(line) {
				g.SetState(int64(x), int64(y), rune(line[x]))
			} else {
				g.SetState(int64(x), int64(y), blank)
			}
		}
	}
	//printG(g)
	return g
}

func printG(g *DataStruct.Grid[rune]) {
	fmt.Printf("%v", g.StateString(func(rune2 rune) string {
		if rune2 == blank {
			return " "
		}
		return string(rune2)
	}))
}

func getDirRune(index int64) rune {
	if index == 0 {
		return 'v'
	}
	if index == 1 {
		return '>'
	}
	if index == 2 {

		return '^'
	}
	if index == 3 {

		return '<'
	}
	panic(index)
}

func getDirValue(index int64) int64 {
	if index == 0 {
		return 1
	}
	if index == 1 {
		return 0
	}
	if index == 2 {
		return 3
	}
	if index == 3 {

		return 2
	}
	panic(index)
}
