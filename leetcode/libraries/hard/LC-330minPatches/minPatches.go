package main

func minPatches(nums []int, n int) int {
	i, x, result := 0, 1, 0
	for x <= n {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			result++
		}
	}
	return result
}

func main() {}
