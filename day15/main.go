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
	var y int64 = 2000000

	noBeaconLines := make([]DataStruct.Line, 0)
	for _, s := range sensors {
		disY := c.Abs(s.pos.Y - y)
		if disY > s.dis {
			continue
		}
		line := DataStruct.NewLine(s.pos.X-(s.dis-disY), s.pos.X+(s.dis-disY))
		//log.Printf("%+v", line)
		noBeaconLines = append(noBeaconLines, *line)
	}
	DataStruct.SortLines(noBeaconLines)
	log.Printf("%+v", noBeaconLines)
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
			log.Printf("overlapped beacon: (%+v)", s.beacon)
			overlapped[s.beacon] = true
		}
		if s.pos.Y == y {
			log.Printf("overlapped beacon: (%+v)", s.beacon)
			overlapped[s.pos] = true
		}
	}
	log.Printf("mergedLines: %+v", mergedLines)
	log.Printf("overlapped: %+v", len(overlapped))
	var total int64 = 0
	for _, l := range mergedLines {
		total += l.End - l.Start + 1
	}

	log.Printf("%+v", total-int64(len(overlapped)))

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

type sensor struct {
	pos    DataStruct.Point
	beacon DataStruct.Point
	dis    int64
}

// 6,853,791
// 4,737,443
