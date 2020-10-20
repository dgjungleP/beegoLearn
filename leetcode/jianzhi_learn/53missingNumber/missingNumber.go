package main

import "fmt"

func missingNumber(nums []int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		m := (left + right) >> 1
		if nums[m] == m {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return left
}

func main() {
	nums := []int{1}
	fmt.Println(missingNumber(nums))
}
