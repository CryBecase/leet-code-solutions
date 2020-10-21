package main

import "math"

// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt32
	pre := -1
	LDR := func(*TreeNode) {}
	LDR = func(root *TreeNode) {
		if root == nil {
			return
		}
		LDR(root.Left)
		if pre != -1 {
			res = minFunc(res, root.Val-pre)
		}
		pre = root.Val
		LDR(root.Right)
	}
	LDR(root)
	return res
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
