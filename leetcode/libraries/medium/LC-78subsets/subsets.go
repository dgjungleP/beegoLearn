package main

import "fmt"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = helper(result, nums)
	return append(result, []int{})
}

func helper(result [][]int, nums []int) [][]int {
	if len(nums) == 0 {
		return result
	}
	for _, num := range result {
		result = append(result, append(num, nums[0]))
	}
	result = append(result, []int{nums[0]})
	return helper(result, nums[1:])
}

func main() {
	fmt.Print(subsets([]int{1, 2, 3}))
}
