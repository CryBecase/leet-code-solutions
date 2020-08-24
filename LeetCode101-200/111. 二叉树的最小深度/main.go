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
	L, R := minDepth(root.Left), minDepth(root.Right)
	if L == 0 || R == 0 { // 此时 root 不是叶子 所以不能返回 1
		return L + R + 1
	}
	return minFunc(L, R) + 1
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
