package main

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	v := make([]int, 0)
	LDR(root, &v)
	L, R := 0, len(v)-1
	for L < R {
		if v[L]+v[R] < k {
			L++
		} else if v[L]+v[R] > k {
			R--
		} else {
			return true
		}
	}
	return false
}

func LDR(root *TreeNode, v *[]int) {
	if root == nil {
		return
	}
	LDR(root.Left, v)
	*v = append(*v, root.Val)
	LDR(root.Right, v)
}

func main() {

}
