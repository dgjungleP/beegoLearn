package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode = new(ListNode)
	head := l3
	left := l1
	right := l2

	for left != nil && right != nil {
		if left.Val <= right.Val {
			l3.Next = left
			left = left.Next
		} else {
			l3.Next = right
			right = right.Next
		}
		l3 = l3.Next
	}
	if left != nil {
		l3.Next = left
	} else if right != nil {
		l3.Next = right
	}
	return head.Next
}
func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4
	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 3
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	mergeTwoLists(l1, l2)
}
