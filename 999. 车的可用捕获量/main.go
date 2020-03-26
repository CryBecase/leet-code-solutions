package main

// https://leetcode-cn.com/problems/available-captures-for-rook/

// 白车 R
// 黑车 r
// 空 .
// 白象 B
// 黑象 b
// 白卒 P
// 黑卒 p

func numRookCaptures(board [][]byte) int {
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}

	Rr, Rc := -1, -1

findR:
	for r, row := range board {
		for c, col := range row {
			if col == 'R' {
				Rr, Rc = r, c
				break findR
			}
		}
	}

	if Rr == -1 {
		return 0
	}

	result := 0

	for i := 0; i < 4; i++ {
		nr, nc := Rr, Rc
		for {
			nr += dr[i]
			nc += dc[i]
			if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) {
				if board[nr][nc] == '.' {
					continue
				} else if board[nr][nc] == 'B' || board[nr][nc] == 'b' || board[nr][nc] == 'P' {
					break
				} else if board[nr][nc] == 'p' {
					result++
					break
				}
			} else {
				break
			}
		}
	}

	return result
}

func main() {

}
