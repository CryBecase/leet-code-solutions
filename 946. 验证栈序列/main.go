package main

// https://leetcode-cn.com/problems/validate-stack-sequences/

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	size := 0
	for i, j := 0, 0; i < len(pushed); i++ {
		pushed[size] = pushed[i]
		size++
		// 每push一个数 都检查是否立即弹出 以及之后的数是否跟着弹出
		for size != 0 && pushed[size-1] == popped[j] {
			size--
			j++
		}
	}
	return size == 0
}

func main() {

}
