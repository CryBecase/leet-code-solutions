package main

// https://leetcode-cn.com/problems/path-sum-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var (
		path []int
		res  [][]int
	)
	dfs := func(*TreeNode, int) {}
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if sum == root.Val && root.Left == nil && root.Right == nil {
			res = append(res, append([]int(nil), path...))
			return
		}

		sum -= root.Val
		if root.Left != nil {
			dfs(root.Left, sum)
			path = path[:len(path)-1]
		}

		if root.Right != nil {
			dfs(root.Right, sum)
			path = path[:len(path)-1]
		}
	}
	dfs(root, sum)
	return res
}

func main() {

}
