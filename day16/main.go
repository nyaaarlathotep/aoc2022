package main

import (
	"aoc2022/util"
	"log"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("16")
	lines := util.GetStringSlice(input, "\n")
	compileRegex := regexp.MustCompile("Valve (.*) has flow rate=(.*);.* to valves? (.*)")
	pipeMap := make(map[string]*Pipe)
	for _, line := range lines {
		arguments := compileRegex.FindStringSubmatch(line)
		pp := &Pipe{
			name: arguments[1],
			rate: util.ParseInt64(arguments[2]),
			des:  strings.Split(strings.Replace(arguments[3], " ", "", -1), ","),
			open: false,
		}
		pipeMap[pp.name] = pp
	}

	for i := 0; i < 20; i++ {

	}

	elapsed := time.Now().Sub(start)

	log.Println("该函数执行完成耗时：", elapsed)
}

type Pipe struct {
	name string
	rate int64
	des  []string
	open bool
}
