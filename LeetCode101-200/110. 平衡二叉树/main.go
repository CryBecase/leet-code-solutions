package main

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
	ok, _ := helper(root)
	return ok
}

func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftOk, leftDepth := helper(root.Left)
	rightOk, rightDepth := helper(root.Right)
	return leftOk && rightOk && abs(leftDepth, rightDepth) <= 1, maxFunc(leftDepth, rightDepth) + 1
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
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
