package main

// https://leetcode-cn.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	if m < 2 || n < 2 {
		return 1
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for c := 1; c < m; c++ {
		dp[0][c] = 1
	}
	for r := 1; r < n; r++ {
		dp[r][0] = 1
	}
	for r := 1; r < n; r++ {
		for c := 1; c < m; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}
	return dp[n-1][m-1]
}

func main() {

}
