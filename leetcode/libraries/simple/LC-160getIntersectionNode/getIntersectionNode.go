package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	topA := headA
	topB := headB
	count := 0
	for topA != nil || topB != nil {
		if topA == nil && count < 2 {
			topA = headB
			count++
		} else if topB == nil && count < 2 {
			topB = headA
			count++
		}
		if topA == topB {
			return topA
		}
		topA = topA.Next
		topB = topB.Next
	}
	return nil
}

func main() {
	headA := new(ListNode)
	headA.Val = 3
	headB := new(ListNode)
	headB.Val = 2
	headB.Next = new(ListNode)
	headB.Next.Val = 3
	getIntersectionNode(headA, headB)

}
