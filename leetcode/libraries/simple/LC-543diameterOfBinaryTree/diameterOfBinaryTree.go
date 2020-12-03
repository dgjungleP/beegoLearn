package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDepth = 1
	if root == nil {
		return 0
	}
	helper(root)
	return maxDepth - 1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := helper(root.Left), helper(root.Right)
	maxDepth = max(maxDepth, left+right+1)
	return max(left, right)

}
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
func main() {

}
