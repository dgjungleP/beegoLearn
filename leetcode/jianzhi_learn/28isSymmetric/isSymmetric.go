package main

import "fmt"

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && helper(left.Left, right.Right) && helper(right.Left, left.Right)
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val =2
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 4
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 3
	fmt.Print(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
