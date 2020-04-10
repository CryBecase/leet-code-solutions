package main

/**
https://leetcode-cn.com/problems/reverse-linked-list/description/

反转一个单链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 三指针
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var leftNode *ListNode
	curNode := head
	rightNode := curNode.Next
	for rightNode != nil {
		if leftNode == nil {
			curNode.Next = nil
		}
		temp := rightNode.Next

		rightNode.Next = curNode

		leftNode = curNode
		curNode = rightNode
		rightNode = temp
	}
	return curNode
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseList2(next)
	next.Next = head
	head.Next = nil
	return newHead
}

// 头插法
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	for head != nil {
		next := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = next
	}
	return newHead.Next
}

func main() {

}
