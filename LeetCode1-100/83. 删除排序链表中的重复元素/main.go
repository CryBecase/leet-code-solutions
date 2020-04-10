package main

/**
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/

给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

logback
swagger
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curHead := head
	for curHead != nil && curHead.Next != nil {
		if curHead.Val == curHead.Next.Val {
			curHead.Next = curHead.Next.Next
		} else {
			curHead = curHead.Next
		}
	}
	return head
}

// 递归
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates2(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	} else {
		return head
	}
}

func main() {

}
