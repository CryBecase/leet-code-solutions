package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	p := h

	base := 0

	for l1 != nil || l2 != nil || base > 0 {
		if l1 != nil {
			base += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			base += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{
			Val:  base % 10,
			Next: p.Next,
		}
		base /= 10
		p = p.Next
	}
	return h.Next
}

func main() {

}
