package main

import "container/list"

// https://leetcode-cn.com/problems/implement-stack-using-queues/

type MyStack struct {
	q1         *list.List
	q2         *list.List
	topElement int
}

func Constructor() MyStack {
	return MyStack{q1: list.New(), q2: list.New(), topElement: -1}
}

func (this *MyStack) Push(x int) {
	this.q1.PushBack(x)
	this.topElement = x
}

func (this *MyStack) Pop() int {
	if this.q1.Len() == 0 {
		return -1
	}

	for this.q1.Len() > 1 {
		e := this.q1.Front()
		this.q1.Remove(e)
		this.q2.PushBack(e.Value.(int))
		this.topElement = e.Value.(int)
	}

	x := this.q1.Front()
	this.q1.Remove(x)
	this.q1, this.q2 = this.q2, this.q1
	if this.q1.Len() == 0 {
		this.topElement = -1
	}

	return x.Value.(int)
}

func (this *MyStack) Top() int {
	return this.topElement
}

func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
