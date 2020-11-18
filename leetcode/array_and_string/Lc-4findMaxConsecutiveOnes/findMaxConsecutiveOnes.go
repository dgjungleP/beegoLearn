package main

func findMaxConsecutiveOnes(nums []int) int {
	low := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if i-low > max {
				max = i - low
			}
			low = i + 1
		}
		if i == len(nums)-1 {
			if i-low+1 > max {
				max = i - low + 1
			}
		}
	}
	return max
}

func main() {
	findMaxConsecutiveOnes([]int{1})
}
