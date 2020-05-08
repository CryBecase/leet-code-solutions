package main

// https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix))

	for r, row := range matrix {
		dp[r] = make([]int, len(row))
		for c, col := range row {
			dp[r][c] = col
		}
	}

	res := 0

	for i := 0; i < len(matrix[0]); i++ {
		res += matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		res += matrix[i][0]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = minFunc(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				res += dp[i][j]
			}
		}
	}

	return res
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

func main() {

}
