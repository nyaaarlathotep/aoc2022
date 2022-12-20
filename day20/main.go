package main

import (
	"aoc2022/util"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	lines := util.GetLines("20")
	zeroLink, startLink, length := initial(lines)
	p := startLink
	moveCount := 0
	for {
		//printLink(zeroLink)
		p = p.right
		movingP := p.left
		if movingP.moved {
			continue
		}
		movingP.moved = true
		moveCount++
		dest := movingP
		moves := movingP.num % (length - 1)
		if moves > 0 {
			for i := 0; i < moves; i++ {
				dest = dest.right
			}
			addToRight(movingP, dest)
		} else if moves < 0 {
			for i := 0; i > moves; i-- {
				dest = dest.left
			}
			addToLeft(movingP, dest)
		}
		if moveCount == length {
			break
		}
	}
	//printLink(zeroLink)
	fmt.Println(getNumAt(1000, length, zeroLink) + getNumAt(2000, length, zeroLink) + getNumAt(3000, length, zeroLink))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func getNumAt(i int, length int, zeroLink *link) int {
	count := i % length
	p := zeroLink
	for i = 0; i < count; i++ {
		p = p.right
	}
	return p.num
}

func addToLeft(movingP *link, dest *link) {
	movingP.right.left = movingP.left
	movingP.left.right = movingP.right
	movingP.left = dest.left
	dest.left.right = movingP
	movingP.right = dest
	dest.left = movingP
}

func addToRight(movingP *link, dest *link) {
	movingP.right.left = movingP.left
	movingP.left.right = movingP.right
	movingP.right = dest.right
	dest.right.left = movingP
	movingP.left = dest
	dest.right = movingP

}

func printLink(zeroLink *link) {
	p := zeroLink.right
	fmt.Printf("%v", zeroLink.num)
	for p.num != zeroLink.num {
		fmt.Printf(" -> %v", p.num)
		p = p.right
	}
	fmt.Println()
}

func initial(lines []string) (*link, *link, int) {
	length := 0
	dummyHead := &link{
		left:  nil,
		right: nil,
		moved: false,
	}
	var zeroLink *link
	p := dummyHead
	for _, line := range lines {
		length++
		num := util.ParseInt(line)
		p.right = &link{
			left:  p,
			right: nil,
			num:   num,
			moved: false,
		}
		p = p.right
		if num == 0 {
			zeroLink = p
		}
	}
	p.right = dummyHead.right
	dummyHead.right.left = p
	return zeroLink, dummyHead.right, length
}

type link struct {
	left  *link
	right *link
	num   int
	moved bool
}
