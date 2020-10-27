package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	base := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return base
}

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 4
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 7
	result := reverseList(head)
	fmt.Printf("%+v", result)
}
