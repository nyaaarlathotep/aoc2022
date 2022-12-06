package main

import (
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	start := time.Now()
	raw := util.GetInput("06")
	partOne(raw)
	partTwo(raw)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(raw string) {
	flag := false
	for i := 3; i < len(raw); i++ {
		for j := i - 3; j < i; j++ {
			for k := j + 1; k <= i; k++ {
				if raw[k] == raw[j] {
					flag = true
					break
				}
			}
		}
		if flag {
			flag = false
		} else {
			log.Printf("%+v", i+1)
			break
		}
	}
}
func partTwo(raw string) {
	flag := false
	for i := 13; i < len(raw); i++ {
		for j := i - 13; j < i; j++ {
			for k := j + 1; k <= i; k++ {
				if raw[k] == raw[j] {
					flag = true
					break
				}
			}
		}
		if flag {
			flag = false
		} else {
			log.Printf("%+v", i+1)
			break
		}
	}
}
