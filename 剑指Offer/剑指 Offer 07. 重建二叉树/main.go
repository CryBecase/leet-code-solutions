package main

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre   []int
	inMap map[int]int
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre = preorder
	// 用 map 记录位置，避免遍历查找
	inMap = make(map[int]int)
	for i, value := range inorder {
		inMap[value] = i
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preL, preR, inL, inR int) *TreeNode {
	if preL > preR || inL > inR {
		return nil
	}

	pivot := pre[preL]
	root := &TreeNode{Val: pivot}

	pivotIndex := inMap[pivot]

	root.Left = build(preL+1, preL+(pivotIndex-inL), inL, pivotIndex-1)
	root.Right = build(preL+(pivotIndex-inL)+1, preR, pivotIndex+1, inR)
	return root
}

func main() {

}
