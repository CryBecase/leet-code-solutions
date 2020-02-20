package main

// https://leetcode-cn.com/problems/palindrome-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head.Next // 慢指针一次1步 快指针一次2步 快指针走到nil 慢指针就是后半部分
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	// 现在 slow 在后半部分
	// 反转链表
	var pre *ListNode
	cur := slow
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

func main() {

}
