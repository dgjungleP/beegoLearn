package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 10))
}
