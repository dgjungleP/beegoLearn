package main

import (
	"bytes"
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	var buffer bytes.Buffer
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			buffer.WriteString(arr[i] + " ")
		}
	}
	return strings.Trim(buffer.String(), " ")
}

func main() {
	fmt.Print(reverseWords("the sky is blue") == "blue is sky the")
}
