package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := head
	next := pre.Next
	for i := 1; i < n; i++ {
		next = next.Next
	}
	if next == nil {
		return head.Next
	}
	for next.Next != nil {
		pre = pre.Next
		next = next.Next
	}
	pre.Next = pre.Next.Next
	return head
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
	removeNthFromEnd(head, 2)
}
