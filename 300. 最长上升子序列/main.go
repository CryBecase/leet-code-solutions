package main

// https://leetcode-cn.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make([]int, length)
	dp[0] = 1

	curMax := 0
	result := 1
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				curMax = maxFunc(dp[j], curMax)
			}
		}
		dp[i] = curMax + 1
		result = maxFunc(dp[i], result)
		curMax = 0
	}
	return result
}

func maxFunc(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {

}
