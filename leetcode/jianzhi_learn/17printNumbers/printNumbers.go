package main

import (
	"fmt"
	"strconv"
)

var baseArr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func printNumbers(n int) []int {
	result := make([]int, 0)
	for _, strValue := range helper("0", n) {
		intValue, _ := strconv.Atoi(strValue)
		result = append(result, intValue)
	}
	return result[1:]
}
func helper(left string, n int) []string {
	helpeArr := make([]string, 0)
	if n == 0 {
		return []string{left}
	}
	for _, value := range baseArr {
		if left != "0" {
			helpeArr = append(helpeArr, helper(left+value, n-1)...)
		} else {
			helpeArr = append(helpeArr, helper(value, n-1)...)
		}
	}
	return helpeArr[:]
}
func main() {
	fmt.Print(printNumbers(2))
}
