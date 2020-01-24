package main

/**
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
双指针
p1先走n次 p2再走 等p1 == nil的时候 p2就是倒数第n个节点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	if p1 == nil {
		return head.Next
	}
	for p1 != nil {
		if p1.Next == nil {
			p2.Next = p2.Next.Next
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return head
}

func main() {

}
