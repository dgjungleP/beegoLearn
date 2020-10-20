package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	value := make([]int, 0)
	max := int(^uint(0) >> 1)
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > max {
			return false
		}
		for len(value) != 0 && value[len(value)-1] > postorder[i] {
			max = value[len(value)-1]
			value = append([]int{}, value[:len(value)-1]...)
		}
		value = append(value, postorder[i])
	}
	return true
}

func main() {
	fmt.Print(verifyPostorder([]int{4, 6, 12, 8, 16, 14, 10}))
}
