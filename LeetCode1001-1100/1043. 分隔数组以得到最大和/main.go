package main

// https://leetcode-cn.com/problems/partition-array-for-maximum-sum/

func maxSumAfterPartitioning(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	dp := make([]int, len(arr))

	max := arr[0]
	for i := 0; i < len(dp); i++ {
		if i < k {
			max = maxFunc(max, arr[i])
			dp[i] = max * (i + 1)
		} else {
			max := arr[i]
			for j := 0; j < k; j++ {
				max = maxFunc(max, arr[i-j])
				dp[i] = maxFunc(dp[i], dp[i-j-1]+max*(j+1))
			}
		}
	}
	return dp[len(dp)-1]
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
