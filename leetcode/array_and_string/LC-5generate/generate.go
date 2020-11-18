package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	tips := make([]int, numRows)
	if numRows == 1 {
		tips[0] = 1
		return [][]int{tips}
	}
	slipe := generate(numRows - 1)
	tips[0] = 1
	tips[numRows-1] = 1
	for i := 1; i < numRows-1; i++ {
		tips[i] = slipe[numRows-2][i-1] + slipe[numRows-2][i]
	}
	slipe = append(slipe, tips)
	return slipe
}

func main() {
	fmt.Print(generate(5))
}
