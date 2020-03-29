package main

// https://leetcode-cn.com/problems/as-far-from-land-as-possible/

type position struct {
	row int
	col int
}

// BFS 用slice模拟队列
func maxDistance(grid [][]int) int {
	var queue []position
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}
	for rowIndex, row := range grid {
		for colIndex, value := range row {
			if value == 1 {
				queue = append(queue, position{rowIndex, colIndex})
			}
		}
	}

	if len(queue) == 0 || len(queue) == len(grid)*len(grid) {
		return -1
	}

	sideLength := len(grid)
	depth := -1
	for len(queue) > 0 {
		var nqueue []position
		depth++
		for _, node := range queue {
			for i := 0; i < 4; i++ {
				nr, nc := node.row+dr[i], node.col+dc[i]
				if nr >= 0 && nr < sideLength && nc >= 0 && nc < sideLength && grid[nr][nc] == 0 {
					grid[nr][nc] = 2
					nqueue = append(nqueue, position{nr, nc})
				}
			}
		}
		queue = nqueue
	}

	return depth
}

func main() {

}
