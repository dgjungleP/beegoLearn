package main

import "fmt"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if result <= dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	fmt.Print(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
