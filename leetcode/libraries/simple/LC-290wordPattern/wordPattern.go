package main

import "strings"

func wordPattern(pattern string, s string) bool {
	matches := make(map[string]byte)
	partMap := make(map[byte]int)
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i, word := range words {
		part := pattern[i]
		if _, ok := matches[word]; !ok {
			if _, ok := partMap[part]; !ok {
				matches[word] = part
				partMap[part] = 1
			} else {
				return false
			}
		}
		if matches[word] != part {
			return false
		}
	}
	return true
}

func main() {
	wordPattern("abba", "dog cat cat dog")
}
