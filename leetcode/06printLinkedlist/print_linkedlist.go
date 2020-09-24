package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	result := make([]int, 1, 1)
	result[0] = head.Val
	if head.Next == nil {
		return result
	} else {
		return append(reversePrint(head.Next), result...)
	}
}

func main() {
	list := new(ListNode)
	list.Val = 1
	list.Next = new(ListNode)
	result := reversePrint(nil)
	fmt.Print(result)
}
