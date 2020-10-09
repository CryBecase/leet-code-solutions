package main

// https://leetcode-cn.com/problems/linked-list-cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
		fast = fast.Next
		if slow == fast {
			return true
		}

		slow = slow.Next
	}
	return false
}

func main() {

}
