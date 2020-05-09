package main

import (
	"container/list"
)

// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()

	queue.PushBack(root)

	for {
		first := queue.Front().Value.(*TreeNode)
		length := queue.Len()

		for i := 0; i < length; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		if queue.Len() == 0 {
			return first.Val
		}
	}
}

func main() {

}
