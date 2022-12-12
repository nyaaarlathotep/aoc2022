package main

import (
	"aoc2022/util"
	"log"
	"math"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("12")
	heightMap := util.Get2dString(input, "\n", "")
	heightRuneMap := initialHeightRuneMap(heightMap)
	stepsMap := initialStepMap(heightMap)
	steps := 0
	x := 0
	y := 0
	fewest := math.MaxInt
	climb(heightRuneMap, steps, &stepsMap, x, y, &fewest)
	log.Println(fewest)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

// 468 wrong
func climb(heightRuneMap [][]rune, steps int, m *[][]int, x int, y int, fewest *int) {

	stepsMap := *m
	if stepsMap[x][y] != -1 && stepsMap[x][y] <= steps {
		return
	}

	stepsMap[x][y] = steps
	if heightRuneMap[x][y] == 'E' {
		if steps < *fewest {
			*fewest = steps
		}
		return
	}
	if heightRuneMap[x][y] == 'S' {
		heightRuneMap[x][y] = 'a'
	}
	if x > 0 {
		r := heightRuneMap[x-1][y]
		if r == 'E' {
			r = 'z'
		}
		if r-heightRuneMap[x][y] <= 1 {
			climb(heightRuneMap, steps+1, m, x-1, y, fewest)
		}
	}
	if x < len(heightRuneMap)-1 && heightRuneMap[x+1][y]-heightRuneMap[x][y] <= 1 {
		r := heightRuneMap[x+1][y]
		if r == 'E' {
			r = 'z'
		}
		if r-heightRuneMap[x][y] <= 1 {
			climb(heightRuneMap, steps+1, m, x+1, y, fewest)
		}
	}
	if y > 0 && heightRuneMap[x][y-1]-heightRuneMap[x][y] <= 1 {
		r := heightRuneMap[x][y-1]
		if r == 'E' {
			r = 'z'
		}
		if r-heightRuneMap[x][y] <= 1 {
			climb(heightRuneMap, steps+1, m, x, y-1, fewest)
		}
	}
	if y < len(heightRuneMap[0])-1 && heightRuneMap[x][y+1]-heightRuneMap[x][y] <= 1 {
		r := heightRuneMap[x][y+1]
		if r == 'E' {
			r = 'z'
		}
		if r-heightRuneMap[x][y] <= 1 {
			climb(heightRuneMap, steps+1, m, x, y+1, fewest)
		}
	}
}

func initialStepMap(heightMap [][]string) [][]int {
	stepsMap := make([][]int, len(heightMap))
	for i := range heightMap {
		stepsMap[i] = make([]int, len(heightMap[0]))
		for j := range heightMap[0] {
			stepsMap[i][j] = -1
		}
	}
	return stepsMap
}

func initialHeightRuneMap(heightMap [][]string) [][]rune {
	stepsMap := make([][]rune, len(heightMap))
	for i := range heightMap {
		stepsMap[i] = make([]rune, len(heightMap[0]))
		for j, s := range heightMap[i] {
			stepsMap[i][j] = rune(s[0])
		}
	}
	return stepsMap
}
