package main

// https://leetcode-cn.com/problems/longest-univalue-path/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = 0

func longestUnivaluePath(root *TreeNode) int {
	result = 0
	dfs(root)
	return result
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left), dfs(root.Right)
	leftPath, rightPath := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		leftPath = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rightPath = right + 1
	}
	result = maxFunc(result, leftPath+rightPath)
	return maxFunc(leftPath, rightPath)
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
