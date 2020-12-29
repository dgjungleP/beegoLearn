package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sold, ice, hold := 0, 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		newHold := max(hold, sold-prices[i])
		newIce := hold + prices[i]
		newSold := max(sold, ice)
		sold, ice, hold = newSold, newIce, newHold
	}
	return max(sold, ice)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit([]int{1, 2, 3, 0, 2})
}
