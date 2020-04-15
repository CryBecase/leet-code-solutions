package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/01-matrix/

func updateMatrix(matrix [][]int) [][]int {
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	rowLength := len(matrix)
	cowLength := len(matrix[0])

	visited := make([][]int, rowLength)
	for i := range visited {
		visited[i] = make([]int, cowLength)
	}

	queue := list.New()
	for r, row := range matrix {
		for c, col := range row {
			if col == 0 {
				queue.PushBack([]int{r, c})
				visited[r][c] = 1
			}
		}
	}

	for queue.Len() > 0 {
		point := queue.Front().Value.([]int)
		queue.Remove(queue.Front())
		for i := 0; i < 4; i++ {
			nr := point[0] + dr[i]
			nc := point[1] + dc[i]
			if nr >= 0 && nr < rowLength && nc >= 0 && nc < cowLength && matrix[nr][nc] != 0 && visited[nr][nc] == 0 {
				matrix[nr][nc] = matrix[point[0]][point[1]] + 1
				queue.PushBack([]int{nr, nc})
				visited[nr][nc] = 1
			}
		}
	}

	return matrix
}

func main() {
	// [[0,0,0],[0,1,0],[1,1,1]]

	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))
}
