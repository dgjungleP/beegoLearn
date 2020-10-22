package main

import "fmt"

func rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	nums = append(nums[k:], nums[:k]...)
	fmt.Print(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
