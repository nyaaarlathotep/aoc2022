package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"fmt"
	"golang.org/x/exp/maps"
	"math"
	"time"
)

const elfPlaceHolder = '#'
const blank = 0
const margin = 100

func main() {
	start := time.Now()
	input := util.GetInput("23")
	lines := util.GetStringSlice(input, "\n")

	partOne(lines)
	partTwo(lines)
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func partOne(lines []string) {
	g := DataStruct.NewGrid[rune](int64(len(lines[0]))+20, int64(len(lines))+20, DataStruct.Directions8)

	elves := make([]DataStruct.Point, 0)

	for i, l := range lines {
		y := i + 10
		for j, r := range l {
			x := j + 10
			if r == '.' {
				r = blank
			}
			g.SetState(int64(x), int64(y), r)
			if r == elfPlaceHolder {
				elves = append(elves, DataStruct.Point{
					X: int64(x),
					Y: int64(y),
				})
			}
		}
	}
	dirs := []func(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool){goNorth, goSouth, goWest, goEast}
	dirIndex := 0
	//printG(g)
	for i := 0; i < 10; i++ {
		dest := make(map[DataStruct.Point]DataStruct.Point)
		for _, elf := range elves {
			neighbours := g.Neighbours(elf)
			move := false
			for _, p := range neighbours {
				r := g.GetPointState(p)
				if r == elfPlaceHolder {
					move = true
					break
				}
			}
			if !move {
				continue
			}

			p := getNextLegalP(elf, g, dirIndex, dirs)
			//fmt.Printf("%+v -> %+v\n", elf, p)
			if _, ok := dest[p]; !ok {
				dest[p] = elf
			} else {
				//fmt.Printf("duplicate: %+v\n", p)
				delete(dest, p)
			}
		}
		for p, elf := range dest {
			g.SetState(elf.X, elf.Y, '.')
			g.SetState(p.X, p.Y, elfPlaceHolder)
		}
		elves = maps.Keys(g.StateMapWhere(func(r rune) bool { return r == elfPlaceHolder }))
		//printG(g)
		dirIndex++
	}
	minX := int64(math.MaxInt64)
	maxX := int64(0)
	minY := int64(math.MaxInt64)
	maxY := int64(0)
	for _, p := range elves {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	fmt.Printf("%v\n", (maxX-minX+1)*(maxY-minY+1)-int64(len(elves)))
}

func partTwo(lines []string) {
	g := DataStruct.NewGrid[rune](int64(len(lines[0]))+2*margin, int64(len(lines))+2*margin, DataStruct.Directions8)

	elves := make([]DataStruct.Point, 0)

	for i, l := range lines {
		y := i + margin
		for j, r := range l {
			x := j + margin
			if r == '.' {
				r = blank
			}
			g.SetState(int64(x), int64(y), r)
			if r == elfPlaceHolder {
				elves = append(elves, DataStruct.Point{
					X: int64(x),
					Y: int64(y),
				})
			}
		}
	}
	dirs := []func(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool){goNorth, goSouth, goWest, goEast}
	dirIndex := 0
	//printG(g)
	rounds := 0
	for ; ; rounds++ {
		dest := make(map[DataStruct.Point]DataStruct.Point)
		for _, elf := range elves {
			neighbours := g.Neighbours(elf)
			move := false
			for _, p := range neighbours {
				r := g.GetPointState(p)
				if r == elfPlaceHolder {
					move = true
					break
				}
			}
			if !move {
				continue
			}

			p := getNextLegalP(elf, g, dirIndex, dirs)
			//fmt.Printf("%+v -> %+v\n", elf, p)
			if _, ok := dest[p]; !ok {
				dest[p] = elf
			} else {
				//fmt.Printf("duplicate: %+v\n", p)
				delete(dest, p)
			}
		}
		if len(dest) == 0 {
			rounds++
			break
		}
		for p, elf := range dest {
			g.SetState(elf.X, elf.Y, '.')
			g.SetState(p.X, p.Y, elfPlaceHolder)
		}
		elves = maps.Keys(g.StateMapWhere(func(r rune) bool { return r == elfPlaceHolder }))
		//printG(g)
		dirIndex++
	}
	fmt.Printf("%v\n", rounds)
}

func getNextLegalP(elf DataStruct.Point, g *DataStruct.Grid[rune], dirIndex int,
	fs []func(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool)) DataStruct.Point {
	ss := DataStruct.Point{
		X: 13,
		Y: 11,
	}
	if elf == ss {
		fmt.Printf("")
	}
	for i := 0; i < 4; i++ {
		fi := (dirIndex + i) % 4
		if point, done := fs[fi](elf, g); done {
			return point
		}
	}
	return elf.Clone()
}

func goNorth(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool) {
	points := [3]DataStruct.Point{elf.Add(DataStruct.Directions8[3]), elf.Add(DataStruct.Directions8[4]), elf.Add(DataStruct.Directions8[5])}
	block := false
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[4]), true
	}
	return DataStruct.Point{}, false
}

func goSouth(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool) {
	block := false
	points := [3]DataStruct.Point{elf.Add(DataStruct.Directions8[0]), elf.Add(DataStruct.Directions8[1]), elf.Add(DataStruct.Directions8[7])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[0]), true
	}
	return DataStruct.Point{}, false
}

func goWest(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool) {
	block := false
	points := [3]DataStruct.Point{elf.Add(DataStruct.Directions8[5]), elf.Add(DataStruct.Directions8[6]), elf.Add(DataStruct.Directions8[7])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[6]), true
	}
	return DataStruct.Point{}, false
}
func goEast(elf DataStruct.Point, g *DataStruct.Grid[rune]) (DataStruct.Point, bool) {
	block := false
	points := [3]DataStruct.Point{elf.Add(DataStruct.Directions8[1]), elf.Add(DataStruct.Directions8[2]), elf.Add(DataStruct.Directions8[3])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[2]), true
	}
	return DataStruct.Point{}, false
}

func printG(g *DataStruct.Grid[rune]) {
	fmt.Printf("%v", g.StateString(func(rune2 rune) string {
		if rune2 == blank {
			return "."
		}
		return string(rune2)
	}))
}
