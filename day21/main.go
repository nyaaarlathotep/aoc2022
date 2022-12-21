package main

import (
	"aoc2022/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("21")
	lines := util.GetStringSlice(input, "\n")
	monkeyMap := make(map[string]*monkey)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			var m *monkey
			m = parseMonkey(parts, m)
			monkeyMap[m.name] = m
			continue
		}
		monkeyMap[parts[0]] = &monkey{
			name:       parts[0],
			hasValue:   true,
			value:      v,
			monkeyName: [2]string{},
			operation:  nil,
		}
	}
	//for _, m := range monkeyMap {
	//	fmt.Printf("%+v\n", *m)
	//}

	fmt.Printf("%+v\n", evaluate("root", monkeyMap))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func evaluate(monkeyName string, monkeyMap map[string]*monkey) int {
	thisMonkey := monkeyMap[monkeyName]
	if thisMonkey.hasValue {
		return thisMonkey.value
	}
	return thisMonkey.operation(evaluate(thisMonkey.monkeyName[0], monkeyMap), evaluate(thisMonkey.monkeyName[1], monkeyMap))
}

func parseMonkey(parts []string, m *monkey) *monkey {
	if strings.Contains(parts[1], "+") {
		waitedMonkey := strings.Split(parts[1], " + ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      0,
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  add,
		}
		return m
	}
	if strings.Contains(parts[1], "-") {
		waitedMonkey := strings.Split(parts[1], " - ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      0,
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  minus,
		}
		return m
	}
	if strings.Contains(parts[1], "*") {
		waitedMonkey := strings.Split(parts[1], " * ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      0,
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  mul,
		}
		return m
	}
	if strings.Contains(parts[1], "/") {
		waitedMonkey := strings.Split(parts[1], " / ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      0,
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  div,
		}
		return m
	}
	panic(parts)
}

type monkey struct {
	name       string
	hasValue   bool
	value      int
	monkeyName [2]string
	operation  func(a, b int) int
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
