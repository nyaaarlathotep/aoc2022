package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("22")
	lines := util.GetStringSlice(input, "\n")
	g := DataStruct.NewGrid[rune](int64(len(lines[0])), int64(len(lines)), DataStruct.Directions4)
	for y, line := range lines {
		for x, r := range line {
			g.SetState(int64(x), int64(y), r)
		}
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
