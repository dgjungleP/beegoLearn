package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 1
	}

	return 1 + helper(maxDepth(root.Left), maxDepth(root.Right))
}
func helper(fir, sec int) int {

	if fir > sec {
		return fir
	}
	return sec
}
func main() {

}
