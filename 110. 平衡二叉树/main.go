package main

import "math"

// https://leetcode-cn.com/problems/balanced-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// 对每一个根 判断它的左右孩子是否高度相差1 递归
	if root == nil {
		return true
	}
	if math.Abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return maxFunc(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxFunc(i, j float64) float64 {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {

}
