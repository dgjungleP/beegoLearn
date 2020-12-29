package main

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	redPoint, whitePoint := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[redPoint], nums[i] = nums[i], nums[redPoint]
			if whitePoint > redPoint {
				nums[whitePoint], nums[i] = nums[i], nums[whitePoint]
			}
			redPoint++
			whitePoint++
		} else if nums[i] == 1 {
			nums[whitePoint], nums[i] = nums[i], nums[whitePoint]
			whitePoint++
		}

	}
}

func main() {
	sortColors([]int{2, 0, 1, 1, 2, 0})
}
