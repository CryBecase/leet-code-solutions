package main

// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = 0 + grid[0][0]
	for r := 1; r < len(grid); r++ {
		dp[r][0] = dp[r-1][0] + grid[r][0]
	}
	for c := 1; c < len(grid[0]); c++ {
		dp[0][c] = dp[0][c-1] + grid[0][c]
	}
	for r := 1; r < len(grid); r++ {
		for c := 1; c < len(grid[r]); c++ {
			dp[r][c] = minFunc(dp[r-1][c]+grid[r][c], dp[r][c-1]+grid[r][c])
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
