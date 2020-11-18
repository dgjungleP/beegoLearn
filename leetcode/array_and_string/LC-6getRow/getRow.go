package main

import "fmt"

func getRow(rowIndex int) []int {
	rowIndex++
	result := make([]int, rowIndex)

	for i := 0; i < rowIndex; i++ {
		result[i] = 1
		for j := i - 1; j > 0 && i > 1; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}

func main() {
	fmt.Println(getRow(16))
}
