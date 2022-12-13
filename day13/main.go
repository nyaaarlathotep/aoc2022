package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"log"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("13")
	pairs := util.Get2dString(input, "\n\n", "\n")
	partOne(pairs)
	partTwo(pairs)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partTwo(pairs [][]string) {
	packets := make([]packet, 0)
	for _, pair := range pairs {
		packets = append(packets, *parsePacket(pair[0]))
		packets = append(packets, *parsePacket(pair[1]))
	}
	markOne := packet{
		isNum: false,
		num:   0,
		packets: []*packet{
			{
				isNum:   true,
				num:     2,
				packets: nil,
			},
		},
	}
	markTwo := packet{
		isNum: false,
		num:   0,
		packets: []*packet{
			{
				isNum:   true,
				num:     6,
				packets: nil,
			},
		},
	}
	packets = append(packets, markOne)
	packets = append(packets, markTwo)
	res := 1
	sort.Sort(packetList(packets))
	for i, p := range packets {
		if !p.isNum && len(p.packets) == 1 && p.packets[0].isNum && p.packets[0].num == 2 {
			res *= i + 1
		}
		if !p.isNum && len(p.packets) == 1 && p.packets[0].isNum && p.packets[0].num == 6 {
			res *= i + 1
		}
	}
	log.Printf("%v", res)
}

func partOne(pairs [][]string) {
	sum := 0
	for i, pair := range pairs {
		left := parsePacket(pair[0])
		right := parsePacket(pair[1])
		flag := compare(left, right)
		if flag == 0 {
			panic(pairs)
		}
		if flag == 1 {
			sum = sum + i + 1
		}
	}
	log.Printf("%v", sum)
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

type packetList []packet

func (p packetList) Len() int {
	return len(p)
}
func (p packetList) Less(i, j int) bool {
	return compare(&p[i], &p[j]) > 0
}
func (p packetList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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
