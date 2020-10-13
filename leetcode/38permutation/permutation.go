package main

import "fmt"

func permutation(s string) []string {
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]) + " ")
	}
	return nil
}

func helper(str []rune, index int) {
	if index == len(str)-1 {
		return
	}
}
func main() {

	permutation("hello")
}
