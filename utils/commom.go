package utils

import (
	"regexp"
)

// Str2Uint string表示的数字大于0
func Str2Uint(str string) int {
	res := 0
	for _, v := range str {
		if '0' <= v && v <= '9' {
			res = res*10 + int(v) - '0'
		} else {
			panic("input string is not just contains pure nums")

		}
	}
	return res
}

func RegFind(pattern, targetString string, index ...int) []string {
	res := make([]string, 0)
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(targetString)
	if len(match) == 0 {
		return res

	}
	//fmt.Printf("%#v\n",match)
	//fmt.Println(len(match))

	for _, v := range index {
		res = append(res, match[v])
	}
	return res
}
