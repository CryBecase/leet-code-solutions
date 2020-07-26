package main

// https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/

var (
	dirs       = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, cols)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans = maxFunc(ans, dfs(matrix, i, j, memo))
		}
	}
	return ans
}

func dfs(matrix [][]int, row, col int, memo [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}
	memo[row][col]++
	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && matrix[newRow][newCol] > matrix[row][col] {
			memo[row][col] = maxFunc(memo[row][col], dfs(matrix, newRow, newCol, memo)+1)
		}
	}
	return memo[row][col]
}

func maxFunc(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
