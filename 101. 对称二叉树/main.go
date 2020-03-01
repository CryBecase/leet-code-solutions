package main

// https://leetcode-cn.com/problems/symmetric-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isReverseEqual(root.Left, root.Right)
}

func isReverseEqual(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isReverseEqual(root1.Left, root2.Right) && isReverseEqual(root1.Right, root2.Left)
}

func main() {

}
