package main

// https://leetcode-cn.com/problems/linked-list-cycle-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {

}
