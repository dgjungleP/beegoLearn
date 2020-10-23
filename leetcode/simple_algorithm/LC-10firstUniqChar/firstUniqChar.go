package main

import "fmt"

func firstUniqChar(s string) int {
	valueMap := make(map[rune]int)

	for _, c := range s {
		valueMap[c]++
	}
	for index, value := range s {
		if valueMap[value] == 1 {
			return index
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
