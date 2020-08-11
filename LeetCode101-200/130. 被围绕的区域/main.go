package main

// https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte) {
	if len(board) < 3 {
		return
	}
	if len(board[0]) < 3 {
		return
	}
	copyBoard := make([][]byte, len(board))
	for i := range copyBoard {
		copyBoard[i] = make([]byte, len(board[i]))
		for j := range copyBoard[i] {
			copyBoard[i][j] = board[i][j]
		}
	}

	for i := 0; i < len(copyBoard[0]); i++ {
		if copyBoard[0][i] == 'O' {
			dfs(copyBoard, 0, i)
		}
		if copyBoard[len(copyBoard)-1][i] == 'O' {
			dfs(copyBoard, len(copyBoard)-1, i)
		}
	}
	for i := 0; i < len(copyBoard); i++ {
		if copyBoard[i][0] == 'O' {
			dfs(copyBoard, i, 0)
		}
		if copyBoard[i][len(copyBoard[i])-1] == 'O' {
			dfs(copyBoard, i, len(copyBoard[i])-1)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if copyBoard[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] != 'O' {
		return
	}
	board[r][c] = 'p'
	dfs(board, r+1, c)
	dfs(board, r-1, c)
	dfs(board, r, c+1)
	dfs(board, r, c-1)
}

func main() {
	solve([][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	})
}
