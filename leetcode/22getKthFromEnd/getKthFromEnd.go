package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre := head
	next := pre.Next
	for i := 0; i < k-1; i++ {
		if next == nil {
			return nil
		}
		next = next.Next
	}
	for next != nil {
		pre = pre.Next
		next = next.Next
	}
	return pre
}
func main() {

}
