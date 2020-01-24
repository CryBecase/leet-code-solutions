package main

/**
https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{
		Val:  0,
		Next: head,
	}

	result := preHead.Next.Next

	for preHead.Next != nil {
		if preHead.Next.Next == nil {
			break
		}
		r := preHead.Next.Next.Next
		l := preHead.Next
		preHead.Next = preHead.Next.Next
		preHead.Next.Next = l
		preHead.Next.Next.Next = r
		preHead = preHead.Next.Next
	}
	return result
}

func main() {

}
