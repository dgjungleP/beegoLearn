package main

func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0)
	pre := 0
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[pre] {
			count++
		} else {
			if count >= 3 {
				result = append(result, []int{pre, i - 1})
			}
			pre = i
			count = 1
		}
	}
	if count >= 3 {
		result = append(result, []int{pre, len(s) - 1})
	}
	return result
}

func main() {
	largeGroupPositions("aaa")
}
