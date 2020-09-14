package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	type coloredNode struct {
		color int
		node  *TreeNode
	}

	stack := make([]*coloredNode, 0)
	stack = append(stack, &coloredNode{
		color: 1,
		node:  root,
	})
	result := make([]int, 0)
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cn.node == nil {
			continue
		}
		if cn.color == 0 {
			result = append(result, cn.node.Val)
		} else {
			stack = append(stack, &coloredNode{
				color: 1,
				node:  cn.node.Right,
			})
			stack = append(stack, &coloredNode{
				color: 0,
				node:  cn.node,
			})
			stack = append(stack, &coloredNode{
				color: 1,
				node:  cn.node.Left,
			})
		}
	}
	return result
}

func main() {

}
