package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := new(strings.Builder)
	for i := 0; i < len(strs[0]); i++ {
		if helper(strs, i) {
			prefix.WriteString(string(strs[0][i]))
		} else {
			return prefix.String()
		}
	}
	return prefix.String()
}
func helper(strs []string, i int) bool {
	for j := 1; j < len(strs); j++ {
		if len(strs[j]) <= i || strs[j-1][i] != strs[j][i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
