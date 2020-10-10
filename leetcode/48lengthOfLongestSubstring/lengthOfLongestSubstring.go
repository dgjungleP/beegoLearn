package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	temp := 0
	maxLength := 0
	chartMap := make(map[byte]int)
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if index, ok := chartMap[s[i]]; !ok {
			chartMap[s[i]] = i
			if i == 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-1] + 1
			}
		} else {
			chartMap[s[i]] = i
			if index < temp {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = i - index
				temp = index
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfLongestSubstring("blqsearxxxbiwqa"))
}
