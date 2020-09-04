package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/binary-tree-paths/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	path := make([]string, 0)
	res := make([]string, 0)
	vlr := func(*TreeNode) {}
	vlr = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right == nil {
			res = append(res, strings.Join(path, "->"))
		}
		vlr(root.Left)
		vlr(root.Right)
		path = path[:len(path)-1]
	}
	vlr(root)
	return res
}

func main() {

}
