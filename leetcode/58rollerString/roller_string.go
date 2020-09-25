package main

import "fmt"

func reverseLeftWords(s string, n int) string {

	return s[n:] + s[:n]
}

func main() {
	fmt.Printf("%s\n", reverseLeftWords("1231asdad", 2))
}
