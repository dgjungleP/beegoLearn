package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	for left != right {
		result = max(result, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(fir, sec int) int {
	if fir > sec {
		return sec
	}
	return fir
}
func max(fir, sec int) int {
	if fir > sec {
		return fir
	}
	return sec
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
