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
	g := DataStruct.NewGrid[rune](200, 200, DataStruct.Directions4)
	for _, line := range lines {
		pointStr := strings.Split(line, " -> ")
		for i := 0; i < len(pointStr)-1; i++ {
			s := strings.Split(pointStr[i], ",")
			e := strings.Split(pointStr[i+1], ",")
			sx := util.ParseInt64(s[0]) - 400
			ex := util.ParseInt64(e[0]) - 400
			sy := util.ParseInt64(s[1])
			ey := util.ParseInt64(e[1])
			for y, x := sy, sx; y != ey || x != ex; y, x = y+c.POne(ey-sy), x+c.POne(ex-sx) {
				log.Printf("(%v,%v)", x, y)
				g.SetState(x, y, '#')
			}
			g.SetState(ex, ey, '#')
			log.Printf("(%v,%v)", ex, ey)
		}
	}

	sand := DataStruct.Point{
		X: 100,
		Y: 0,
	}
	for true {
		fmt.Printf("%v", g.StateString(func(r rune) string {
			if r == 0 {
				return "."
			}
			return string(r)
		}))
		if g.GetState(sand.X, sand.Y+1) == '0' {
			sand.Y++
			if sand.Y > g.YLen() {
				goto end
			}
			continue
		}
		if blocked(g, sand.X, sand.Y+1) {
			if validPos(g, sand.X, sand.Y) {
				g.SetState(sand.X, sand.Y, 'o')
				sand.X = 100
				sand.Y = 0
				continue
			}
			if !blocked(g, sand.X-1, sand.Y+1) {
				sand.X--
				sand.Y++
				continue
			}
			if !blocked(g, sand.X+1, sand.Y+1) {
				sand.X++
				sand.Y++
				continue
			}
		}
	}
end:
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func blocked(g *DataStruct.Grid[rune], x, y int64) bool {
	return g.GetState(x, y) == 'o' || g.GetState(x, y) == '#'
}

func validPos(g *DataStruct.Grid[rune], x int64, y int64) bool {
	return blocked(g, x-1, y+1) && blocked(g, x, y+1) && blocked(g, x+1, y+1)
}
