package main

import (
	"aoc2022/util"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("03")
	rucksacks := util.GetStringSlice(input, "\n")
	indices := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	partOne(rucksacks, indices)

	totalPartTwo := 0
	for i := 0; i < len(rucksacks)/3; i++ {
		group := rucksacks[3*i : 3*i+3]
		for _, r := range group[0] {
			if strings.Contains(group[1], string(r)) &&
				strings.Contains(group[2], string(r)) {
				totalPartTwo += strings.Index(indices, string(r)) + 1
				break
			}
		}
	}
	log.Printf("%v", totalPartTwo)

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(rucksacks []string, indices string) {
	totalPartOne := 0
	for _, rucksack := range rucksacks {
		partOne, partTwo := rucksack[0:len(rucksack)/2], rucksack[len(rucksack)/2:]
		for _, r := range partOne {
			if strings.Contains(partTwo, string(r)) {
				totalPartOne += strings.Index(indices, string(r)) + 1
				break
			}
		}
	}
	log.Printf("%+v", totalPartOne)
}
