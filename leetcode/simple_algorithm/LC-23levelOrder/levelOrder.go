package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	roots := []*TreeNode{root}
	for i := 0; len(roots) > 0; i++ {
		result = append(result, []int{})
		change := make([]*TreeNode, 0)
		for j := 0; j < len(roots); j++ {
			result[i] = append(result[i], roots[j].Val)
			if roots[j].Left != nil {
				change = append(change, roots[j].Left)
			}
			if roots[j].Right != nil {
				change = append(change, roots[j].Right)
			}
		}
		roots = change
	}
	return result
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Left.Val = 3
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	root.Right.Left.Val = 5
	root.Right.Right.Val = 6
	levelOrder(root)
}
