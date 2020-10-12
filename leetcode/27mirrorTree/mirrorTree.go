package main

func mirrorTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	temp := mirrorTree(root.Left)
	root.Left = mirrorTree(root.Right)
	root.Right = temp
	return root
}

func main() {
	root := new(TreeNode)
	root.Val = 4
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 7
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 9
	mirrorTree(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
