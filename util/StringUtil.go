package util

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func GetDecimalFromBinaryInt(binarySlice []int) int {
	var buffer bytes.Buffer
	for _, v := range binarySlice {
		buffer.WriteString(strconv.Itoa(v))
	}
	res, _ := strconv.ParseInt(buffer.String(), 2, 64)
	return int(res)
}

func GetDecimalFromBinaryString(binarySlice []string) int {
	binary := strings.Join(binarySlice, "")
	res, _ := strconv.ParseInt(binary, 2, 64)
	return int(res)
}

func GetSliceFromString(str string) []string {
	return strings.Split(str, "")
}

func StringAt(str string, index int) string {
	tmp := strings.Split(str, "")
	return tmp[index]
}

func Get2dString(input string, sep1 string, sep2 string) [][]string {
	input = strings.Replace(input, "  ", " ", -1)
	lines := strings.Split(input, sep1)
	tDString := make([][]string, 0)
	for _, v := range lines {
		v = strings.Trim(v, " ")
		line := strings.Split(v, sep2)
		tDString = append(tDString, line)
	}
	return tDString
}

func GetStringSlice(input string, sep string) []string {
	return strings.Split(input, sep)
}

func TwoDStringToInt(s [][]string) [][]int {

	res := make([][]int, 0)
	for i := range s {
		res = append(res, make([]int, 0))
		for j := range s[0] {
			v, _ := strconv.Atoi(s[i][j])
			res[i] = append(res[i], v)
		}
	}
	return res
}

// IsUpper 判断字符 r 是否为大写格式
func IsUpper(s string) bool {
	for _, r := range s {
		// 判断字符是否为大写
		if unicode.IsUpper(r) {
			return true
		} else {
			return false
		}
	}
	return false
}

func CopyStringSlice(in []string) []string {
	re := make([]string, 0)
	for _, s := range in {
		re = append(re, s)
	}
	return re
}

func RemoveByLoop(slc []string) []string {
	var res []string
	for i := range slc {
		flag := true
		for j := range res {
			if slc[i] == res[j] {
				flag = false
				break
			}
			if flag {
				res = append(res, slc[i])
			}
		}

	}
	return res
}

func Print2DSlice[T int | string | float64](TDString [][]T, join string) {
	for _, v := range TDString {
		line := ""
		for _, s := range v {
			line = line + fmt.Sprintf("%v%v", s, join)
		}
		log.Println(line)
	}
}
func Print2DRune(TDString [][]rune, join string) {
	for _, v := range TDString {
		line := ""
		for _, s := range v {
			line = line + fmt.Sprintf("%v%v", string(s), join)
		}
		log.Println(line)
	}
}
func ParseInt(num string) int {
	num = strings.Trim(num, " ")
	v, err := strconv.Atoi(num)
	Check(err)
	return v
}
func Parse2DString(num []int) []string {
	str := make([]string, len(num))
	for i, v := range num {
		str[i] = strconv.Itoa(v)
	}
	return str
}
func ParseInt64(num string) int64 {
	v, err := strconv.Atoi(num)
	Check(err)
	return int64(v)
}

func SixToTwo(input string) string {
	convMap := make(map[string]string)
	convMap["0"] = "0000"
	convMap["1"] = "0001"
	convMap["2"] = "0010"
	convMap["3"] = "0011"
	convMap["4"] = "0100"
	convMap["5"] = "0101"
	convMap["6"] = "0110"
	convMap["7"] = "0111"
	convMap["8"] = "1000"
	convMap["9"] = "1001"
	convMap["A"] = "1010"
	convMap["B"] = "1011"
	convMap["C"] = "1100"
	convMap["D"] = "1101"
	convMap["E"] = "1110"
	convMap["F"] = "1111"
	res := ""
	for i := range input {
		res = res + convMap[string(input[i])]
	}
	return res
}
