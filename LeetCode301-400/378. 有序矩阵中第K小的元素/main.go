package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/

type Position struct {
	Value int
	Raw   int
	Col   int
}

type minHeap []Position

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i].Value < (*m)[j].Value
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(Position))
}

func (m *minHeap) Pop() interface{} {
	t := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return t
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	minH := make(minHeap, 0)
	for i := 0; i < len(matrix); i++ {
		heap.Push(&minH, Position{
			Value: matrix[i][0],
			Raw:   i,
			Col:   0,
		})
	}
	for i := 0; i < k; i++ {
		p := heap.Pop(&minH).(Position)
		if i == k-1 {
			return p.Value
		}
		if p.Col == len(matrix[p.Raw])-1 {
			// 如果是这行的最后一个
			continue
		}
		heap.Push(&minH, Position{
			Value: matrix[p.Raw][p.Col+1],
			Raw:   p.Raw,
			Col:   p.Col + 1,
		})
	}
	return 0
}

func main() {
	fmt.Println(kthSmallest([][]int{{-5}}, 1))
}
