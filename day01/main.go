package main

import (
	"aoc2022/util"
	"log"
	"strconv"
)

func main() {
	input := util.GetInput("01")
	lines := util.GetStringSlice(input, "\n")
	cals := make([][]int, 1)
	index := 0
	for _, line := range lines {
		if line == "" {
			index++
			cals = append(cals, make([]int, 0))
			continue
		}
		calNum, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)

		}
		cals[index] = append(cals[index], calNum)
	}
	calSum := make([]int, len(cals))
	for i, cal := range cals {
		total := 0
		for _, num := range cal {
			total += num
		}
		calSum[i] = total
	}
	partOne(calSum)
	partTwo(calSum)
}

func partTwo(calSum []int) {
	maxOne := 0
	maxTwo := 0
	maxThree := 0
	for _, calEach := range calSum {
		if calEach > maxOne {
			maxThree = maxTwo
			maxTwo = maxOne
			maxOne = calEach
			continue
		}
		if calEach > maxTwo {
			maxThree = maxTwo
			maxTwo = calEach
			continue
		}
		if calEach > maxThree {
			maxThree = calEach
			continue
		}
	}
	log.Printf("%+v", maxTwo+maxOne+maxThree)
}

func partOne(calSum []int) {
	max := 0
	for _, calEach := range calSum {
		if calEach > max {
			max = calEach

		}
	}
	log.Printf("%+v", max)
}
