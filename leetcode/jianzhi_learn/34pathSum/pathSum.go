package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = make([][]int, 0)
var path = make([]int, 0)

func pathSum(root *TreeNode, sum int) [][]int {
	helper(root, sum)
	temp := append([][]int{}, result...)
	result = make([][]int, 0)
	path = make([]int, 0)
	return temp
}

func helper(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		temp := append([]int{}, path[:]...)
		result = append(result, temp)
	}
	helper(root.Left, sum)
	helper(root.Right, sum)
	path = path[:len(path)-1]
}

func main() {
	root := new(TreeNode)
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 4
	root.Right = new(TreeNode)
	root.Right.Val = 8
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 11
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 7
	root.Left.Left.Right = new(TreeNode)
	root.Left.Left.Right.Val = 2
	// root.Left.Right = new(TreeNode)
	// root.Left.Right.Val = 5
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 13
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 4
	root.Right.Right.Left = new(TreeNode)
	root.Right.Right.Left.Val = 5
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 1
	fmt.Print(pathSum(root, 22))
}
