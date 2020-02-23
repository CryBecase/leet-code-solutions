package main

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxFunc(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {

}
