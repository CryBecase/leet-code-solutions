package main

// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	uniMap := make(map[int]int)
	node := &ListNode{
		Val:  0,
		Next: head,
	}
	for node != nil && node.Next != nil {
		for {
			if _, ok := uniMap[node.Next.Val]; ok {
				// 重复
				node.Next = node.Next.Next
				if node.Next == nil {
					break
				}
			} else {
				uniMap[node.Next.Val] = 1
				break
			}
		}
		node = node.Next
	}
	return head
}

func main() {

}
