package main

import (
	"aoc2022/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("25")
	lines := util.GetStringSlice(input, "\n")
	total := 0
	for _, line := range lines {
		thisNum := 0
		for i, r := range line {
			num := parseRune(r)
			thisNum += num * int(math.Pow(5, float64(len(line)-i-1)))
		}
		total += thisNum
	}
	res := strings.Builder{}
	get5Num(total, &res)
	total5Str := res.String()
	numSlice := make([]int, 0)
	numSlice = append(numSlice, 0)
	for _, n := range total5Str {
		numSlice = append(numSlice, util.ParseInt(string(n)))
	}
	for i := len(numSlice) - 1; i > 0; i-- {
		if numSlice[i] > 2 {
			numSlice[i] = numSlice[i] - 5
			numSlice[i-1]++
		}
	}
	res = strings.Builder{}
	for _, i := range numSlice {
		if i >= 0 {
			res.WriteString(strconv.Itoa(i))
		}
		if i == -1 {
			res.WriteString("-")
		}
		if i == -2 {
			res.WriteString("=")
		}
	}
	fmt.Println(res.String())
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func get5Num(total int, res *strings.Builder) {
	for {
		i, n := total/5, total%5
		defer func() {
			res.WriteString(strconv.Itoa(n))
		}()
		total = i
		if i == 0 {
			break
		}
	}
}

func parseRune(r int32) int {
	if r == '1' {
		return 1
	}
	if r == '2' {
		return 2
	}
	if r == '=' {
		return -2
	}
	if r == '-' {
		return -1
	}
	if r == '0' {
		return 0
	}
	panic(r)
}
