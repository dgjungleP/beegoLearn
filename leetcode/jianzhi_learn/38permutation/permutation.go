package main

import "fmt"

var str = make([]rune, 0)
var result = make([]string, 0)

func permutation(s string) []string {
	str = []rune(s)
	helper(0)
	return result
}

func helper(index int) {
	if index == len(str)-1 {
		result = append(result, string(str[:]))
		return
	}
	var chartMap = make(map[rune]int)
	for i := index; i < len(str); i++ {
		if _, ok := chartMap[str[i]]; ok {
			continue
		}
		chartMap[str[i]] = 1
		str[i], str[index] = str[index], str[i]
		helper(index + 1)
		str[i], str[index] = str[index], str[i]
	}
}

func main() {
	fmt.Print(permutation("abc"))
}
