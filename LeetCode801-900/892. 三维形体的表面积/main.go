package main

// https://leetcode-cn.com/problems/surface-area-of-3d-shapes/

func surfaceArea(grid [][]int) int {
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}

	width := len(grid)
	if width == 0 {
		return 0
	}
	result := 0

	for r, row := range grid {
		for c, v := range row {
			if v > 0 {
				result += 2
				for i := 0; i < 4; i++ {
					nr := r + dr[i]
					nc := c + dc[i]
					if nr >= 0 && nr < width && nc >= 0 && nc < width {
						result += maxFunc(v-grid[nr][nc], 0)
					} else {
						result += v
					}
				}
			}
		}
	}

	return result
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
