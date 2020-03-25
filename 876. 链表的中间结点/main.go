package main

// https://leetcode-cn.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		fast = fast.Next
		if fast == nil {
			return slow
		}
		fast = fast.Next
		if fast == nil {
			return slow.Next
		}
		slow = slow.Next
	}
}

func main() {

}
