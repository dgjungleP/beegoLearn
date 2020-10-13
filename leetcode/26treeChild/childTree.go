package main

import "fmt"

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return pure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func pure(A *TreeNode, B *TreeNode) bool {

	if B == nil {
		return true
	} else if A == nil {
		return false
	}

	if A.Val == B.Val {
		return pure(A.Left, B.Left) && pure(A.Right, B.Right)
	}
	return false
}

func main() {
	A := new(TreeNode)
	A.Val = 2
	A.Left = new(TreeNode)
	A.Left.Val = 3
	A.Right = new(TreeNode)
	A.Right.Val = 2
	A.Left.Left = new(TreeNode)
	A.Left.Left.Val = 1

	B := new(TreeNode)
	B.Val = 1
	B.Right = new(TreeNode)
	B.Right.Val = 2
	B.Right.Left = new(TreeNode)
	B.Right.Left.Val = 2
	fmt.Print(isSubStructure(A, B))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
