package main

/**
https://leetcode-cn.com/problems/add-two-numbers-ii/description/

给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//挨个入栈
	var l1Stack = make([]*ListNode, 0)
	var l2Stack = make([]*ListNode, 0)

	for l1 != nil {
		l1Stack = append(l1Stack, l1)
		l1 = l1.Next
	}

	for l2 != nil {
		l2Stack = append(l2Stack, l2)
		l2 = l2.Next
	}

	carry := 0
	var node *ListNode
	len1, len2 := len(l1Stack), len(l2Stack)
	for len1 > 0 || len2 > 0 {
		num1, num2 := 0, 0
		if len1 > 0 {
			num1 = l1Stack[len1-1].Val
		}
		if len2 > 0 {
			num2 = l2Stack[len2-1].Val
		}
		total := num2 + num1

		node = &ListNode{(total + carry) % 10, node}
		if total+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		len1--
		len2--
	}
	if carry == 1 {
		node = &ListNode{1, node}
	}
	return node
}

func main() {

}
