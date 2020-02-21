package main

// https://leetcode-cn.com/problems/odd-even-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstEvenPointer := head.Next
	oddPointer, evenPointer := head, head.Next
	for {
		if oddPointer != nil && oddPointer.Next != nil && oddPointer.Next.Next != nil {
			oddPointer.Next = oddPointer.Next.Next
			oddPointer = oddPointer.Next
		} else {
			break
		}

		if evenPointer != nil && evenPointer.Next != nil {
			evenPointer.Next = evenPointer.Next.Next
			evenPointer = evenPointer.Next
		} else {
			break
		}
	}
	oddPointer.Next = firstEvenPointer
	return head
}

func main() {

}
