package main

import (
	"aoc2022/util"
	"log"
	"path"
	"time"
)

func main() {
	start := time.Now()
	input := util.GetInput("07")
	lines := util.Get2dString(input, "\n", " ")
	dirNow := &dir{
		path:   "/",
		parent: nil,
		isDir:  true,
		sons:   make([]*dir, 0),
	}
	root := dirNow
	readTree(lines, dirNow)
	res := scanPartOne(root, make([]*dir, 0))
	var total int64
	for _, dir := range res {
		total += dir.length
	}
	log.Printf("%+v", total)
	limit := 30000000 - (70000000 - root.length)
	log.Println(limit)
	smallest := scanPartTwo(root, 70000000, limit)
	log.Printf("%+v", smallest)
	elapsed := time.Now().Sub(start)
	log.Println("该函数执行完成耗时：", elapsed)
}

func readTree(lines [][]string, dirNow *dir) {
	for _, line := range lines {
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					dirNow = dirNow.parent
					continue
				}
				dir := &dir{
					path:   path.Join(dirNow.path, line[2]),
					parent: dirNow,
					isDir:  true,
					sons:   make([]*dir, 0),
				}
				dirNow.sons = append(dirNow.sons, dir)
				dirNow = dir
			}
		} else if line[0] != "dir" {
			file := &dir{
				path:   path.Join(dirNow.path, line[1]),
				parent: dirNow,
				isDir:  false,
				length: util.ParseInt64(line[0]),
				sons:   nil,
			}
			dirNow.sons = append(dirNow.sons, file)
			addFileLength(dirNow, file)
		}
	}
}

func scanPartOne(dirNow *dir, res []*dir) []*dir {

	if dirNow.isDir && dirNow.length <= 100000 {
		res = append(res, dirNow)
	}
	for _, son := range dirNow.sons {
		if son.isDir {
			res = scanPartOne(son, res)
		}
	}
	return res
}

func scanPartTwo(dirNow *dir, smallest int64, limit int64) int64 {

	if dirNow.length < limit {
		return 70000000
	}
	if dirNow.isDir && dirNow.length >= limit && dirNow.length < smallest {
		smallest = dirNow.length
	}
	for _, son := range dirNow.sons {
		if son.isDir {
			sonSmallest := scanPartTwo(son, smallest, limit)
			if sonSmallest < smallest {
				smallest = sonSmallest
			}
		}
	}
	return smallest
}
func addFileLength(dirNow *dir, file *dir) {
	//atomic.AddInt64(&(dirNow.length), file.length)
	dirNow.length = dirNow.length + file.length
	if dirNow.parent != nil {
		addFileLength(dirNow.parent, file)
	}
}

type dir struct {
	path   string
	parent *dir
	isDir  bool
	length int64
	sons   []*dir
}
