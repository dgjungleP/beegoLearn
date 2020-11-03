package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target <= nums[0] {
		return 0
	}
	left, right := 0, len(nums)-1
	for left != right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	return left
}
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
