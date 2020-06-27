package main

// https://leetcode-cn.com/problems/same-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return preOrder(p, q)
}

func preOrder(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return preOrder(p.Left, q.Left) && preOrder(p.Right, q.Right)
}

func main() {

}
