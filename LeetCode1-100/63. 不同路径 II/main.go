package main

// https://leetcode-cn.com/problems/unique-paths-ii/

// dp[r][c] 表示 到达第r行第c列有多少种走法
// dp[r][c] = dp[r-1][c] + dp[r][c-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	dp[0][0] = 1
	for r := 1; r < len(obstacleGrid); r++ {
		if obstacleGrid[r][0] == 1 {
			continue
		}
		dp[r][0] = dp[r-1][0]
	}
	for c := 1; c < len(obstacleGrid[0]); c++ {
		if obstacleGrid[0][c] == 1 {
			continue
		}
		dp[0][c] = dp[0][c-1]
	}
	for r := 1; r < len(obstacleGrid); r++ {
		for c := 1; c < len(obstacleGrid[r]); c++ {
			if obstacleGrid[r][c] == 1 {
				continue
			}
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	uniquePathsWithObstacles([][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	})
}
