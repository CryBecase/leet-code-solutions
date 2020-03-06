package main

// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	leftValue, rightValue := root.Left.Val, root.Right.Val
	if leftValue == root.Val {
		leftValue = findSecondMinimumValue(root.Left)
	}
	if rightValue == root.Val {
		rightValue = findSecondMinimumValue(root.Right)
	}
	if leftValue != -1 && rightValue != -1 {
		return minFunc(leftValue, rightValue)
	}
	if leftValue != -1 {
		return leftValue
	}
	return rightValue
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
