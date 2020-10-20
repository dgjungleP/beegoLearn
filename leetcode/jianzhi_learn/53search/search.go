package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	right := i
	if j > 0 && nums[j] < target {
		return 0
	}
	i = 0
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	left := j
	return right - left - 1
}

func main() {
	nums := []int{1}
	target := 1
	fmt.Print(search(nums, target))
}
