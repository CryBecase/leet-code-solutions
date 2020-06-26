package main

// https://leetcode-cn.com/problems/best-sightseeing-pair/

func maxScoreSightseeingPair(A []int) int {
	ans := 0
	maxI := A[0] - 0
	for i := 1; i < len(A); i++ {
		if A[i-1]+(i-1) > maxI {
			maxI = A[i-1] + (i - 1)
		}
		if maxI+A[i]-i > ans {
			ans = maxI + A[i] - i
		}
	}
	return ans
}

func main() {

}
