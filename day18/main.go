package main

import (
	"aoc2022/DataStruct"
	"aoc2022/util"
	"log"
	"strings"
	"time"
)

const size = 10

func main() {
	start := time.Now()
	input := util.GetInput("18")
	lines := util.GetStringSlice(input, "\n")

	m := make(map[DataStruct.Point3d]int64)
	for _, line := range lines {
		pointAru := strings.Split(line, ",")
		p := DataStruct.Point3d{
			X: util.ParseInt64(pointAru[0]),
			Y: util.ParseInt64(pointAru[1]),
			Z: util.ParseInt64(pointAru[2]),
		}
		neighbours := getNeighbours(p)
		neighbourNum := 0
		for _, maybeNeighbour := range neighbours {
			if _, ok := m[maybeNeighbour]; ok {
				m[maybeNeighbour]--
				neighbourNum++
			}
		}
		m[p] = int64(6 - neighbourNum)
	}

	total := int64(0)
	for _, i := range m {
		total += i
	}

	log.Printf("%v", total)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func getNeighbours(p DataStruct.Point3d) []DataStruct.Point3d {
	neighbours := make([]DataStruct.Point3d, 0)
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: 1,
		Y: 0,
		Z: 0,
	}))
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: -1,
		Y: 0,
		Z: 0,
	}))
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: 0,
		Y: 1,
		Z: 0,
	}))
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: 0,
		Y: -1,
		Z: 0,
	}))
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: 0,
		Y: 0,
		Z: 1,
	}))
	neighbours = append(neighbours, p.Add(DataStruct.Point3d{
		X: 0,
		Y: 0,
		Z: -1,
	}))
	return neighbours
}
