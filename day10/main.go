package main

import (
	"aoc2022/util"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("10")
	instructs := util.Get2dString(input, "\n", " ")
	total := partOne(instructs)
	log.Printf("%+v", total)

	tick := 0
	x := 1
	for _, instruct := range instructs {
		if instruct[0] == "noop" {
			tick = tickPartTwo(tick, x)
		} else if instruct[0] == "addx" {

			tick = tickPartTwo(tick, x)
			tick = tickPartTwo(tick, x)
			x += util.ParseInt(instruct[1])
		} else {
			panic(instruct[0])
		}
	}
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func tickPartTwo(tick int, x int) int {
	tick++
	if tick <= x+2 && tick >= x {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	fmt.Printf(" ")
	if tick%40 == 0 {
		tick = 0
		fmt.Printf("\n")
	}
	return tick
}

func partOne(instructs [][]string) int {
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
	return total
}

func ticktock(tick int, x int, total int) (int, int) {
	tick++
	if tick == 20 || tick == 60 || tick == 100 || tick == 140 || tick == 180 || tick == 220 {
		total = x*tick + total
	}
	return tick, total
}
