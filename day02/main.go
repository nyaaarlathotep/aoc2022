package main

import (
	"aoc2022/util"
	"log"
	"time"
)

type elfMove string
type myMove string

const (
	Rock       elfMove = "A"
	Paper      elfMove = "B"
	Scissors   elfMove = "C"
	MyRock     myMove  = "X"
	MyPaper    myMove  = "Y"
	MyScissors myMove  = "Z"
)

func main() {
	start := time.Now()
	input := util.GetInput("02")
	rawMoves := util.Get2dString(input, "\n", " ")
	partOne(rawMoves)
	partTwo(rawMoves)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(rawMoves [][]string) {
	total := 0
	for _, rawMove := range rawMoves {
		elf := parseElfMove(rawMove[0])
		my := parseMyMove(rawMove[1])
		winScore := my.win(elf)
		total = total + winScore + my.moveScore()
	}
	log.Printf("%+v", total)
}

func partTwo(rawMoves [][]string) {
	total := 0
	for _, rawMove := range rawMoves {
		elf := parseElfMove(rawMove[0])
		winScore := scorePartTwo(rawMove[1], elf)
		total = total + winScore
	}
	log.Printf("%+v", total)
}

func parseElfMove(move string) elfMove {
	if move == "A" {
		return Rock
	}
	if move == "B" {
		return Paper
	}
	if move == "C" {
		return Scissors
	}
	log.Fatalln("???")
	return Rock
}

func parseMyMove(move string) myMove {
	if move == "X" {
		return MyRock
	}
	if move == "Y" {
		return MyPaper
	}
	if move == "Z" {
		return MyScissors
	}
	log.Fatalln("???")
	return MyRock
}
func scorePartTwo(move string, move2 elfMove) int {
	if move == "X" {
		if move2 == Rock {
			return 3
		}
		if move2 == Scissors {
			return 2
		}
		if move2 == Paper {
			return 1
		}
		log.Fatalln("???")
		return 0
	}
	if move == "Y" {
		if move2 == Rock {
			return 4
		}
		if move2 == Scissors {
			return 6
		}
		if move2 == Paper {
			return 5
		}
		log.Fatalln("???")
		return 0
	}
	if move == "Z" {
		if move2 == Rock {
			return 8
		}
		if move2 == Scissors {
			return 7
		}
		if move2 == Paper {
			return 9
		}
		log.Fatalln("???")
		return 0
	}
	panic("scorePartTwo")
}

func (my *myMove) win(move elfMove) int {
	if *my == MyRock {
		if move == Rock {
			return 3
		}
		if move == Scissors {
			return 6
		}
		if move == Paper {
			return 0
		}
		log.Fatalln("???")
		return 0
	}

	if *my == MyPaper {
		if move == Rock {
			return 6
		}
		if move == Scissors {
			return 0
		}
		if move == Paper {
			return 3
		}
		log.Fatalln("???")
		return 0
	}
	if *my == MyScissors {
		if move == Rock {
			return 0
		}
		if move == Scissors {
			return 3
		}
		if move == Paper {
			return 6
		}
		log.Fatalln("???")
		return 0
	}
	log.Fatalln("???")
	return 0
}

func (my *myMove) moveScore() int {
	if *my == MyRock {
		return 1
	}
	if *my == MyPaper {
		return 2
	}
	if *my == MyScissors {
		return 3
	}
	panic("moveScore")
}
