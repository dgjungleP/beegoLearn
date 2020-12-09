package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int
var in int

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre = 0
	in = 0
	return helper(preorder, inorder, int(^uint(0)>>1))
}

func helper(preorder, inorder []int, stop int) *TreeNode {
	if pre == len(preorder) {
		return nil
	}
	if inorder[in] == stop {
		in++
		return nil
	}
	rootValue := preorder[pre]
	pre++
	root := &TreeNode{Val: rootValue}
	root.Left = helper(preorder, inorder, rootValue)
	root.Right = helper(preorder, inorder, stop)
	return root
}

func main() {
	buildTree([]int{-1}, []int{-1})
}
