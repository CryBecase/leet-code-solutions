package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0] 第i+1天 手上持有股票的最大收益
	// dp[i][1] 第i+1天 手上不持有股票，并且处于冷冻期中的累计最大收益
	// dp[i][2] 第i+1天 手上不持有股票，并且不在冷冻期中的累计最大收益
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxFunc(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = maxFunc(dp[i-1][1], dp[i-1][2])
		max = maxFunc(max, maxFunc(maxFunc(dp[i][0], dp[i][1]), dp[i][2]))
	}
	return max
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{2, 1}))
}
