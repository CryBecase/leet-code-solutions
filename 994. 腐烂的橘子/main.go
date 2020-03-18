package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/rotting-oranges/

type position struct {
	r int
	c int
}

func orangesRotting(grid [][]int) int {
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}
	// (+1,0) (0,+1) (-1,0) (0,-1)
	R, C := len(grid), len(grid[0])

	list := list.New()
	depth := make(map[position]int)

	// 保存坏掉的水果
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 2 {
				list.PushBack(position{r, c})
				depth[position{r, c}] = 0
			}
		}
	}
	// BFS广度优先遍历
	result := 0
	for list.Len() > 0 {
		code := list.Front().Value.(position)
		list.Remove(list.Front())
		r, c := code.r, code.c
		for i := 0; i < 4; i++ {
			nr, nc := r+dr[i], c+dc[i]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				ncode := position{nr, nc}
				list.PushBack(ncode)
				depth[ncode] = depth[code] + 1
				result = depth[ncode]
			}
		}
	}
	// 检查有没有新鲜的
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 2, 2, 1, 1}}))
}
