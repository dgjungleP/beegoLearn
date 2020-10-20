package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for nums[left]+nums[right] != target {
		if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
		if left >= right {
			return nil
		}
	}
	return []int{nums[left], nums[right]}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
