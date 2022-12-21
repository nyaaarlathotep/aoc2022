package util

import (
	"io/ioutil"
	"log"
	"path"
)

func GetInput(day string) string {
	f, err := ioutil.ReadFile(path.Join("day"+day, "input"))
	if err != nil {
		log.Fatalln("read fail", err)
	}
	return string(f)
}

func GetLines(day string) []string {
	f, err := ioutil.ReadFile(path.Join("day"+day, "input"))
	if err != nil {
		log.Fatalln("read fail", err)
	}
	return GetStringSlice(string(f), "\n")
}

func GetLineInt(day string) []int {
	f, err := ioutil.ReadFile(path.Join("day"+day, "input"))
	if err != nil {
		log.Fatalln("read fail", err)
	}
	strLine := GetStringSlice(string(f), "\n")
	nums := make([]int, len(strLine))
	for i := range strLine {
		nums[i] = ParseInt(strLine[i])
	}
	return nums
}
