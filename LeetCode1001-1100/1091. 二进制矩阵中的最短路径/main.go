package main

import "container/list"

// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/

func shortestPathBinaryMatrix(grid [][]int) int {
	rowLength := len(grid)
	if rowLength == 0 {
		return -1
	}
	colLength := len(grid[0])
	if colLength != rowLength || grid[0][0] == 1 || grid[rowLength-1][colLength-1] == 1 {
		return -1
	}

	if rowLength == 1 {
		return 1
	}

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	queue := list.New()
	queue.PushBack([]int{0, 0})

	term := 1
	for queue.Len() > 0 {
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			p := queue.Front().Value.([]int)
			queue.Remove(queue.Front())
			for j := 0; j < 8; j++ {
				nr, nc := p[0]+dr[j], p[1]+dc[j]
				if nr >= 0 && nr < rowLength && nc >= 0 && nc < colLength && grid[nr][nc] == 0 {
					if nr == rowLength-1 && nc == colLength-1 {
						return term + 1
					}
					queue.PushBack([]int{nr, nc})
					grid[nr][nc] = term
				}
			}
		}
		term++
	}
	return -1
}

func main() {

}
