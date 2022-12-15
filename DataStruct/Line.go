package DataStruct

import "sort"

type Line struct {
	Start int64
	End   int64
}

func NewLine(start int64, end int64) *Line {
	return &Line{
		Start: start,
		End:   end,
	}
}

func (l *Line) Involve(l1 *Line) bool {
	return (l1.Start >= l.Start-1 && l1.Start <= l.End+1) ||
		(l1.End >= l.Start-1 && l1.End <= l.End+1) ||
		(l1.Start <= l.Start && l1.End >= l.End) ||
		(l1.Start >= l.Start && l1.End <= l.End)
}

func (l *Line) Contain(l1 *Line) bool {
	return l1.Start >= l.Start && l1.End <= l.End
}

func (l *Line) Add(l1 Line) Line {
	if l1.Start > l.Start {
		l1.Start = l.Start
	}
	if l1.End < l.End {
		l1.End = l.End
	}
	return l1
}

func SortLines(lines []Line) {
	sort.Sort(L(lines))
}

type L []Line

func (p L) Len() int {
	return len(p)
}
func (p L) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}
func (p L) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
