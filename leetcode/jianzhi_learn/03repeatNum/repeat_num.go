package main

import "fmt"

func findRepeatNumber(nums []int) int {
	mapset := make(map[int]bool)
	for _, num := range nums {
		if _, ok := mapset[num]; ok {
			return num
		} else {
			mapset[num] = true
		}
	}

	return -1
}
func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	replayNum := findRepeatNumber(nums)
	fmt.Printf("%d", replayNum)
}
