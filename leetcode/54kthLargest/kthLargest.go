package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sortValues = []int{}

func kthLargest(root *TreeNode, k int) int {
	sortValues = make([]int, 0)
	helper(root)
	return sortValues[k-1]
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	sortValues = append(sortValues, root.Val)
	helper(root.Left)
}

func main() {

}
