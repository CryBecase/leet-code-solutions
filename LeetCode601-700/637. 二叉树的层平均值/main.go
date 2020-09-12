package main

// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	queue := []*TreeNode{root}
	var nextLayer []*TreeNode
	var sum float64
	for len(queue) != 0 {
		for _, v := range queue {
			if v.Left != nil {
				nextLayer = append(nextLayer, v.Left)
			}
			if v.Right != nil {
				nextLayer = append(nextLayer, v.Right)
			}
			sum += float64(v.Val)
		}
		result = append(result, sum/float64(len(queue)))
		sum = 0
		queue = nextLayer
		nextLayer = []*TreeNode{}
	}
	return result
}

func main() {

}
