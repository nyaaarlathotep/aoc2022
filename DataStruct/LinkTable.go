package DataStruct

import (
	"aoc2022/util"
)

type LinkTable struct {
	Name string
	big  bool
	Ptr  []*LinkTable
}

func NewLinkTable(name string) *LinkTable {
	t := &LinkTable{}
	t.Name = name
	t.Ptr = make([]*LinkTable, 0)
	if util.IsUpper(name) {
		t.big = true
	} else {
		t.big = false
	}

	return t
}

func (linkTable *LinkTable) IsBig() bool {
	return linkTable.big
}

func (linkTable *LinkTable) AddWay(dest *LinkTable) {
	linkTable.Ptr = append(linkTable.Ptr, dest)
}
