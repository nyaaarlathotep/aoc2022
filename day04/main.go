package main

import (
	"aoc2022/util"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("04")
	pairs := util.Get2dString(input, "\n", ",")
	partOneTotal := 0
	partTwoTotal := 0
	for _, pair := range pairs {
		left := strings.Split(pair[0], "-")
		right := strings.Split(pair[1], "-")
		leftNum := [2]int{util.ParseInt(left[0]), util.ParseInt(left[1])}
		rightNum := [2]int{util.ParseInt(right[0]), util.ParseInt(right[1])}
		if (leftNum[0] <= rightNum[0] && leftNum[1] >= rightNum[1]) ||
			(leftNum[0] >= rightNum[0] && leftNum[1] <= rightNum[1]) {
			partOneTotal++
		}
		if (leftNum[0] <= rightNum[1] && leftNum[0] >= rightNum[0]) ||
			(leftNum[1] <= rightNum[1] && leftNum[1] >= rightNum[0]) ||
			(leftNum[0] <= rightNum[0] && leftNum[1] >= rightNum[1]) ||
			(leftNum[0] >= rightNum[0] && leftNum[1] <= rightNum[1]) {
			partTwoTotal++
		}
	}
	log.Printf("%+v", partOneTotal)
	log.Printf("%+v", partTwoTotal)

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}
