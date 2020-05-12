package main

import (
	"container/list"
)

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	realStack *list.List
	minStack  *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		realStack: list.New(),
		minStack:  list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.realStack.PushBack(x)
	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		min := this.minStack.Back().Value.(int)
		if x < min {
			this.minStack.PushBack(x)
		} else {
			this.minStack.PushBack(min)
		}
	}
}

func (this *MinStack) Pop() {
	this.realStack.Remove(this.realStack.Back())
	this.minStack.Remove(this.minStack.Back())
}

func (this *MinStack) Top() int {
	return this.realStack.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.minStack.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
