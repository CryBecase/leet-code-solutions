package main

import "fmt"

// https://leetcode-cn.com/problems/integer-break/

/**

状态数组dp[i]表示：数字 i 拆分为至少两个正整数之和的最大乘积。为了方便计算，dp 的长度是 n + 1，值初始化为 1。

显然dp[2]等于 1，外层循环从 3 开始遍历，一直到 n 停止。
内层循环 j 从 1 开始遍历，一直到 i 之前停止，它代表着数字 i 可以拆分成 j + (i - j)。
但 j * (i - j)不一定是最大乘积，因为i-j不一定大于dp[i - j]（数字i-j拆分成整数之和的最大乘积），这里要选择最大的值作为 dp[i] 的结果。

*/

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = maxFunc(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func maxFunc(i1, i2, i3 int) int {
	largest := i1
	if i2 > i1 {
		largest = i2
	}
	if i3 > largest {
		return i3
	}
	return largest
}

func main() {
	integerBreak(10)
}
