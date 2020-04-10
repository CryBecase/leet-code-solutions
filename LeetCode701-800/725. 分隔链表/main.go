package main

// https://leetcode-cn.com/problems/split-linked-list-in-parts/description/

/**
root 的长度范围： [0, 1000].
输入的每个节点的大小范围：[0, 999].
k 的取值范围： [1, 50].
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	// 确实 原链表 长度
	length := 0
	p := root
	for p != nil {
		length++
		p = p.Next
	}
	// 确定分组长度
	var leftNodeLength int  // 左边链表的长度
	leftNodeCount := 1      // 左边链表的个数
	var rightNodeLength int // 右边链表的长度
	if length <= k {
		leftNodeLength = 1
		rightNodeLength = leftNodeLength
	} else {
		if length%k == 0 { // 4/2 = 2   [2,2]
			leftNodeLength = length / k
			rightNodeLength = leftNodeLength
		} else { // 5/2 = 2···1 [3, 2]   5/3 = 1···2 [2, 2, 1]
			leftNodeLength = length/k + 1
			rightNodeLength = leftNodeLength - 1
			// 确认左边这类链表的个数
			leftNodeCount = length % k
		}
	}

	result := make([]*ListNode, k, k)
	p = root
	for i := 0; i < k && p != nil; i++ {
		result[i] = p
		if i < leftNodeCount {
			for j := 0; j < leftNodeLength-1 && p != nil; j++ {
				p = p.Next
			}
		} else {
			for j := 0; j < rightNodeLength-1 && p != nil; j++ {
				p = p.Next
			}
		}
		if p != nil {
			tmp := p.Next
			p.Next = nil
			p = tmp
		}
	}
	return result
}

func main() {

}
