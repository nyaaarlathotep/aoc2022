package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("17")
	winds := util.GetStringSlice(input, "")
	g := DataStruct.NewGrid[rune](7, 7600, DataStruct.Directions4)

	rocks := []rock{{
		points: []DataStruct.Point{{X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: 0}},
		height: 1,
	}, {
		points: []DataStruct.Point{{X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 3, Y: 2}, {X: 3, Y: 0}},
		height: 3,
	}, {
		points: []DataStruct.Point{{X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 4, Y: 1}, {X: 4, Y: 2}},
		height: 3,
	}, {
		points: []DataStruct.Point{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
		height: 4,
	}, {
		points: []DataStruct.Point{{X: 2, Y: 0}, {X: 3, Y: 0}, {X: 2, Y: 1}, {X: 3, Y: 1}},
		height: 2,
	}}
	windIndex := 0
	height := int64(3)
	for rockIndex := int64(0); rockIndex < 2022; rockIndex++ {
		rockNow := rocks[rockIndex%5]
		heightNow := height
		xOfAll := int64(0)
		for true {
			//fmt.Println("height!!!!", heightNow)
			wind := getWind(winds, &windIndex)
			xNow := getXMove(wind)
			crash := getCrash(rockNow, xOfAll+xNow, g, heightNow)
			if !crash {
				xOfAll += xNow
			}
			heightNow--
			crash = getCrash(rockNow, xOfAll, g, heightNow)
			if crash {
				heightNow++
				for _, p := range rockNow.points {
					g.SetState(p.X+xOfAll, p.Y+heightNow, '#')
				}
				if heightNow+rockNow.height+3 > height {
					height = heightNow + rockNow.height + 3
				}
				//fmt.Printf("%v", g.StateString(func(r rune) string {
				//	if r == 0 {
				//		return "."
				//	}
				//	return string(r)
				//}))
				//fmt.Println("-----------")
				break
			}
		}
	}
	fmt.Println(height - 3)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func getWind(winds []string, windIndex *int) string {
	wind := winds[*windIndex]
	*windIndex++
	*windIndex = *windIndex % len(winds)
	return wind
}

func getCrash(rockNow rock, xMove int64, g *DataStruct.Grid[rune], heightNow int64) bool {
	for _, p := range rockNow.points {
		xAfterWind := p.X + xMove
		if heightNow < 0 || xAfterWind < 0 || xAfterWind >= g.XLen() {
			return true
		}
		if g.GetState(p.Y+heightNow, xAfterWind) == '#' {
			return true
		}
	}
	return false
}

func getXMove(wind string) int64 {
	if wind == ">" {
		return 1
	} else {
		return -1
	}
}

type rock struct {
	points []DataStruct.Point
	height int64
}
