package main

// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	path := make([]int, 0)
	sum := 0
	dfs := func(*TreeNode) {}
	dfs = func(root *TreeNode) {
		if root == nil {
			s := 0
			for i, n := range path {
				base := 1
				for j := 1; j < len(path)-i; j++ {
					base *= 10
				}
				s += n * base
			}
			sum += s
			return
		}
		path = append(path, root.Val)

		if root.Right == nil && root.Left == nil {
			dfs(nil)
		}

		if root.Right != nil {
			dfs(root.Right)
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		path = path[:len(path)-1]
	}
	dfs(root)
	return sum
}

func main() {

}
