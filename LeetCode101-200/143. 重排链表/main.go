package main

// https://leetcode-cn.com/problems/reorder-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	helper := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		helper = append(helper, cur)
		temp := cur.Next
		cur.Next = nil
		cur = temp
	}

	L, R := 0, len(helper)-1
	for L < R {
		helper[L].Next = helper[R]
		if L+1 != R {
			helper[R].Next = helper[L+1]
		}

		L++
		R--
	}
}

func main() {

}
