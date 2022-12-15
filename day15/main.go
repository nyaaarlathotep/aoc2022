package main

import (
	"aoc2022/DataStruct"
	"aoc2022/c"
	"aoc2022/util"
	"log"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("15")
	lines := util.GetStringSlice(input, "\n")
	compileRegex := regexp.MustCompile("Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)")
	sensors := make([]*sensor, 0)
	for _, line := range lines {
		arguments := compileRegex.FindStringSubmatch(line)
		s := &sensor{
			pos: DataStruct.Point{
				X: util.ParseInt64(arguments[1]),
				Y: util.ParseInt64(arguments[2]),
			},
			beacon: DataStruct.Point{
				X: util.ParseInt64(arguments[3]),
				Y: util.ParseInt64(arguments[4]),
			},
			dis: 0,
		}
		s.dis = c.Abs(s.pos.X-s.beacon.X) + c.Abs(s.pos.Y-s.beacon.Y)
		sensors = append(sensors, s)
	}
	l := DataStruct.Line{
		Start: 0,
		End:   4000000,
	}
	//partOne(sensors, 2000000)
	for y := 0; y < 4000000; y++ {
		lines := partOne(sensors, int64(y))
		for _, line := range lines {
			if !line.Contain(&l) {
				log.Printf("(%v,%v)", line.End+1, y)
				log.Printf("%v", (line.End+1)*4000000+int64(y))
				break
			}
		}
	}

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func partOne(sensors []*sensor, y int64) []DataStruct.Line {

	noBeaconLines := make([]DataStruct.Line, 0)
	for _, s := range sensors {
		disY := c.Abs(s.pos.Y - y)
		if disY > s.dis {
			continue
		}
		line := DataStruct.NewLine(s.pos.X-(s.dis-disY), s.pos.X+(s.dis-disY))
		noBeaconLines = append(noBeaconLines, *line)
	}
	DataStruct.SortLines(noBeaconLines)
	mergedLines := make([]DataStruct.Line, 0)
	mergedLines = append(mergedLines, noBeaconLines[0])
	for i := 1; i < len(noBeaconLines); i++ {
		if mergedLines[len(mergedLines)-1].Involve(&noBeaconLines[i]) {
			mergedLines[len(mergedLines)-1] = mergedLines[len(mergedLines)-1].Add(noBeaconLines[i])
			continue
		}
		mergedLines = append(mergedLines, noBeaconLines[i])
	}
	overlapped := make(map[DataStruct.Point]bool)
	for _, s := range sensors {
		if s.beacon.Y == y {
			overlapped[s.beacon] = true
		}
		if s.pos.Y == y {
			overlapped[s.pos] = true
		}
	}
	//log.Printf("mergedLines: %+v", mergedLines)
	//log.Printf("overlapped: %+v", len(overlapped))
	var total int64 = 0
	for _, l := range mergedLines {
		total += l.End - l.Start + 1
	}
	//log.Printf("%+v", total-int64(len(overlapped)))
	return mergedLines
}

type sensor struct {
	pos    DataStruct.Point
	beacon DataStruct.Point
	dis    int64
}

// 6,853,791
// 4,737,443
