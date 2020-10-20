package main

import "fmt"

func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	for i != j {
		m := (i + j) / 2
		if numbers[j] > numbers[m] {
			j = m
		} else if numbers[j] < numbers[m] {
			i = m + 1
		} else {
			j--
		}
	}

	return numbers[i]
}
func main() {
	fmt.Print(minArray([]int{3, 4, 5, 6, 1, 2}))
}
