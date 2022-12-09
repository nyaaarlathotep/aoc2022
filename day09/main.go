package main

import (
	"aoc2022/util"
	"log"
	"math"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("09")
	instructs := util.Get2dString(input, "\n", " ")

	head := &point{
		x: 0,
		y: 0,
	}
	tail := &point{
		x: 0,
		y: 0,
	}
	ropeKonts := make([]*point, 10)
	ropeKonts[0] = head
	for i := 1; i < 10; i++ {
		ropeKonts[i] = &point{
			x: 0,
			y: 0,
		}
	}

	traces := make(map[point]bool)
	tracesPartTwo := make(map[point]bool)
	for _, line := range instructs {
		for i := 0; i < util.ParseInt(line[1]); i++ {
			traces[*tail] = false
			tracesPartTwo[*ropeKonts[9]] = false
			if line[0] == "D" {
				head.y--
			} else if line[0] == "U" {
				head.y++
			} else if line[0] == "L" {
				head.x--
			} else if line[0] == "R" {
				head.x++
			} else {
				panic("???")
			}
			moveTail(head, tail)
			for j := 0; j < len(ropeKonts)-1; j++ {
				moveTail(ropeKonts[j], ropeKonts[j+1])
			}
		}
	}
	traces[*tail] = false
	tracesPartTwo[*ropeKonts[9]] = false
	log.Printf("%+v", len(traces))
	log.Printf("%+v", len(tracesPartTwo))
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)

}

func moveTail(head *point, tail *point) {
	if math.Abs(float64(head.x-tail.x)) <= 1 && math.Abs(float64(head.y-tail.y)) <= 1 {
		return
	} else if math.Abs(float64(head.x-tail.x))+math.Abs(float64(head.y-tail.y)) == 2 {
		tail.x = tail.x + (head.x-tail.x)/2
		tail.y = tail.y + (head.y-tail.y)/2
	} else {
		if math.Signbit(float64(head.x - tail.x)) {
			tail.x--
		} else {
			tail.x++
		}
		if math.Signbit(float64(head.y - tail.y)) {
			tail.y--
		} else {
			tail.y++
		}
	}
}

type point struct {
	x int
	y int
}
