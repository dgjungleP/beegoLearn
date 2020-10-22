package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	values := make(map[int]int)
	if len(nums1) >= len(nums2) {
		for _, num := range nums1 {
			values[num]++
		}
		for _, num := range nums2 {
			if value, ok := values[num]; ok {
				if value > 0 {
					result = append(result, num)
					values[num]--
				}
			}
		}
	} else {
		for _, num := range nums2 {
			values[num]++
		}
		for _, num := range nums1 {
			if value, ok := values[num]; ok {
				if value > 0 {
					result = append(result, num)
					values[num]--
				}
			}
		}
	}
	return result
}
func main() {
	fmt.Print(intersect([]int{1, 2, 3, 3, 4, 5, 5, 6, 6, 7}, []int{4, 3, 3, 2, 1, 5, 7}))
}
