package main

import "fmt"

func exchange(nums []int) []int {

	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			for i <= right && right >= 0 {
				if nums[right]%2 != 0 {
					nums[i], nums[right] = nums[right], nums[i]
					right--
					break
				} else {
					right--
				}
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
