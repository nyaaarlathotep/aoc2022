package main

import (
	"aoc2022/DataStruct"
	"aoc2022/c"
	"aoc2022/util"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"log"
	"strings"
	"time"
)

const size = 10
const limit = 500

func main() {
	start := time.Now()
	input := util.GetInput("18")
	lines := util.GetStringSlice(input, "\n")

	//partOne(lines)
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
			if _, ok := m[maybeNeighbour]; ok && m[maybeNeighbour] > 0 {
				m[maybeNeighbour]--
				neighbourNum++
			} else if !ok {
				m[maybeNeighbour] = -1
			} else {
				m[maybeNeighbour]--
			}
		}
		m[p] = int64(6 - neighbourNum)
	}
	airPockets := make([][]DataStruct.Point3d, 0)
	edges := c.Select(maps.Keys(m), func(p DataStruct.Point3d) bool { return m[p] < 0 })
	for _, edge := range edges {
		if contain(airPockets, edge) {
			continue
		}
		//ppp := DataStruct.Point3d{X: 2, Y: 2, Z: 5}
		//if edge == ppp {
		//	log.Printf("%v", "111")
		//}
		newAirPocket := make(map[DataStruct.Point3d]int)
		newAirPocket[edge] = 1
		getAirPocket(edge, m, &newAirPocket)

		if len(newAirPocket) > limit {
			continue
		}
		airPockets = append(airPockets, maps.Keys(newAirPocket))
	}
	meaningP := c.Select(maps.Values(m), func(num int64) bool { return num > 0 })
	total2 := getTotal(meaningP)
	totalAir := int64(0)
	for _, airPocket := range airPockets {
		airThisTime := int64(0)
		for _, air := range airPocket {
			airThisTime += m[air]
		}
		totalAir += airThisTime
	}
	log.Printf("%v", totalAir)
	log.Printf("%v", total2+totalAir)

	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func getTotal(meaningP []int64) int64 {
	total2 := int64(0)
	for _, p := range meaningP {
		total2 += p
	}
	return total2
}

func getAirPocket(edge DataStruct.Point3d, m map[DataStruct.Point3d]int64, newAirPocket *map[DataStruct.Point3d]int) {
	for _, neighbour := range getNeighbours(edge) {
		if _, ok := (*newAirPocket)[neighbour]; !ok {
			if m[neighbour] > 0 {
				continue
			}
			(*newAirPocket)[neighbour] = 0
			if len(*newAirPocket) > limit {
				return
			}
			getAirPocket(neighbour, m, newAirPocket)
		}
	}
}

func contain(airPockets [][]DataStruct.Point3d, edge DataStruct.Point3d) bool {
	for _, airPocket := range airPockets {
		if slices.Contains(airPocket, edge) {
			return true
		}
	}
	return false
}

func partOne(lines []string) {
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
