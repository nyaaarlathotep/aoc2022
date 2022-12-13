package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"log"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("13")
	pairs := util.Get2dString(input, "\n\n", "\n")
	indexes := make([]int, 0)
	sum := 0
	for i, pair := range pairs {
		left := parsePacket(pair[0])
		right := parsePacket(pair[1])
		flag := compare(left, right)
		if flag == 0 {
			panic(pairs)
		}
		if flag == 1 {
			indexes = append(indexes, i+1)
			sum = sum + i + 1
		}
	}
	log.Printf("%v", indexes)
	log.Printf("%v", sum)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func compare(left *packet, right *packet) int {
	if left.isNum && right.isNum {
		if left.num < right.num {
			return 1
		}
		if left.num > right.num {
			return -1
		}
		if left.num == right.num {
			return 0
		}
	}
	if !left.isNum && !right.isNum {
		for i := range left.packets {
			if i >= len(right.packets) {
				return -1
			}
			res := compare(left.packets[i], right.packets[i])
			if res != 0 {
				return res
			}
		}
		if len(left.packets) != len(right.packets) {
			return 1
		} else {
			return 0
		}
	}
	if left.isNum {
		left = transToList(left)
	}
	if right.isNum {
		right = transToList(right)
	}
	return compare(left, right)
}

func transToList(left *packet) *packet {
	return &packet{
		isNum:   false,
		num:     0,
		packets: []*packet{left},
	}
}

func parsePacket(packetStr string) *packet {
	pack := &packet{
		isNum:   false,
		num:     0,
		packets: make([]*packet, 0),
	}
	packets := spiltIgnore(packetStr[1:len(packetStr)-1], ",", "[", "]")
	for _, innerPacketStr := range packets {
		var p *packet
		if len(innerPacketStr) == 0 {
			return pack
		}
		if innerPacketStr[0] == '[' {
			p = parsePacket(innerPacketStr)
		} else {
			p = &packet{
				isNum:   true,
				num:     util.ParseInt64(innerPacketStr),
				packets: nil,
			}
		}
		pack.packets = append(pack.packets, p)
	}
	return pack
}

type packet struct {
	isNum   bool
	num     int64
	packets []*packet
}

func spiltIgnore(str string, separator string, l string, r string) []string {
	res := make([]string, 0)
	index := 0
	stack := DataStruct.NewStack()
	for i := range str {
		if string(str[i]) == l {
			stack.Push(false)
			continue
		}
		if string(str[i]) == r {
			stack.Pop()
			continue
		}
		if stack.GetLength() != 0 {
			continue
		}
		if string(str[i]) == separator {
			res = append(res, str[index:i])
			index = i + 1
		}
	}
	res = append(res, str[index:])
	return res
}
