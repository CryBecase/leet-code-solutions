package main

// https://leetcode-cn.com/problems/sorted-merge-lcci/

func merge(A []int, m int, B []int, n int) {
	helper := make([]int, m)
	for i := 0; i < m; i++ {
		helper[i] = A[i]
	}
	i, j, index := 0, 0, 0
	for i < m && j < n {
		if helper[i] <= B[j] {
			A[index] = helper[i]
			i++
		} else {
			A[index] = B[j]
			j++
		}
		index++
	}
	for i < m {
		A[index] = helper[i]
		i++
		index++
	}
	for j < n {
		A[index] = B[j]
		j++
		index++
	}
}

func main() {

}
