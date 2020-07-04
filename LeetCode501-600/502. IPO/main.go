package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode-cn.com/problems/ipo/

type Node struct {
	Capital int
	Profit  int
}

type MinCapitalHeap []Node

func (m MinCapitalHeap) Len() int {
	return len(m)
}

func (m MinCapitalHeap) Less(i, j int) bool {
	return m[i].Capital < m[j].Capital
}

func (m MinCapitalHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinCapitalHeap) Push(x interface{}) {
	*m = append(*m, x.(Node))
}

func (m *MinCapitalHeap) Pop() interface{} {
	res := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return res
}

type MaxProfitHeap []Node

func (m MaxProfitHeap) Len() int {
	return len(m)
}

func (m MaxProfitHeap) Less(i, j int) bool {
	return m[i].Profit > m[j].Profit
}

func (m MaxProfitHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxProfitHeap) Push(x interface{}) {
	*m = append(*m, x.(Node))
}

func (m *MaxProfitHeap) Pop() interface{} {
	res := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return res
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	nodes := make([]Node, 0)
	for i := 0; i < len(Profits); i++ {
		nodes = append(nodes, Node{
			Capital: Capital[i],
			Profit:  Profits[i],
		})
	}
	minCapitalHeap := MinCapitalHeap(nodes)
	heap.Init(&minCapitalHeap)
	maxProfitHeap := make(MaxProfitHeap, 0)
	for i := 0; i < k; i++ {
		for minCapitalHeap.Len() > 0 && minCapitalHeap[0].Capital <= W {
			heap.Push(&maxProfitHeap, heap.Pop(&minCapitalHeap))
		}
		if maxProfitHeap.Len() == 0 {
			return W
		}
		W += heap.Pop(&maxProfitHeap).(Node).Profit
	}
	return W
}

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}
