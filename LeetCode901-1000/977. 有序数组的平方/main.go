package main

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(A []int) []int {
	L, R := 0, 0
	for R < len(A) && A[R] <= 0 {
		R++
	}
	L = R - 1
	res := make([]int, 0)
	for R < len(A) && L >= 0 {
		if A[R]+A[L] > 0 {
			res = append(res, A[L]*A[L])
			L--
		} else {
			res = append(res, A[R]*A[R])
			R++
		}
	}
	for R < len(A) {
		res = append(res, A[R]*A[R])
		R++
	}
	for L >= 0 {
		res = append(res, A[L]*A[L])
		L--
	}
	return res
}

func main() {

}
