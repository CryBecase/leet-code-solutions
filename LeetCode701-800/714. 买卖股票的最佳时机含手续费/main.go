package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0] 表示第i+1天没有股票时的最大利润
	// dp[i][1]表示第i+1天有股票时的最大利润
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxFunc(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxFunc(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
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
