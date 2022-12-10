package main

import (
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("10")
	instructs := util.Get2dString(input, "\n", " ")
	tick := 0
	x := 1
	total := 0
	for _, instruct := range instructs {
		if instruct[0] == "noop" {
			tick, total = ticktock(tick, x, total)
		} else if instruct[0] == "addx" {

			tick, total = ticktock(tick, x, total)
			tick, total = ticktock(tick, x, total)
			x += util.ParseInt(instruct[1])
		} else {
			panic(instruct[0])
		}
	}
	log.Printf("%+v", total)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func ticktock(tick int, x int, total int) (int, int) {
	tick++
	if tick == 20 || tick == 60 || tick == 100 || tick == 140 || tick == 180 || tick == 220 {
		log.Printf("tick: %+v\n", tick)
		log.Printf("x: %+v\n", x)
		log.Printf("x*tick: %+v\n", x*tick)
		total = x*tick + total
	}
	return tick, total
}
