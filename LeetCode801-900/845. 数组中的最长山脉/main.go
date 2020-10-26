package main

// https://leetcode-cn.com/problems/longest-mountain-in-array/

func longestMountain(A []int) int {
	L, R := 0, 0
	res := 0
	for R < len(A) {
		L = R
		// 从山脚到山顶
		for R+1 < len(A) && A[R] < A[R+1] {
			R++
		}
		// 没有山顶 重置山脚
		if L == R {
			R++
			continue
		}
		// 从山顶到山脚
		hasRight := false
		for R+1 < len(A) && A[R] > A[R+1] {
			hasRight = true
			R++
		}
		if hasRight {
			res = maxFunc(res, R-L+1)
		}
	}
	return res
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
