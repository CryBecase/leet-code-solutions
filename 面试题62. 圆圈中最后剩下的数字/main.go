package main

import "fmt"

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

// (当前index + m) % 上一轮剩余数字的个数
func lastRemaining(n int, m int) int {
	ans := 0
	// 最后一轮剩下2个人，所以从2开始反推
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	fmt.Println(lastRemaining(5, 1))
}
