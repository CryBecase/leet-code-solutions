package main

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

type CQueue struct {
	PopStack  []int
	PushStack []int
}

func Constructor() CQueue {
	return CQueue{
		PopStack:  make([]int, 0),
		PushStack: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.PushStack = append(this.PushStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.PopStack) == 0 {
		if len(this.PushStack) == 0 {
			return -1
		}
		length := len(this.PushStack)
		for i := length - 1; i >= 0; i-- {
			this.PopStack = append(this.PopStack, this.PushStack[i])
			this.PushStack = this.PushStack[:i]
		}
	}
	res := this.PopStack[len(this.PopStack)-1]
	this.PopStack = this.PopStack[:len(this.PopStack)-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {

}
