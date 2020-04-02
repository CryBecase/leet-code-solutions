package main

// https://leetcode-cn.com/problems/game-of-life/

func gameOfLife(board [][]int) {
	var result [][]int

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for r, row := range board {
		var arr []int
		for c, col := range row {
			live := 0
			for i := 0; i < 8; i++ {
				nr := r + dr[i]
				nc := c + dc[i]
				if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) && board[nr][nc] == 1 {
					live++
				}
			}
			if live < 2 || live > 3 {
				arr = append(arr, 0)
			} else if live == 3 {
				arr = append(arr, 1)
			} else {
				arr = append(arr, col)
			}
		}
		result = append(result, arr)
	}
	copy(board, result)
}

func main() {

}
