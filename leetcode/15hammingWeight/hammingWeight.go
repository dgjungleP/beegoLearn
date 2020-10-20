package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0

	for num != 0 {
		res++
		num &= num - 1
	}
	return res

}
func main() {
	var num uint32 = 9
	fmt.Print(hammingWeight(num))
}
