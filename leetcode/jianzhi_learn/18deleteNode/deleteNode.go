package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	pre := head
	next := head.Next
	if pre.Val == val {
		return head.Next
	}
	for next != nil {
		if next.Val == val {
			pre.Next = next.Next
			break
		}
		pre, next = next, next.Next

	}
	return head
}
func main() {

}
