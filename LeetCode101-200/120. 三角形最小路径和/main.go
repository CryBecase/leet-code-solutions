package main

import "math"

// https://leetcode-cn.com/problems/triangle/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	// dp[i][j] 表示从顶点到第i+1行第j+1个元素需要走的最短路经和
	dp := make([][]int, len(triangle))
	for i, row := range triangle {
		dp[i] = make([]int, len(row))
	}

	dp[0][0] = triangle[0][0]
	min := math.MaxInt32
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				continue
			}
			dp[i][j] = minFunc(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		min = minFunc(min, dp[len(triangle)-1][j])
	}
	return min
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
