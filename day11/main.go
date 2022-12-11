package main

import (
	"aoc2022/util"
	"gopkg.in/eapache/queue.v1"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("11")
	monkeyLines := util.Get2dString(input, "\n\n", "\n")
	monkeys := make([]*monkey, 0)
	for _, monkeyLine := range monkeyLines {
		monkey := initialMonkey(monkeyLine)
		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 10000; i++ {
		//log.Printf("----------round: %v", i)
		for k, monkey := range monkeys {
			itemNum := monkey.items.Length()
			for j := 0; j < itemNum; j++ {
				//partOne(monkey, monkeys, k)
				partTwo(monkey, monkeys, k)
			}
		}
	}
	for _, monkey := range monkeys {
		log.Printf("%v", monkey.inspectCount)
	}
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(monkey *monkey, monkeys []*monkey, k int) {
	monkey.inspectCount++
	itemLevel := monkey.items.Remove().(int)
	newLevel := monkey.operation(itemLevel)
	passTest := monkey.test(newLevel / 3)
	if passTest {
		monkeys[monkey.ifTrue].items.Add(newLevel / 3)
		log.Printf("%v - %v(%v) -> %v", k, newLevel/3, itemLevel, monkey.ifTrue)
	} else {
		monkeys[monkey.ifFalse].items.Add(newLevel / 3)
		log.Printf("%v - %v(%v) -> %v", k, newLevel/3, itemLevel, monkey.ifFalse)
	}
}

func partTwo(monkey *monkey, monkeys []*monkey, k int) {
	monkey.inspectCount++
	itemLevel := monkey.items.Remove().(int)
	newLevel := monkey.managePanic(monkey.operation(itemLevel))
	passTest := monkey.test(newLevel)
	if passTest {
		monkeys[monkey.ifTrue].items.Add(newLevel)
		//log.Printf("%v - %v(%v) -> %v", k, newLevel, itemLevel, monkey.ifTrue)
	} else {
		monkeys[monkey.ifFalse].items.Add(newLevel)
		//log.Printf("%v - %v(%v) -> %v", k, newLevel, itemLevel, monkey.ifFalse)
	}
}

func initialMonkey(monkeyLine []string) *monkey {
	monkey := &monkey{
		items:     queue.New(),
		operation: nil,
		test:      nil,
		ifTrue:    0,
		ifFalse:   0,
	}
	initialItems(monkeyLine, monkey)
	initialOp(monkeyLine, monkey)
	initTest(monkeyLine, monkey)
	initManagePanic(monkey)
	initialNext(monkey, monkeyLine)
	return monkey
}

func initialNext(monkey *monkey, monkeyLine []string) {
	a := strings.Split(strings.Trim(monkeyLine[4], " "), " ")[5]
	monkey.ifTrue = util.ParseInt(a)
	monkey.ifFalse = util.ParseInt(strings.Split(strings.Trim(monkeyLine[5], " "), " ")[5])
}

func initTest(monkeyLine []string, monkey *monkey) {
	testLine := strings.Split(strings.Trim(monkeyLine[3], " "), " ")
	num := util.ParseInt(testLine[3])
	monkey.test = func(old int) bool {
		return old%num == 0
	}
}
func initManagePanic(monkey *monkey) {
	monkey.managePanic = func(old int) int {
		return old % 9699690
	}
}

func initialOp(monkeyLine []string, monkey *monkey) {
	opLine := strings.Split(strings.Trim(monkeyLine[2], " "), " ")
	v, err := strconv.Atoi(opLine[5])

	if opLine[4] == "+" {
		monkey.operation = func(old int) int {
			return old + v
		}
	} else if opLine[4] == "-" {
		monkey.operation = func(old int) int {
			return old - v
		}
	} else if opLine[4] == "*" {
		monkey.operation = func(old int) int {
			if err != nil {
				return old * old
			}
			return old * v
		}
	} else {
		panic(opLine[4])
	}
}

func initialItems(monkeyLine []string, monkey *monkey) {
	items := strings.Split(strings.Split(monkeyLine[1], ":")[1], ",")
	for _, item := range items {
		monkey.items.Add(util.ParseInt(item))
	}
}

type monkey struct {
	inspectCount int
	items        *queue.Queue
	operation    func(old int) int
	test         func(old int) bool
	managePanic  func(old int) int
	ifTrue       int
	ifFalse      int
}
