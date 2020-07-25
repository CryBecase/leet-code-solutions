package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][k][0] 表示第i+1天允许k笔交易且手上没有股票的时候的最大利润
	// dp[i][k][1] 表示第i+1天允许k笔交易且手上有股票的时候的最大利润
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, 3)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			dp[i][k][0] = maxFunc(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = maxFunc(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
