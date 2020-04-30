package main

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return maxFunc(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
