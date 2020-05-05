package main

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	max := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= max {
			return false
		}
		max = root.Val
		root = root.Right
	}
	return true
}

func main() {

}
