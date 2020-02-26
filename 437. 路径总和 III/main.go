package main

// https://leetcode-cn.com/problems/path-sum-iii/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = 0

func pathSum(root *TreeNode, sum int) int {
	result = 0
	DLR(root, sum)
	return result
}

// 前序遍历
func DLR(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	canGetSum(root, sum)
	DLR(root.Left, sum)
	DLR(root.Right, sum)
}

func canGetSum(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if sum == root.Val {
		result++
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	sum -= root.Val
	canGetSum(root.Left, sum)
	canGetSum(root.Right, sum)
}

func main() {

}
