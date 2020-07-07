package main

// https://leetcode-cn.com/problems/maximum-sum-of-two-non-overlapping-subarrays/

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	// 计算前缀和
	for i := 1; i < len(A); i++ {
		A[i] += A[i-1]
	}

	res := A[L+M-1]
	Lmax, Mmax := A[L-1], A[M-1]
	for i := L + M; i < len(A); i++ {
		Lmax = maxFunc(Lmax, A[i-M]-A[i-M-L]) // A[i-M]-A[i-M-L] 表示连续长度为L的子数组和
		Mmax = maxFunc(Mmax, A[i-L]-A[i-L-M]) // A[i-M]-A[i-M-L] 表示连续长度为L的子数组和
		res = maxFunc(res, maxFunc(Lmax+A[i]-A[i-M], Mmax+A[i]-A[i-L]))
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
