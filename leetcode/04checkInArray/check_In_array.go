package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	length := len(matrix[0]) - 1
	high := len(matrix) - 1
	//第一步过滤范围
	for i, v := range matrix[0] {
		if v >= target {
			length = i
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] >= target {
			high = i
			break
		}
	}
	//缩小范围之后开始甄别
	for i := 0; i <= high; i++ {
		for j := 0; j <= length; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func main() {
	var findIt bool
	// matrix := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30}}
	matrix := [][]int{{-5}}
	target := -5
	findIt = findNumberIn2DArray(matrix, target)
	fmt.Printf("%d:is  %t\n", target, findIt)
	target = 20
	findIt = findNumberIn2DArray(matrix, target)
	fmt.Printf("%d:is  %t", target, findIt)
}
