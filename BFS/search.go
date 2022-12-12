package BFS

import (
	// "log"

	"aoc2022/DataStruct"
	"aoc2022/routing"
)

func MineBFS(g routing.Graph, start DataStruct.Point, goal DataStruct.Point) []DataStruct.Point {
	var frontier DataStruct.Queue[DataStruct.Point]
	frontier.Put(start)

	cameFrom := make(map[DataStruct.Point]*DataStruct.Point)
	cameFrom[start] = nil

	for !frontier.Empty() {
		current := frontier.Get()
		if current == goal {
			break
		}

		//log.Println(g.Neighbours(current))

		for _, n := range g.Neighbours(current) {
			if _, ok := cameFrom[n]; !ok {
				frontier.Put(n)
				cameFrom[n] = &current
			}
		}
	}

	ret := []DataStruct.Point{goal}
	for n := cameFrom[goal]; n != nil; n = cameFrom[*n] {
		ret = append(ret, *n)
	}

	return ret
}
