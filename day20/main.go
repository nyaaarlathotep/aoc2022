package main

import (
	"aoc2022/util"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	lines := util.GetLines("20")
	mixOrder := partOne(lines)
	partTwo(lines, mixOrder)

	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func partTwo(lines []string, mixOrder []*link) {
	zeroLink, startLink, length := initial(lines)
	p := startLink
	for {
		p.partTwoNum = p.num * 811589153
		p = p.right
		if p == startLink {
			break
		}
	}
	for i := 0; i < 10; i++ {
		mixTwo(zeroLink, length, mixOrder)
	}
	printPartTwo(startLink)
	fmt.Println(getGroveCoordinates(zeroLink))
}

func printPartTwo(zeroLink *link) *link {
	p := zeroLink.right
	fmt.Printf("%v", zeroLink.partTwoNum)
	for p.num != zeroLink.num {
		fmt.Printf(" -> %v", p.partTwoNum)
		p = p.right
	}
	fmt.Println()
	return p
}

func partOne(lines []string) []*link {
	zeroLink, startLink, length := initial(lines)
	mixOrder := mixOne(startLink, length)
	fmt.Println(getNumAt(1000, length, zeroLink) + getNumAt(2000, length, zeroLink) + getNumAt(3000, length, zeroLink))
	return mixOrder
}

func mixOne(startLink *link, length int) []*link {
	p := startLink
	moveCount := 0
	mixOrder := make([]*link, 0)
	for {
		p = p.right
		movingP := p.left
		if movingP.moved {
			continue
		}
		movingP.moved = true
		moveCount++
		mixOrder = append(mixOrder, movingP)
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
	return mixOrder
}

func mixTwo(startLink *link, length int, mixOrder []*link) {
	for _, order := range mixOrder {
		movingP := startLink
		for {
			if movingP.num == order.num {
				break
			}
			movingP = movingP.right
		}
		dest := movingP
		moves := movingP.partTwoNum % (length - 1)
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
	}
}

func getNumAt(i int, length int, zeroLink *link) int {
	count := i % length
	p := zeroLink
	for i = 0; i < count; i++ {
		p = p.right
	}
	return p.num
}

func getPartTwoNumAt(i int, length int, zeroLink *link) int {
	count := i % length
	p := zeroLink
	for i = 0; i < count; i++ {
		p = p.right
	}
	return p.partTwoNum
}

func getGroveCoordinates(zero *link) int {
	gc := []int{}

	cur := zero
	//for cur.partTwoNum != 0 {
	//	cur = cur.right
	//}

	for len(gc) < 3 {
		for i := 0; i < 1000; i++ {
			cur = cur.right
		}
		gc = append(gc, cur.partTwoNum)
	}
	return gc[0] + gc[1] + gc[2]
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
	left       *link
	right      *link
	num        int
	partTwoNum int
	moveNum    int
	moved      bool
}
