package main

import (
	"aoc2022/DataStruct"
	"aoc2022/c"
	"aoc2022/util"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("14")
	lines := util.GetStringSlice(input, "\n")
	g := DataStruct.NewGrid[rune](360, 360, DataStruct.Directions4)
	var maxY int64 = 0
	for _, line := range lines {
		pointStr := strings.Split(line, " -> ")
		for i := 0; i < len(pointStr)-1; i++ {
			s := strings.Split(pointStr[i], ",")
			e := strings.Split(pointStr[i+1], ",")
			sx := util.ParseInt64(s[0]) - 320
			ex := util.ParseInt64(e[0]) - 320
			sy := util.ParseInt64(s[1])
			ey := util.ParseInt64(e[1])
			for y, x := sy, sx; y != ey+c.POne(ey-sy) || x != ex+c.POne(ex-sx); y, x = y+c.POne(ey-sy), x+c.POne(ex-sx) {
				g.SetState(x, y, '#')
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	floor := maxY + 2
	for x := 0; x < int(g.XLen()); x++ {
		g.SetState(int64(x), floor, '#')
	}
	sand := DataStruct.Point{
		X: 180,
		Y: 0,
	}
	partOne(sand, g)
	partTwo(g, sand)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partTwo(g *DataStruct.Grid[rune], sand DataStruct.Point) {
	for true {
		if blocked(g, sand.Y, sand.X) {
			goto end
		}
		if g.GetState(sand.Y+1, sand.X) == 0 {
			sand.Y++
			continue
		}
		if blocked(g, sand.Y+1, sand.X) {
			if validPos(g, sand.Y+1, sand.X) {
				g.SetState(sand.X, sand.Y, 'o')
				sand.X = 180
				sand.Y = 0
				continue
			}
			if !blocked(g, sand.Y+1, sand.X-1) {
				sand.X--
				sand.Y++
				continue
			}
			if !blocked(g, sand.Y+1, sand.X+1) {
				sand.X++
				sand.Y++
				continue
			} else {
				panic("?")
			}
		}
	}
end:
	fmt.Printf("%v", g.StateString(func(r rune) string {
		if r == 0 {
			return "."
		}
		return string(r)
	}))
	sandPos := g.StateMapWhere(func(rune rune) bool { return rune == 'o' })
	log.Printf("%v", len(sandPos))
}

func partOne(sand DataStruct.Point, g *DataStruct.Grid[rune]) {
	for true {
		if sand.Y+1 >= g.YLen() {
			goto end
		}
		if g.GetState(sand.Y+1, sand.X) == 0 {
			sand.Y++
			continue
		}
		if blocked(g, sand.Y+1, sand.X) {
			if validPos(g, sand.Y+1, sand.X) {
				g.SetState(sand.X, sand.Y, 'o')
				sand.X = 100
				sand.Y = 0
				continue
			}
			if !blocked(g, sand.Y+1, sand.X-1) {
				sand.X--
				sand.Y++
				continue
			}
			if !blocked(g, sand.Y+1, sand.X+1) {
				sand.X++
				sand.Y++
				continue
			} else {
				fmt.Printf("%v", g.StateString(func(r rune) string {
					if r == 0 {
						return "."
					}
					return string(r)
				}))
				panic("?")
			}
		}
	}
end:
	fmt.Printf("%v", g.StateString(func(r rune) string {
		if r == 0 {
			return "."
		}
		return string(r)
	}))
	sandPos := g.StateMapWhere(func(rune rune) bool { return rune == 'o' })
	log.Printf("%v", len(sandPos))
}

func blocked(g *DataStruct.Grid[rune], y, x int64) bool {
	return isSand(g, y, x) || isRock(g, y, x)
}

func isRock(g *DataStruct.Grid[rune], y int64, x int64) bool {
	return g.GetState(y, x) == '#'
}

func isSand(g *DataStruct.Grid[rune], y int64, x int64) bool {
	return g.GetState(y, x) == 'o'
}

func validPos(g *DataStruct.Grid[rune], yy int64, x int64) bool {
	return blocked(g, yy, x-1) && blocked(g, yy, x) && blocked(g, yy, x+1)
}
