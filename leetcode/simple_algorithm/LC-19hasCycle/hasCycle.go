package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	low := head
	faster := head
	for faster != nil {
		if faster.Next == nil {
			return false
		}
		faster = faster.Next.Next
		if low == faster {
			return true
		}
		low = low.Next

	}
	return false
}

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 1
}
