package main

import (
	"aoc2022/c"
	"aoc2022/util"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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
		thisPipe := &Pipe{
			name:    arguments[1],
			rate:    util.ParseInt(arguments[2]),
			next:    strings.Split(strings.Replace(arguments[3], " ", "", -1), ","),
			farNext: make([]string, 0),
			dis:     make(map[string]int),
			open:    false,
		}
		if thisPipe.rate == 0 {
			for _, p := range c.Select(maps.Values(pipeMap), func(pipe *Pipe) bool { return slices.Contains(pipe.next, thisPipe.name) }) {
				for _, s := range thisPipe.next {
					p.farNext = append(p.farNext, s)
				}
			}
		}
		pipeMap[thisPipe.name] = thisPipe
	}
	calDis(pipeMap["AA"], pipeMap["AA"], 0, pipeMap)
	for _, s := range pipeMap {
		if s.rate != 0 {
			calDis(s, s, 0, pipeMap)
		}
	}

	a, _ := check(pipeMap["AA"], pipeMap, make([]string, 0), 0, 0)
	log.Printf("%+v", a)

	elapsed := time.Now().Sub(start)

	log.Println("该函数执行完成耗时：", elapsed)
}

func check(pipeNow *Pipe, pipeMap map[string]*Pipe, roadMap []string, minute int, flowNow int) (int, []string) {
	bestFlow := flowNow
	bestRoadMap := roadMap
	for endName, dis := range pipeNow.dis {
		if pipeMap[endName].rate != 0 && !slices.Contains(roadMap, endName) {
			minuteNow := minute + dis + 1
			if minuteNow >= 30 {
				//log.Printf("%+v", roadMap)
				continue
			}
			f := flowNow + (30-minuteNow)*pipeMap[endName].rate
			r, rm := check(pipeMap[endName], pipeMap, append(roadMap, endName), minuteNow, f)
			if r > bestFlow {
				bestFlow = r
				bestRoadMap = rm
			}
		}
	}
	// DD BB JJ HH EE CC
	//log.Printf("%+v", bestRoadMap)
	//log.Printf("%+v", minute)
	return bestFlow, bestRoadMap
}

func calDis(s *Pipe, now *Pipe, distance int, pipeMap map[string]*Pipe) {
	for _, p := range now.next {
		if pipeMap[p].name == s.name {
			continue
		}
		if _, ok := s.dis[p]; !ok {
			s.dis[p] = distance + 1
			calDis(s, pipeMap[p], distance+1, pipeMap)
		} else if distance < s.dis[p] {
			s.dis[p] = distance + 1
			calDis(s, pipeMap[p], distance+1, pipeMap)
		}
	}
	for _, p := range now.farNext {
		if pipeMap[p].name == s.name {
			continue
		}
		if _, ok := s.dis[p]; !ok {
			s.dis[p] = distance + 2
			calDis(s, pipeMap[p], distance+2, pipeMap)
		} else if distance < s.dis[p] {
			s.dis[p] = distance + 2
			calDis(s, pipeMap[p], distance+2, pipeMap)
		}
	}
}

type Pipe struct {
	name    string
	rate    int
	next    []string
	farNext []string
	dis     map[string]int
	open    bool
}
