package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	// dp[i][0] 表示第i天持股的最大利润
	// dp[i][1] 表示第i天不持股的最大利润

	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 0
	for i, price := range prices {
		if i == 0 {
			dp[i+1][0] = -price
			dp[i+1][1] = 0
		} else {
			dp[i+1][0] = maxFunc(dp[i][0], dp[i][1]-price)
			dp[i+1][1] = maxFunc(dp[i][1], dp[i][0]+prices[i])
		}
	}
	return dp[len(prices)][1]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
