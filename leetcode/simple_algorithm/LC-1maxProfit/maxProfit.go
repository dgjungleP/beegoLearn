package main

import "fmt"

func maxProfit(prices []int) int {
	result := make([]int, len(prices))
	if len(result) == 0 {
		return 0
	}
	result[0] = 0
	for i := 1; i < len(result); i++ {
		if prices[i]-prices[i-1] > 0 {
			result[i] = result[i-1] + prices[i] - prices[i-1]
		} else {
			result[i] = result[i-1]
		}
	}
	return result[len(result)-1]
}

func main() {
	fmt.Print(maxProfit([]int{7, 1, 3, 5, 6, 7, 2, 6}))
}
