package main

// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return build(arr, 0, len(arr)-1)
}

func build(arr []int, L, R int) *TreeNode {
	if L >= len(arr) || L > R {
		return nil
	}

	if L == R {
		return &TreeNode{Val: arr[L]}
	}

	mid := (L + R) / 2
	if (R-L)%2 == 1 {
		mid += 1
	}
	tree := &TreeNode{}
	tree.Val = arr[mid]
	tree.Left = build(arr, L, mid-1)
	tree.Right = build(arr, mid+1, R)
	return tree
}

func main() {

}
