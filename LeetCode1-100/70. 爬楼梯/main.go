package main

// https://leetcode-cn.com/problems/climbing-stairs/description/

func climbStairs(n int) int {
	// dp[i] 表示到第i层有多少种方式
	dp := make([]int, 0)
	dp = append(dp, 0)
	dp = append(dp, 1)
	dp = append(dp, 2)

	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

func main() {

}
