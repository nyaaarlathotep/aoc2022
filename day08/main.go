package main

import (
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("08")
	trees := util.TwoDStringToInt(util.Get2dString(input, "\n", ""))
	highTreeMark := make([][]int, len(trees))
	highTreeMark[0] = make([]int, len(trees[0]))
	for i := 1; i < len(trees)-1; i++ {
		highestIndex := 0
		highTreeMark[i] = make([]int, len(trees[0]))
		for j := 1; j < len(trees[0])-1; j++ {
			if trees[i][j] > trees[i][highestIndex] {
				highestIndex = j
				highTreeMark[i][j] = 0
			} else {
				highTreeMark[i][j] = 1
			}
		}
		highestIndexOtherSide := len(trees[0]) - 1
		for j := len(trees[0]) - 2; j > highestIndex; j-- {
			if trees[i][j] > trees[i][highestIndexOtherSide] {
				highestIndexOtherSide = j
				highTreeMark[i][j] = 0
			}
		}
	}
	highTreeMark[len(trees)-1] = make([]int, len(trees[0]))

	for i := 1; i < len(trees[0])-1; i++ {
		highestIndex := 0
		for j := 1; j < len(trees)-1; j++ {
			if trees[j][i] > trees[highestIndex][i] {
				highestIndex = j
				highTreeMark[j][i] = 0
			}
		}
		highestIndexOtherSide := len(trees) - 1
		for j := len(trees) - 2; j > highestIndex; j-- {
			if trees[j][i] > trees[highestIndexOtherSide][i] {
				highestIndexOtherSide = j
				highTreeMark[j][i] = 0
			}
		}
	}
	//util.Print2DInt(trees, " ")
	//util.Print2DInt(highTreeMark, " ")
	count := 0
	for _, line := range highTreeMark {
		for _, vis := range line {
			if vis == 0 {
				count++
			}
		}
	}
	log.Printf("%+v", count)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}
