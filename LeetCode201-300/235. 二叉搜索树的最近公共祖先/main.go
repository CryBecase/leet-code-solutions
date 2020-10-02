package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 找到第一个分岔点 就是答案
	res := root
	for {
		if res.Val > p.Val && res.Val > q.Val {
			res = res.Left
		} else if res.Val < p.Val && res.Val < q.Val {
			res = res.Right
		} else {
			return res
		}
	}
}

func main() {

}
