package main

/**
https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	currentLayer := []*TreeNode{root}
	var nextLayer []*TreeNode
	var sum float64
	for len(currentLayer) != 0 {
		// 广度优先遍历 BFS
		for _, v := range currentLayer {
			if v.Left != nil {
				nextLayer = append(nextLayer, v.Left)
			}
			if v.Right != nil {
				nextLayer = append(nextLayer, v.Right)
			}
			sum = sum + float64(v.Val)
		}
		result = append(result, sum/float64(len(currentLayer)))
		sum = 0
		currentLayer = nextLayer
		nextLayer = []*TreeNode{}
	}
	return result
}

func main() {

}
