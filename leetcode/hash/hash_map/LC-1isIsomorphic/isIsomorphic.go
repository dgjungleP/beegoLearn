package main

import "strconv"

func isIsomorphic(s string, t string) bool {
	return helper(s) == helper(t)
}

func helper(s string) string {
	indexMap := make(map[rune]int)
	result := ""
	for index, value := range s {
		if v, ok := indexMap[value]; ok {
			result += strconv.Itoa(v)
		} else {
			result += strconv.Itoa(index)
			indexMap[value] = index
		}
	}
	return result
}

func main() {
	isIsomorphic("abb", "bcc")
}
