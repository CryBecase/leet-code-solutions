package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0] 表示第i天没有股票时的最大利润
	// dp[i][1] 表示第i天有股票时的最大利润
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxFunc(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxFunc(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
