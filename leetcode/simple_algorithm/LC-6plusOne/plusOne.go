package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			digits[i-1]++
		}
	}
	if digits[0] > 9 {
		digits[0] %= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Print(plusOne([]int{9}))
}
