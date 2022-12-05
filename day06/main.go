package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("06")
	stacks := initialStack()

	rawInstructs := util.Get2dString(input, "\n", " ")
	instructs := make([][3]int, 0)
	for _, rawInstruct := range rawInstructs {
		instructs = append(instructs,
			[3]int{util.ParseInt(rawInstruct[1]), util.ParseInt(rawInstruct[3]), util.ParseInt(rawInstruct[5])})
	}

	for _, instruct := range instructs {
		for i := 0; i < instruct[0]; i++ {
			stacks[instruct[2]-1].Push(stacks[instruct[1]-1].Pop())
			//for _, s := range stacks {
			//	s.Traverse()
			//}
		}
	}

	for _, stack := range stacks {
		log.Printf("%v", stack.Pop())
	}

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func initialTestStack() []DataStruct.Stack {
	initialStacks := make([]DataStruct.Stack, 3)
	initialStacks[0] = DataStruct.NewStack("Z", "N")
	initialStacks[1] = DataStruct.NewStack("M", "C", "D")
	initialStacks[2] = DataStruct.NewStack("P")
	return initialStacks
}

func initialStack() []DataStruct.Stack {
	initialStacks := make([]DataStruct.Stack, 9)
	initialStacks[0] = DataStruct.NewStack("D", "T", "W", "F", "J", "S", "H", "N")
	initialStacks[1] = DataStruct.NewStack("H", "R", "P", "Q", "T", "N", "B", "G")
	initialStacks[2] = DataStruct.NewStack("L", "Q", "V")
	initialStacks[3] = DataStruct.NewStack("N", "B", "S", "W", "R", "Q")
	initialStacks[4] = DataStruct.NewStack("N", "D", "F", "T", "V", "M", "B")
	initialStacks[5] = DataStruct.NewStack("M", "D", "B", "V", "H", "T", "R")
	initialStacks[6] = DataStruct.NewStack("D", "B", "Q", "J")
	initialStacks[7] = DataStruct.NewStack("D", "N", "J", "V", "R", "Z", "H", "Q")
	initialStacks[8] = DataStruct.NewStack("B", "N", "H", "M", "S")
	return initialStacks
}
