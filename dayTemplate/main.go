package main

import (
	"aoc2022/util"
	"log"
	"time"
)


func main() {
	start := time.Now()
	input := util.GetInput("02")
	_ = util.GetStringSlice(input,"\n")
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}
