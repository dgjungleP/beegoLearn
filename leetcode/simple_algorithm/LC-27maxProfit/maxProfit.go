package main

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt64
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func main() {

}
