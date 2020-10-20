package main

import "fmt"

func replaceSpace(s string) string {
	replace := make([]rune, 0, len(s))

	for _, c := range s {
		if c == ' ' {
			replace = append(replace, '%', '2', '0')
		} else {
			replace = append(replace, c)
		}
	}

	return string(replace)
}

func main() {
	s := "We are happy."
	s = replaceSpace(s)
	fmt.Println(s)
}
