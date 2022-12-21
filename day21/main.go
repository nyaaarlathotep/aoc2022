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
		if parts[0] == "humn" {
			monkeyMap[parts[0]] = &monkey{
				name:     parts[0],
				hasValue: true,
				value: valueWithHuman{
					hasHuman:  true,
					functions: make([]func(value int) int, 0),
				},
			}
			continue
		}
		monkeyMap[parts[0]] = &monkey{
			name:     parts[0],
			hasValue: true,
			value: valueWithHuman{
				hasHuman:  false,
				functions: nil,
				value:     v,
			},
			monkeyName: [2]string{},
			operation:  nil,
		}
	}
	//for _, m := range monkeyMap {
	//	fmt.Printf("%+v\n", *m)
	//}

	valueWithHumanA := evaluate(monkeyMap["root"].monkeyName[0], monkeyMap)
	valueWithHumanB := evaluate(monkeyMap["root"].monkeyName[1], monkeyMap)
	if valueWithHumanA.hasHuman {
		v := valueWithHumanB.value
		for i := len(valueWithHumanA.functions) - 1; i >= 0; i-- {
			v = valueWithHumanA.functions[i](v)
		}
		fmt.Printf("%+v\n", v)
	}
	if valueWithHumanB.hasHuman {
		v := valueWithHumanA.value
		for i := len(valueWithHumanB.functions) - 1; i >= 0; i-- {
			v = valueWithHumanB.functions[i](v)
		}
		fmt.Printf("%+v\n", v)
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func evaluate(monkeyName string, monkeyMap map[string]*monkey) valueWithHuman {
	thisMonkey := monkeyMap[monkeyName]
	if thisMonkey.hasValue {
		return thisMonkey.value
	}
	left := evaluate(thisMonkey.monkeyName[0], monkeyMap)
	right := evaluate(thisMonkey.monkeyName[1], monkeyMap)
	res := thisMonkey.operation(left, right)
	return res
}

func parseMonkey(parts []string, m *monkey) *monkey {
	if strings.Contains(parts[1], "+") {
		waitedMonkey := strings.Split(parts[1], " + ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      valueWithHuman{},
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  addHuman,
		}
		return m
	}
	if strings.Contains(parts[1], "-") {
		waitedMonkey := strings.Split(parts[1], " - ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      valueWithHuman{},
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  minusHuman,
		}
		return m
	}
	if strings.Contains(parts[1], "*") {
		waitedMonkey := strings.Split(parts[1], " * ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      valueWithHuman{},
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  mulHuman,
		}
		return m
	}
	if strings.Contains(parts[1], "/") {
		waitedMonkey := strings.Split(parts[1], " / ")
		m = &monkey{
			name:       parts[0],
			hasValue:   false,
			value:      valueWithHuman{},
			monkeyName: [2]string{waitedMonkey[0], waitedMonkey[1]},
			operation:  divHuman,
		}
		return m
	}
	panic(parts)
}

type monkey struct {
	name       string
	hasValue   bool
	value      valueWithHuman
	monkeyName [2]string
	operation  func(a, b valueWithHuman) valueWithHuman
}
type valueWithHuman struct {
	hasHuman  bool
	functions []func(value int) int
	value     int
}

func addHuman(a, b valueWithHuman) valueWithHuman {
	if a.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(a.functions, func(value int) int {
				return value - b.value
			}),
		}
	}
	if b.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(b.functions, func(value int) int {
				return value - a.value
			}),
		}
	}
	return valueWithHuman{
		hasHuman: false,
		value:    add(a.value, b.value),
	}
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func minusHuman(a, b valueWithHuman) valueWithHuman {
	if a.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(a.functions, func(value int) int {
				return value + b.value
			}),
		}
	}
	if b.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(b.functions, func(value int) int {
				return value - a.value
			}),
		}
	}
	return valueWithHuman{
		hasHuman: false,
		value:    minus(a.value, b.value),
	}
}
func mul(a, b int) int {
	return a * b
}

func mulHuman(a, b valueWithHuman) valueWithHuman {
	if a.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(a.functions, func(value int) int {
				return value / b.value
			}),
		}
	}
	if b.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(b.functions, func(value int) int {
				return value / a.value
			}),
		}
	}
	return valueWithHuman{
		hasHuman: false,
		value:    mul(a.value, b.value),
	}
}

func div(a, b int) int {
	return a / b
}

func divHuman(a, b valueWithHuman) valueWithHuman {
	if a.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(a.functions, func(value int) int {
				return value * b.value
			}),
		}
	}
	if b.hasHuman {
		return valueWithHuman{
			hasHuman: true,
			functions: append(b.functions, func(value int) int {
				panic("1/value")
			}),
		}
	}
	return valueWithHuman{
		hasHuman: false,
		value:    mul(a.value, b.value),
	}
}
