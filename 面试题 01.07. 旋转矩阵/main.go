package main

// https://leetcode-cn.com/problems/rotate-matrix-lcci/

// (r, c) -> (c, n - r - 1)
// 等价于
// (r, c) 上下翻转-> (n - r - 1, c) 沿对角线折叠-> (c, n - r - 1)
func rotate(matrix [][]int) {
	n := len(matrix)
	for r := range matrix {
		if r < n/2 {
			for c := range matrix[r] {
				matrix[r][c], matrix[n-r-1][c] = matrix[n-r-1][c], matrix[r][c]
			}
		}
	}
	//for r := range matrix {
	//	for c := range matrix[r] {
	//		fmt.Printf("%d ", matrix[r][c])
	//	}
	//	fmt.Println()
	//}

	for r := range matrix {
		for c := 0; c < r; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
}

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}
