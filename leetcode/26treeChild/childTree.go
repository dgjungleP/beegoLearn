package main

import "fmt"

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return pure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func pure(A *TreeNode, B *TreeNode) bool {

	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		return false
	}

	if A.Val == B.Val {

		return pure(A.Left, B.Left) && pure(A.Right, B.Right)
	}
	return false
}

func main() {
	A := new(TreeNode)
	A.Val = 1
	A.Left = new(TreeNode)
	A.Left.Val = 0
	A.Right = new(TreeNode)
	A.Right.Val = 1
	A.Left.Left = new(TreeNode)
	A.Left.Left.Val = -4
	A.Left.Right = new(TreeNode)
	A.Left.Right.Val = -3
	B := new(TreeNode)
	B.Val = 1
	B.Left = new(TreeNode)
	B.Left.Val = -4
	fmt.Print(isSubStructure(A, B))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
