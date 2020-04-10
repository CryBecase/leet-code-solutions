package main

import "fmt"

// https://leetcode-cn.com/problems/diameter-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := depth(root.Left, max), depth(root.Right, max)
	*max = maxFunc(*max, leftDepth+rightDepth)
	return maxFunc(leftDepth, rightDepth) + 1
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(diameterOfBinaryTree(&TreeNode{}))
}
