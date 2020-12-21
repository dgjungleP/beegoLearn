package main

func minCostClimbingStairs(cost []int) int {
	pre, next := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, next = next, min(next+cost[i-1], pre+cost[i-2])
	}
	return next
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	minCostClimbingStairs([]int{0, 0, 1, 1})
}
