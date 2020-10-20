package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([]int, len(prices))
	result := 0
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - prices[i-1]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
