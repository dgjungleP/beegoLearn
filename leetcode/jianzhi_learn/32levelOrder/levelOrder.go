package main

import "fmt"

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	helper := []*TreeNode{root}
	for len(helper) > 0 {
		size := len(helper)
		for i := 0; i < size; i++ {
			node := helper[i]
			if node.Left != nil {
				helper = append(helper, node.Left)
			}
			if node.Right != nil {
				helper = append(helper, node.Right)
			}
			result = append(result, node.Val)
		}
		helper = helper[size:]
	}
	return result
}
func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 5
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7
	fmt.Print(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
