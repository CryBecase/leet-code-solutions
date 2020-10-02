package main

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	head := root
	for {
		if val < head.Val {
			if head.Left == nil {
				head.Left = &TreeNode{Val: val}
				break
			}
			head = head.Left
		} else {
			if head.Right == nil {
				head.Right = &TreeNode{Val: val}
				break
			}
			head = head.Right
		}
	}

	return root
}

func main() {

}
