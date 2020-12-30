package main

import "fmt"

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	return helper(root, sum, true)
}
func helper(root *TreeNode, sum int, flush bool) int {
	if root == nil {
		return 0
	}

	result := helper(root.Left, sum-root.Val, false) + helper(root.Right, sum-root.Val, false)
	if flush {
		result += helper(root.Left, sum, true) + helper(root.Right, sum, true)
	}
	if root.Val == sum {
		result++
	}
	return result
}

func main() {
	root := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}},
	}
	fmt.Println(pathSum(&root, 8))
}
