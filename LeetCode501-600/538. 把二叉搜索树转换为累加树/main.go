package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	RDL(root)
	return root
}

func RDL(root *TreeNode) {
	if root == nil {
		return
	}
	RDL(root.Right)
	sum += root.Val
	root.Val = sum
	RDL(root.Left)
}

func main() {

}
