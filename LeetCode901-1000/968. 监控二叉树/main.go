package main

import "math"

// https://leetcode-cn.com/problems/binary-tree-cameras/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const inf = math.MaxInt32 / 2

func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = minFunc(a, minFunc(la+rb, ra+lb))
		c = minFunc(a, lb+rb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
