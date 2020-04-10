package main

// https://leetcode-cn.com/problems/max-area-of-island/

// dfs 沉岛思想
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = maxFunc(getArea(grid, i, j), maxArea)
			}
		}
	}
	return maxArea
}

func getArea(grid [][]int, i, j int) int {
	if i == len(grid) || i < 0 {
		return 0
	} else if j == len(grid[0]) || j < 0 {
		return 0
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + getArea(grid, i-1, j) + getArea(grid, i+1, j) + getArea(grid, i, j-1) + getArea(grid, i, j+1)
	}
	return 0
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
