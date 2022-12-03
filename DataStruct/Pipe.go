package DataStruct

import (
	"aoc2022/util"
	"strconv"
)

type Pipe struct {
	cacheIndex int
	cache      string
}

func NewPipe(input string) *Pipe {
	p := &Pipe{
		cacheIndex: 0,
		cache:      util.SixToTwo(input),
	}
	return p
}

func CacheNewPipe(cache string) *Pipe {
	p := &Pipe{
		cacheIndex: 0,
		cache:      cache,
	}
	return p
}

func (p *Pipe) NextChar() string {
	res := string(p.cache[p.cacheIndex])
	p.cacheIndex++
	return res
}

func getFourBits(str string) string {
	ten, _ := strconv.ParseInt(str, 16, 32)
	bits := strconv.FormatInt(ten, 2)

	zeroNum := 4 - len(bits)
	for i := 0; i < zeroNum; i++ {
		bits = "0" + bits
	}

	return bits
}

func (p *Pipe) GetChars(length int) string {
	res := ""
	for i := 0; i < length; i++ {
		res = res + p.NextChar()
	}
	return res
}

func (p *Pipe) IsEnd() bool {
	if p.cacheIndex > len(p.cache) {
		return true
	} else {
		for i := p.cacheIndex; i < len(p.cache); i++ {
			if string(p.cache[i]) == "1" {
				return false
			}
		}
		return true
	}
}
