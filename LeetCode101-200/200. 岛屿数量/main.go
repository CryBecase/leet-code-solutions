package main

// https://leetcode-cn.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	result := 0

	for r, row := range grid {
		for c, col := range row {
			if col == '1' {
				bfs(grid, r, c)
				result++
			}
		}
	}
	return result
}

func bfs(grid [][]byte, r, c int) {
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		nr, nc := r+dr[i], c+dc[i]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] == '1' {
			grid[nr][nc] = '0'
			bfs(grid, nr, nc)
		}
	}

}

func main() {

}
