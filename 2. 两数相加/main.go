package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := &ListNode{}
	root := result
	carry := 0
	v := 0

	for l1 != nil && l2 != nil {
		v = l1.Val + l2.Val + carry
		root.Val = v % 10
		carry = v / 10
		l1 = l1.Next
		l2 = l2.Next
		
		root.Next = &ListNode{}
		root = root.Next
	}

	return result
}

func main() {

}
