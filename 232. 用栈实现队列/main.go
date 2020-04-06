package main

// https://leetcode-cn.com/problems/implement-queue-using-stacks/

type Stack struct {
	list []int
	size int
	cap  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		list: make([]int, cap, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	}
	r := s.list[s.size-1]
	s.size--
	return r
}

func (s *Stack) Push(data int) {
	if s.size+1 > s.cap {
		s.cap *= 2
		newlist := make([]int, s.cap, s.cap)
		copy(newlist, s.list)
		s.list = newlist
	}
	s.list[s.size] = data
	s.size++
}

func (s *Stack) Peek() int {
	if s.size == 0 {
		return -1
	}
	return s.list[s.size-1]
}

type MyQueue struct {
	s1 *Stack
	s2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(100),
		s2: NewStack(100),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.size == 0 {
		size := this.s1.size
		for i := 0; i < size; i++ {
			this.s2.Push(this.s1.Pop())
		}
	}
	r := this.s2.Pop()
	return r
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.s2.size == 0 {
		size := this.s1.size
		for i := 0; i < size; i++ {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.s1.size == 0 && this.s2.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

}
