package main

import (
	"aoc2022/util"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("02")
	_ = util.GetStringSlice(input, "\n")
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
