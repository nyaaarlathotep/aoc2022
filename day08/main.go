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

	highTreeMark, count := partOne(trees)
	highScore := partTwo(highTreeMark, trees)

	log.Printf("%+v", count)
	log.Printf("%+v", highScore)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(trees [][]int) ([][]int, int) {
	highTreeMark := make([][]int, len(trees))
	highTreeMark[0] = make([]int, len(trees[0]))
	lookFromSide(trees, highTreeMark)
	highTreeMark[len(trees)-1] = make([]int, len(trees[0]))
	lookFromUp(trees, highTreeMark)
	count := 0
	for _, line := range highTreeMark {
		for _, vis := range line {
			if vis == 0 {
				count++
			}
		}
	}
	return highTreeMark, count
}

func partTwo(highTreeMark [][]int, trees [][]int) int {
	highScore := 0
	for i := range highTreeMark {
		for j := range highTreeMark[i] {
			if highTreeMark[i][j] == 0 {
				if i == 0 || j == 0 || i == len(highTreeMark) || j == len(highTreeMark[0]) {
					continue
				}
				left, right, up, down := j, j, i, i
				for left > 0 {
					left--
					if trees[i][left] >= trees[i][j] {
						break
					}
				}

				for right < len(highTreeMark)-1 {
					right++
					if trees[i][right] >= trees[i][j] {
						break
					}
				}

				for up > 0 {
					up--
					if trees[up][j] >= trees[i][j] {
						break
					}
				}
				for down < len(highTreeMark[0])-1 {
					down++
					if trees[down][j] >= trees[i][j] {
						break
					}
				}
				score := (i - up) * (down - i) * (j - left) * (right - j)
				if score > highScore {
					highScore = score
				}
			}
		}
	}
	return highScore
}

func lookFromUp(trees [][]int, highTreeMark [][]int) {
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
}

func lookFromSide(trees [][]int, highTreeMark [][]int) {
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
}
