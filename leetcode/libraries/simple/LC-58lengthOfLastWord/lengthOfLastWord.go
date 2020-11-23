package main

import "fmt"

func lengthOfLastWord(s string) int {
	tail := len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			tail--
		} else {
			break
		}
	}
	for i := tail; i >= 0; i-- {
		if s[i] == ' ' {
			return tail - i
		}
	}
	return tail + 1
}

func main() {
	fmt.Println(lengthOfLastWord("a"))
}
