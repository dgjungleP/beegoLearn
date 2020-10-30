package main

func climbStairs(n int) int {
	sum := 1
	left := 0
	right := 0
	for i := 0; i < n; i++ {
		left = right
		right = sum
		sum = left + right
	}
	return sum
}

func main() {

}
