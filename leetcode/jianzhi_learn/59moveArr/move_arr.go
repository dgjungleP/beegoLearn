package main

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, 0)
	if len(nums) == 0 || k == 0 {
		return result
	}
	for i := 0; i < len(nums)-k+1; i++ {
		result = append(result, maxNum(nums[i:i+k]))
	}
	return result
}
func maxNum(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{1, -1}
	maxSlidingWindow(nums, 1)

}
