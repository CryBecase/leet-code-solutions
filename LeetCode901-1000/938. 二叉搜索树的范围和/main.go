package main

// https://leetcode-cn.com/problems/range-sum-of-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	add := func(v int) {
		sum += v
	}
	preOrder(root, L, R, add)
	return sum
}

func preOrder(root *TreeNode, L, R int, add func(v int)) {
	if root == nil {
		return
	}
	if root.Val > L {
		preOrder(root.Left, L, R, add)
	}
	if root.Val >= L && root.Val <= R {
		add(root.Val)
	}
	if root.Val < R {
		preOrder(root.Right, L, R, add)
	}
}

func main() {

}
