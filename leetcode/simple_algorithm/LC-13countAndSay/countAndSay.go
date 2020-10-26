package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	str := "1"
	if n == 1 {
		return str
	}
	if n == 2 {
		return "11"
	}
	n--
	str = countAndSay(n)
	str = countStr(str)
	return str
}
func countStr(str string) string {
	buffer := strings.Builder{}
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			count++
		} else {
			buffer.WriteString(strconv.Itoa(count))
			buffer.WriteString(string(str[i-1]))
			count = 1
		}
		if i == len(str)-1 {
			buffer.WriteString(strconv.Itoa(count))
			buffer.WriteString(string(str[i]))
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
	fmt.Println(countAndSay(9))
}
