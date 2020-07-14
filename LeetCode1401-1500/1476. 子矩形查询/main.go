package main

// https://leetcode-cn.com/problems/subrectangle-queries/

type SubrectangleQueries struct {
	data [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{data: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	if row1 < 0 || row1 >= len(this.data) {
		return
	}
	if col1 < 0 || col1 >= len(this.data[row1]) {
		return
	}
	if row2 < row1 || row2 >= len(this.data) {
		return
	}
	if col2 < col1 || col2 >= len(this.data[row2]) {
		return
	}
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.data[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	if row < 0 || row >= len(this.data) {
		return -1
	}
	if col < 0 || col >= len(this.data[row]) {
		return -1
	}
	return this.data[row][col]
}

func main() {

}
