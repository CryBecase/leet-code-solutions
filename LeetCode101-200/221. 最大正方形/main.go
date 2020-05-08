package main

// https://leetcode-cn.com/problems/maximal-square/

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	max := 0
	for r, row := range matrix {
		dp[r] = make([]int, len(row))
		for c, col := range row {
			dp[r][c] = int(col - '0')
			if dp[r][c] == 1 {
				max = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = minFunc(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				max = maxFunc(max, dp[i][j])
			}
		}
	}
	return max * max
}

func minFunc(i, j, z int) int {
	least := i
	if j < i {
		least = j
	}
	if z < least {
		return z
	}
	return least
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
