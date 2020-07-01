package main

// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	// init
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	res := 0
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			res = maxFunc(res, dp[i][j])
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
	findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1})
}
