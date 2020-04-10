package main

// https://leetcode-cn.com/problems/sum-of-left-leaves/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = 0

func sumOfLeftLeaves(root *TreeNode) int {
	result = 0
	LDR(root, false)
	return result
}

func LDR(root *TreeNode, b bool) {
	if root == nil {
		return
	}
	if b && root.Left == nil && root.Right == nil {
		result += root.Val
	}
	LDR(root.Left, true)
	LDR(root.Right, false)
}

func main() {

}
