package main

// https://leetcode-cn.com/problems/word-search/

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	check := func(i, j, k int) bool { return false }
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 { // 单词存在于网格中
			return true
		}

		visited[i][j] = true
		defer func() { visited[i][j] = false }()

		for _, dir := range directions {
			if nr, nc := i+dir.x, j+dir.y; nr >= 0 && nr < h && nc >= 0 && nc < w && !visited[nr][nc] {
				if check(nr, nc, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {

}
