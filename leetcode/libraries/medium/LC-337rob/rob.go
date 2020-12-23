package main

/**
 * Definition for a binary tree node.
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left, right := dfs(root.Left), dfs(root.Right)
	selected := root.Val + left[1] + right[1]
	noselected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, noselected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
