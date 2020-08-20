package main

// https://leetcode-cn.com/problems/minesweeper/

func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}

	row, col := click[0], click[1]
	clickPosition := board[row][col]
	if clickPosition == 'M' {
		return handleM(board, row, col)
	}

	if clickPosition == 'E' {
		return handleE(board, row, col)
	}

	return board
}

func handleM(board [][]byte, row, col int) [][]byte {
	board[row][col] = 'X'
	return board
}

func handleE(board [][]byte, row, col int) [][]byte {
	if board[row][col] != 'E' {
		return board
	}

	v := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	boomCount := 0
	for i := 0; i < 8; i++ {
		nc, nr := row+v[i][0], col+v[i][1]
		if nc >= 0 && nc < len(board) && nr >= 0 && nr < len(board[0]) && board[nc][nr] == 'M' {
			boomCount++
		}
	}
	if boomCount > 0 {
		board[row][col] = byte('0' + boomCount)
		return board
	}

	board[row][col] = 'B'

	for i := 0; i < 8; i++ {
		nc, nr := row+v[i][0], col+v[i][1]
		if nc >= 0 && nc < len(board) && nr >= 0 && nr < len(board[0]) {
			board = handleE(board, nc, nr)
		}
	}
	return board
}

func main() {

}
