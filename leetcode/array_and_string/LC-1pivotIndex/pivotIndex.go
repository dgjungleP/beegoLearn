package main

func pivotIndex(nums []int) int {
	sum := helper(nums)
	left := 0
	for i := 0; i < len(nums); i++ {
		if left == sum-left-nums[i] {
			return i
		}
		left += nums[i]
	}
	return -1
}

func helper(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {

}
