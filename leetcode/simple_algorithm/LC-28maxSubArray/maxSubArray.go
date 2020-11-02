package main

func maxSubArray(nums []int) int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if result[i-1] > 0 {
			result[i] = result[i-1] + nums[i]
		} else {
			result[i] = nums[i]
		}
		if result[i] > max {
			max = result[i]
		}
	}
	return max
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
