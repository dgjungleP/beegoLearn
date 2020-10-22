package main

import "fmt"

func moveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}
	fmt.Print(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
