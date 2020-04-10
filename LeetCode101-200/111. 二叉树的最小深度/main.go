package main

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 { // 此时 root 不是叶子 所以不能返回 1
		return left + right + 1
	}
	return minFunc(left, right) + 1
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
