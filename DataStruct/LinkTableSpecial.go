package DataStruct

import (
	"aoc2022/util"
)

type LinkTableSpecial struct {
	Name        string
	big         bool
	enterChance int
	Ptr         []*LinkTableSpecial
}

func NewLinkTableSpecial(name string) *LinkTableSpecial {
	t := &LinkTableSpecial{}
	t.Name = name
	t.Ptr = make([]*LinkTableSpecial, 0)
	if util.IsUpper(name) {
		t.big = true
		t.enterChance = -1
	} else {
		t.big = false
		t.enterChance = 2
	}

	return t
}

func (linkTable *LinkTableSpecial) IsBig() bool {
	if linkTable.big {
		return true
	} else {
		if linkTable.enterChance > 0 {
			linkTable.enterChance--
			return true
		}
		return false

	}
}

func (linkTable *LinkTableSpecial) AddWay(dest *LinkTableSpecial) {
	linkTable.Ptr = append(linkTable.Ptr, dest)
}
