package main

// https://leetcode-cn.com/problems/house-robber-iii/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	return robInternal(root, &m)
}

func robInternal(root *TreeNode, m *map[*TreeNode]int) int {
	if tmp, ok := (*m)[root]; ok {
		return tmp
	}

	if root == nil {
		return 0
	}
	v := root.Val
	if root.Left != nil {
		v += robInternal(root.Left.Left, m) + robInternal(root.Left.Right, m)
	}
	if root.Right != nil {
		v += robInternal(root.Right.Left, m) + robInternal(root.Right.Right, m)
	}
	v2 := robInternal(root.Left, m) + robInternal(root.Right, m)
	(*m)[root] = maxFunc(v, v2)
	return maxFunc(v, v2)
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
