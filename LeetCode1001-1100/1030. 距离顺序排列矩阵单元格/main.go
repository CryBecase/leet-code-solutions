package main

// https://leetcode-cn.com/problems/matrix-cells-in-distance-order/

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDistance := R + C - 2

	res := make([][]int, 0)
	res = append(res, []int{r0, c0})
	distance := 1
	for distance <= maxDistance {
		nr, nc := r0+distance, c0
		if nr < R && nc < C && nr >= 0 && nc >= 0 {
			res = append(res, []int{nr, nc})
		}
		for i := 1; i <= distance; i++ {
			if nr-i < R && nc-i < C && nr-i >= 0 && nc-i >= 0 {
				res = append(res, []int{nr - i, nc - i})
			}
			if nr-i < R && nc+i < C && nr-i >= 0 && nc+i >= 0 {
				res = append(res, []int{nr - i, nc + i})
			}
		}

		nr, nc = r0-distance, c0
		if nr < R && nc < C && nr >= 0 && nc >= 0 {
			res = append(res, []int{nr, nc})
		}
		for i := 1; i < distance; i++ {
			if nr+i < R && nc-i < C && nr+i >= 0 && nc-i >= 0 {
				res = append(res, []int{nr + i, nc - i})
			}
			if nr+i < R && nc+i < C && nr+i >= 0 && nc+i >= 0 {
				res = append(res, []int{nr + i, nc + i})
			}
		}

		distance++
	}

	return res
}

func main() {

}
