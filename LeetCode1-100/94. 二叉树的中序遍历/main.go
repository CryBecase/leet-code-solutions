package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 颜色标记法 模拟递归
func inorderTraversal(root *TreeNode) []int {
	type Node struct {
		color int // color 为 1 表示进入递归，为 0 表示执行操作
		node  *TreeNode
	}

	stack := make([]*Node, 0)

	stack = append(stack, &Node{
		color: 1,
		node:  root,
	})
	result := make([]int, 0)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.node == nil {
			continue
		}
		if n.color == 0 {
			result = append(result, n.node.Val)
		} else {
			stack = append(stack, &Node{
				color: 1,
				node:  n.node.Right,
			})
			stack = append(stack, &Node{
				color: 0,
				node:  n.node,
			})
			stack = append(stack, &Node{
				color: 1,
				node:  n.node.Left,
			})
		}
	}
	return result
}

func main() {

}
