package main

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if n == 0 {
		return result
	}
	return append(result, newTree(1, n)...)
}

func newTree(start, end int) []*TreeNode {
	result := make([]*TreeNode, 0)

	if start > end {
		result = append(result, nil)
		return result
	}
	if start == end {
		result = append(result, &TreeNode{Val: start})
		return result
	}

	for i := start; i <= end; i++ {
		leftTrees := newTree(start, i-1)
		rightTress := newTree(i+1, end)
		for l := range leftTrees {
			for r := range rightTress {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  leftTrees[l],
					Right: rightTress[r],
				})
			}
		}
	}
	return result
}

func main() {

}
