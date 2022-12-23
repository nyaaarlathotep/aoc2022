package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"fmt"
	"golang.org/x/exp/maps"
	"time"
)

const elfPlaceHolder = '#'
const blank = 0

func main() {
	start := time.Now()
	input := util.GetInput("23")
	lines := util.GetStringSlice(input, "\n")
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
	printG(g)
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

			p := getNextLegalP(elf, g)
			fmt.Printf("%+v -> %+v\n", elf, p)
			if _, ok := dest[p]; !ok {
				dest[p] = elf
			} else {
				fmt.Printf("duplicate: %+v\n", p)
				delete(dest, p)
			}
		}
		for p, elf := range dest {
			g.SetState(elf.X, elf.Y, '.')
			g.SetState(p.X, p.Y, elfPlaceHolder)
		}
		elves = maps.Keys(g.StateMapWhere(func(r rune) bool { return r == elfPlaceHolder }))
		printG(g)
	}

	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func getNextLegalP(elf DataStruct.Point, g *DataStruct.Grid[rune]) DataStruct.Point {
	ss := DataStruct.Point{
		X: 13,
		Y: 11,
	}
	if elf == ss {
		fmt.Printf("")
	}
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
		return elf.Add(DataStruct.Directions8[4])
	}
	block = false
	points = [3]DataStruct.Point{elf.Add(DataStruct.Directions8[0]), elf.Add(DataStruct.Directions8[1]), elf.Add(DataStruct.Directions8[7])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[0])
	}
	block = false
	points = [3]DataStruct.Point{elf.Add(DataStruct.Directions8[5]), elf.Add(DataStruct.Directions8[6]), elf.Add(DataStruct.Directions8[7])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[6])
	}
	block = false
	points = [3]DataStruct.Point{elf.Add(DataStruct.Directions8[1]), elf.Add(DataStruct.Directions8[2]), elf.Add(DataStruct.Directions8[3])}
	for _, p := range points {
		r := g.GetPointState(p)
		if r == elfPlaceHolder {
			block = true
			break
		}
	}
	if !block {
		return elf.Add(DataStruct.Directions8[2])
	}
	return elf.Clone()
}
func printG(g *DataStruct.Grid[rune]) {
	fmt.Printf("%v", g.StateString(func(rune2 rune) string {
		if rune2 == blank {
			return "."
		}
		return string(rune2)
	}))
}
