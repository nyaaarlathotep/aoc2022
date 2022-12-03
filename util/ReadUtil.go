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
