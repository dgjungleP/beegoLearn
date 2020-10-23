package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for right > left {
		if !isalnum(s[left]) {
			left++
			continue
		}
		if !isalnum(s[right]) {
			right--
			continue
		}
		if right > left {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isalnum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
